package report

import "strings"

type Model struct {
	Title   string
	Source  string
	Target  string
	Mode    string
	Entries []string
}

func Markdown(report Model) string {
	var b strings.Builder
	b.WriteString("# " + report.Title + "\n\n")
	b.WriteString("- Source: `" + report.Source + "`\n")
	b.WriteString("- Target: `" + report.Target + "`\n")
	b.WriteString("- Mode: `" + report.Mode + "`\n\n")
	b.WriteString("## Name Mappings\n\n")
	if len(report.Entries) == 0 {
		b.WriteString("No mappings detected.\n")
		return b.String()
	}
	for _, entry := range report.Entries {
		b.WriteString("- " + entry + "\n")
	}
	return b.String()
}
