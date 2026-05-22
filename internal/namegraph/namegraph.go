package namegraph

import (
	"net/url"
	"path"
	"strings"
)

type Mapping struct {
	Source     string
	Target     string
	Category   string
	Confidence string
}

type Graph struct {
	Mappings []Mapping
}

func Build(sourceURL, targetURL string) Graph {
	source := identityParts(sourceURL)
	target := identityParts(targetURL)
	mappings := []Mapping{
		{Source: source.Full, Target: target.Full, Category: "url", Confidence: "high"},
		{Source: source.Owner, Target: target.Owner, Category: "namespace", Confidence: "high"},
		{Source: source.Repo, Target: target.Repo, Category: "repository", Confidence: "high"},
	}
	if source.Base != "" && target.Base != "" {
		mappings = append(mappings,
			Mapping{Source: strings.ToLower(source.Base), Target: strings.ToLower(target.Base), Category: "brand-lower", Confidence: "high"},
			Mapping{Source: title(source.Base), Target: title(target.Base), Category: "brand-title", Confidence: "high"},
			Mapping{Source: strings.ToUpper(source.Base), Target: strings.ToUpper(target.Base), Category: "brand-upper", Confidence: "high"},
		)
	}
	return Graph{Mappings: dedupe(mappings)}
}

type parts struct {
	Full  string
	Owner string
	Repo  string
	Base  string
}

func identityParts(raw string) parts {
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Host == "" {
		trimmed := strings.Trim(raw, "/")
		return parts{Full: trimmed, Repo: path.Base(trimmed), Base: brandBase(path.Base(trimmed))}
	}
	segments := strings.Split(strings.Trim(parsed.Path, "/"), "/")
	out := parts{Full: strings.TrimRight(raw, "/")}
	if len(segments) > 0 {
		out.Owner = segments[0]
	}
	if len(segments) > 1 {
		out.Repo = segments[len(segments)-1]
		out.Base = brandBase(out.Repo)
	}
	return out
}

func brandBase(repo string) string {
	repo = strings.TrimSuffix(repo, ".git")
	repo = strings.TrimSuffix(repo, ".com")
	fields := strings.FieldsFunc(repo, func(r rune) bool {
		return r == '-' || r == '_' || r == '.'
	})
	if len(fields) == 0 {
		return repo
	}
	return fields[0]
}

func title(value string) string {
	if value == "" {
		return value
	}
	return strings.ToUpper(value[:1]) + strings.ToLower(value[1:])
}

func dedupe(input []Mapping) []Mapping {
	seen := map[string]bool{}
	out := make([]Mapping, 0, len(input))
	for _, item := range input {
		key := item.Source + "\x00" + item.Target + "\x00" + item.Category
		if item.Source == "" || item.Target == "" || seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, item)
	}
	return out
}

