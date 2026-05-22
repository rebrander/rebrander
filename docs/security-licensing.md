# Security and Licensing

Rebrander is designed primarily for projects owned by the operator or projects the operator is explicitly authorized to modify.

## Local and Owned-Project Mode

When the source and target projects belong to the same owner, company, or authorized client, Rebrander may perform a full transformation:

- Project names.
- Package names.
- Domains and URLs.
- CI/CD identifiers.
- Infrastructure names.
- Assets.
- Rebuilt binaries when the user owns or controls the source.

The generated report should record that the user declared the right to perform the operation.

## SaaS Compliance Mode

The hosted `rebrander.sh` service must be stricter than the local CLI:

- It must require confirmation of ownership or authorization.
- It must respect license, copyright, and trademark constraints.
- It must keep or generate provenance reports for third-party code.
- It must block or downgrade unclear jobs to plan-only mode.
- It must never copy secret values from source repositories.

## Secrets

Rebrander should detect and report secret locations, but it must not migrate raw secret values. Target projects should receive a checklist such as `REBRANDER_SECRETS.md` with names, scopes, and setup instructions.

## License

The Rebrander repository uses MIT. This license applies to Rebrander itself, not to repositories processed by Rebrander.

