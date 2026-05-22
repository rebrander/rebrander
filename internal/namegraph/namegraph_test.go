package namegraph

import "testing"

func TestBuildIncludesBrandVariants(t *testing.T) {
	graph := Build("https://gitlab.com/netol/netol.com", "https://gitlab.com/polococa/polococa")
	got := map[string]string{}
	for _, mapping := range graph.Mappings {
		got[mapping.Source] = mapping.Target
	}
	if got["netol"] != "polococa" {
		t.Fatalf("expected lowercase brand mapping, got %#v", got)
	}
	if got["Netol"] != "Polococa" {
		t.Fatalf("expected title brand mapping, got %#v", got)
	}
	if got["NETOL"] != "POLOCOCA" {
		t.Fatalf("expected upper brand mapping, got %#v", got)
	}
}

