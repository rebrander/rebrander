package app

import (
	"github.com/rebrander/rebrander/internal/namegraph"
	"github.com/rebrander/rebrander/internal/report"
)

type Plan struct {
	Source string
	Target string
	Mode   string
	Names  namegraph.Graph
}

func (p Plan) Report() report.Model {
	entries := make([]string, 0, len(p.Names.Mappings))
	for _, mapping := range p.Names.Mappings {
		entries = append(entries, mapping.Source+" -> "+mapping.Target+" ("+mapping.Category+")")
	}
	return report.Model{
		Title:   "Rebrander Plan",
		Source:  p.Source,
		Target:  p.Target,
		Mode:    p.Mode,
		Entries: entries,
	}
}
