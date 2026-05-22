package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/rebrander/rebrander/internal/app"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "rebrander: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet("rebrander", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	opts := app.Options{}
	flags.BoolVar(&opts.PlanOnly, "plan-only", false, "generate a plan without writing target repositories")
	flags.BoolVar(&opts.LocalOnly, "local-only", false, "write local output only")
	flags.BoolVar(&opts.Push, "push", false, "push to the target repository after validation")
	flags.StringVar(&opts.Mode, "mode", "safe", "safe, reviewed, or aggressive")
	flags.StringVar(&opts.OutputDir, "output", "rebrander-output", "output directory for reports")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 2 {
		return fmt.Errorf("usage: rebrander <source> <target> [flags]")
	}
	opts.Source = flags.Arg(0)
	opts.Target = flags.Arg(1)
	return app.New().Run(context.Background(), opts)
}

