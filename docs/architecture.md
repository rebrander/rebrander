# Architecture

Rebrander has two product surfaces:

- Go CLI for local, auditable rebranding of projects owned by the user or explicitly authorized by a client.
- TypeScript web frontend for `rebrander.sh`, backed later by a SaaS API and isolated workers.

## System Shape

```text
CLI / Web UI
  -> Orchestrator
    -> Source connectors: GitHub, GitLab, local git
    -> Discovery engine: files, languages, CI, infra, assets, binaries, secrets
    -> Name graph: source and target identity variants
    -> Transform engine: structured config, AST, text, path renames, assets
    -> Validation engine: tests, builds, secret scan, license scan
    -> Publish engine: target repo creation, branch push, PR/MR, release artifacts
    -> Report engine: markdown/json/html reports
```

## Rebranding Pipeline

1. Parse source and target URLs.
2. Detect whether the input is a single repository or an organization/group.
3. Clone repositories into an isolated workspace.
4. Discover project identity: slug, owner, package names, domains, config keys, image names, CI variables, and infrastructure names.
5. Build the name graph with confidence levels.
6. Generate a dry-run plan.
7. Apply structured transformations first, then AST transformations, then reviewed text fallback.
8. Rename paths after content transforms.
9. Run validation in CI/CD.
10. Generate reports and publish to the target repository when explicitly requested.

## Go Packages

- `cmd/rebrander`: CLI entrypoint.
- `internal/app`: orchestration and options validation.
- `internal/namegraph`: initial mapping of source identity to target identity.
- `internal/report`: markdown reports.

Future packages:

- `internal/connectors/github`
- `internal/connectors/gitlab`
- `internal/discovery`
- `internal/transform`
- `internal/validate`
- `internal/publish`

## Frontend

The `web/` directory contains the `rebrander.sh` frontend. It is intentionally static in the first phase so GitHub Pages can host it while the backend API is still being designed.

