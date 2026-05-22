# Rebrander

Rebrander is a Go-first CLI and future SaaS for controlled project rebranding. It clones a GitHub or GitLab project, discovers project-name references, builds a rename plan, applies safe transformations, validates the result in CI/CD, and prepares the target repository with reports and infrastructure instructions.

The primary use case is rebranding projects you own or are authorized to modify. The public SaaS flow for `rebrander.sh` must only execute legal changes that respect repository ownership, license terms, trademarks, and copyright.

## Status

This repository is in bootstrap phase. The initial implementation provides:

- Go CLI skeleton.
- Name variant graph for brand/project identifiers.
- Dry-run planning model.
- Markdown report generation.
- TypeScript frontend scaffold for `rebrander.sh`.
- GitHub Actions CI/CD for Go and frontend validation.
- Architecture, security, licensing, operations, and roadmap documentation.

## CLI

```bash
rebrander <source> <target> [flags]
```

Example:

```bash
rebrander https://github.com/original/project https://github.com/newbrand/project --plan-only
```

Current bootstrap flags:

```text
--plan-only      Generate a plan without writing target repositories.
--local-only     Write local output only.
--push           Push to the target repository after validation.
--mode           safe, reviewed, or aggressive.
--output         Output directory for reports.
```

## Development

The project language is Go. The frontend under `web/` is TypeScript.

Local Go is not required for the initial workstation flow. CI/CD performs Go formatting checks, tests, and builds on GitHub-hosted runners.

## Documentation

- [Architecture](docs/architecture.md)
- [CI/CD](docs/cicd.md)
- [Operations](docs/operations.md)
- [Security and licensing](docs/security-licensing.md)
- [Roadmap](docs/roadmap.md)

## License

MIT. See [LICENSE](LICENSE).

