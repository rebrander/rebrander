# Product Concept

Rebrander changes the identity of a project while preserving its architecture and behavior.

## Inputs

- Source GitHub repository.
- Source GitLab repository.
- Source GitHub organization.
- Source GitLab group.
- Target repository or organization path.

## Outputs

- Rebranded repository or repositories.
- Change report.
- Infrastructure checklist.
- Secret setup checklist.
- Validation report.
- Provenance report when needed.

## Modes

### Safe

Only high-confidence mappings are applied. Everything else becomes a review item.

### Reviewed

Medium-risk mappings require user confirmation.

### Aggressive

Broad rewrite for owned projects. This mode is not appropriate for unclear third-party repositories or hosted SaaS jobs without proof of authorization.

