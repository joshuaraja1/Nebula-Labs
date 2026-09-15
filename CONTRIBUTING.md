# Contributing to Nebula Labs

Thank you for helping improve Nebula Labs. Contributions should be focused,
reviewable, tested, and clear about the behavior they change.

## Development workflow

1. Create a focused branch from `main`.
2. Make one logical change per commit.
3. Format and test the affected Go module.
4. Update documentation when behavior or configuration changes.
5. Open a pull request that explains the problem, approach, and verification.

Suggested branch names:

- `feature/<short-description>`
- `fix/<short-description>`
- `docs/<short-description>`
- `test/<short-description>`

## Local verification

Run these commands from the component you changed:

```bash
gofmt -w .
go test ./...
go vet ./...
```

For `services/nebula-api`, copy `.env.template` to `.env` only when local
integration testing requires configured services. Never commit credentials,
tokens, API keys, or personal data.

## Commit guidelines

Use concise, imperative commit subjects:

- `Add scholarship schema validation`
- `Fix empty parser input handling`
- `Document uploader retry behavior`

Avoid unrelated formatting changes in the same commit as a functional change.

## Pull-request checklist

- [ ] The change has a single clear purpose.
- [ ] New behavior includes tests or an explanation of why tests are unnecessary.
- [ ] `gofmt`, `go test ./...`, and `go vet ./...` pass for the affected module.
- [ ] Public APIs, schemas, and environment variables are documented.
- [ ] No secrets, generated binaries, or personal data are included.
- [ ] Attribution and upstream license notices remain intact.

## Reporting issues

A useful issue includes:

- the expected behavior;
- the actual behavior;
- steps or a minimal input that reproduces the problem;
- relevant logs with secrets removed;
- the operating system and Go version.

## Licensing and attribution

By contributing, you agree that your contribution may be distributed under the
license applicable to the component you modify. Do not remove upstream copyright,
license, or attribution notices.
