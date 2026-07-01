## Contributing

Hi! Thanks for your interest in contributing to the GitHub CLI Internationalization (i18n) & Localization (l10n) project.

This project is a fork of the official GitHub CLI with the goal of adding multi-language support. All contributions must remain compatible with the upstream repository.

---

### Development Environment

Prerequisites:
- Go 1.26+
- Git 2.30+
- A GitHub account

Build with:
- Unix-like systems: `make`
- Windows: `go run script/build.go`

Run the new binary as:
- Unix-like systems: `bin/gh`
- Windows: `bin\gh`

Run tests with: `go test ./...`

---

### Branch Workflow

We follow a simple branching strategy:

- `trunk` — Mirrors the upstream `cli/cli` `trunk` branch. Do not commit directly here.
- `develop` — Base branch for all feature development. All new work branches from and merges into `develop`.

Feature branches use the pattern: `feature/<short-description>` or `fix/<short-description>`.

---

### Syncing with Upstream

To keep your fork in sync with the official GitHub CLI repository:

```bash
# Fetch upstream changes
git fetch upstream

# Update trunk (mirror of upstream)
git checkout trunk
git merge upstream/trunk

# Rebase develop on trunk
git checkout develop
git rebase trunk
```

---

### Submitting a Pull Request

1. Create a new branch from `develop`:
   ```bash
   git checkout -b feature/my-feature develop
   ```
2. Make your change, add tests, and ensure tests pass:
   ```bash
   go test ./...
   ```
3. Keep your branch updated:
   ```bash
   git fetch upstream
   git rebase upstream/trunk
   ```
4. Submit a pull request against the `develop` branch:
   ```bash
   gh pr create --base develop
   ```

### Pull Request Checklist

Before submitting your PR:
- [ ] Code compiles without errors (`go build` or `make`)
- [ ] Tests pass (`go test ./...`)
- [ ] No lint warnings (`make lint`)
- [ ] Branch is up to date with `develop`
- [ ] Commit messages are clear and descriptive
- [ ] Changes do not break compatibility with upstream

---

### Coding Style

- Follow the existing code style of the project
- Use `gofmt` / `go fmt` before committing
- Run `make lint` (golangci-lint) before submitting a PR
- Add godoc comments to all exported functions, types, and constants

---

### Issue Templates

Please use the appropriate template when filing issues:
- [Bug report](.github/ISSUE_TEMPLATE/bug_report.md)
- [Feature request](.github/ISSUE_TEMPLATE/submit-a-request.md)
- [i18n/l10n specific request](.github/ISSUE_TEMPLATE/i18n-request.md)

---

### Code of Conduct

Please note that this project adheres to a [Contributor Code of Conduct](./CODE-OF-CONDUCT.md). By participating in this project you agree to abide by its terms.

### Resources

- [How to Contribute to Open Source](https://opensource.guide/how-to-contribute/)
- [Using Pull Requests](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests)
- [GitHub Help](https://docs.github.com/)

Contributions to this project are [released](https://docs.github.com/en/free-pro-team@latest/github/site-policy/github-terms-of-service#6-contributions-under-repository-license) to the public under the [project's open source license](../LICENSE).
