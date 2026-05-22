package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rebrander/rebrander/internal/namegraph"
	"github.com/rebrander/rebrander/internal/report"
)

type Options struct {
	Source    string
	Target    string
	Mode      string
	OutputDir string
	PlanOnly  bool
	LocalOnly bool
	Push      bool
}

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context, opts Options) error {
	if err := opts.validate(); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	plan := Plan{
		Source: opts.Source,
		Target: opts.Target,
		Mode:   opts.Mode,
		Names:  namegraph.Build(opts.Source, opts.Target),
	}
	if err := os.MkdirAll(opts.OutputDir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}
	path := filepath.Join(opts.OutputDir, "REBRANDER_REPORT.md")
	if err := os.WriteFile(path, []byte(report.Markdown(plan.Report())), 0o644); err != nil {
		return fmt.Errorf("write report: %w", err)
	}
	fmt.Printf("plan written to %s\n", path)
	return nil
}

func (o Options) validate() error {
	if o.Source == "" {
		return fmt.Errorf("source is required")
	}
	if o.Target == "" {
		return fmt.Errorf("target is required")
	}
	switch o.Mode {
	case "safe", "reviewed", "aggressive":
		return nil
	default:
		return fmt.Errorf("invalid mode %q", o.Mode)
	}
}

