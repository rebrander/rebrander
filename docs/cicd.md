# CI/CD

Local Go is not required for the current workstation flow. GitHub Actions is the source of validation.

## Workflows

### CI

`.github/workflows/ci.yml` runs on pushes and pull requests to `main`.

Checks:

- `gofmt` for all Go files.
- `go test ./...`.
- `go build ./cmd/rebrander`.
- TypeScript frontend check and static build.

### Release

`.github/workflows/release.yml` runs for tags matching `v*.*.*`.

It builds binaries for:

- Linux amd64/arm64.
- macOS amd64/arm64.
- Windows amd64.

The resulting artifacts are attached to the GitHub release.

### Pages

`.github/workflows/pages.yml` builds the TypeScript frontend and deploys `web/dist` to GitHub Pages. The static site writes a `CNAME` file for `rebrander.sh`.

## Deployment Start

The first deployment step is pushing `main`. GitHub Actions should then create the first CI run and Pages deployment. Repository settings still need to enable GitHub Pages with GitHub Actions as the source if GitHub does not auto-configure it.

