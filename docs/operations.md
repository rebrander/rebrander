# Operations

## Repository

Target repository: `github.com/rebrander/rebrander`.

The repository is public. The product code is MIT licensed. Users remain responsible for having the legal right to rebrand source repositories.

## Required GitHub Permissions

The operator account needs:

- `repo` for repository creation, push, settings, and releases.
- `workflow` for GitHub Actions workflow files.
- `read:org` for organization visibility checks.

## Local Development Constraint

The current workstation does not have Go installed. Do not rely on local Go commands. Use GitHub Actions for Go validation.

Local actions allowed:

- Edit files.
- Commit.
- Push.
- Inspect GitHub Actions via `gh`.

## SaaS Operations Direction

The future SaaS must run rebranding jobs in isolated workers:

- One job per container or microVM.
- Short-lived GitHub/GitLab tokens.
- Masked logs.
- Workspace destruction after job completion.
- Plan-only fallback when license or ownership is unclear.

