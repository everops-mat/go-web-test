# PRE-COMMIT

In order to make sure that processes are followed, we are using [pre-commit](https://pre-commit.com/)
during the software develop lifecycle to improve the CI/CD process.

For this to work correctly, the following will need to be installed on the developers workstation.

* pre-commit
* goimports
  * `go install golang.org/x/tools/cmd/goimports@latest`
* golangci-lint
  * `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

## Setup

There is a make file that will install the pre-commit hooks in your local copy of the repository.

`make pre-commit`

Occassionaly, we'll want to run `make pre-commit-update` to update the plugins.

## Basic tests

We use some base test from pre-commit:

```yaml
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v5.0.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
      - id: check-merge-conflict
      - id: check-shebang-scripts-are-executable
      - id: detect-private-key
```

* trailing-whitespace - removes extra whitespace at end of line of files.
* end-of-file-fixer - make sure all files end with a new line.
* check-yaml: check that yaml files are formatted correctly.
* check-add-large-files: we don't wish to add large files into the repository.
* check-merge-conflict: make sure we don't have merge conflicts.
* check-shebang-scripts-are-exectable: make sure our scripts and able to be run.
* detect-private-key: Try to make sure we don't including private keys in a commit.

All these commands are run pre-commit, if they fail, your commit will fail. You can correct errors,
add your changes, and attempt the commit again. You also can run pre-commit on all your files by
running `pre-commit run --all-files`, which can be done using make, `make run-pre-commit`.

## Branch Name Enforcement

There is a custom script that will enforce the branch name schema for this repository.

We require the following branch names:

* feature/
* bugfix/
* hotfix/
* ci/
* relealse/

Which must be used on prefixes for new branches.

The [script](/scripts/enforce-branch-naming.sh) will be run from pre-commit once for each commit.
If you branch does not following the naming requirement, pre-commit will fail.

This is controled in the [.pre-commit-config.yaml](/.pre-commit-config.yaml) with the following
setup

```yaml
  - repo: local
    hooks:
      - id: enforce-branch-naming
        name: Enforce Branch Naming
        entry: bash ./scripts/enforce-branch-naming.sh
        language: system
        pass_filenames: false
        always_run: true
```

There is also a [GitHub Action](/.github/workflows/enforce-branch.yml) which will check push and
pull request to make sure proper branch naming is enforced.

__NOTE: If the name schema is changed, both script and Github Action needs to be updated.__

## Code Specific Pre Commit tasks

### Go specific pre-commit checks
```yaml
  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.5.1
    hooks:
      - id: validate-toml
      - id: go-mod-tidy
      - id: go-imports
      - id: go-fmt
      - id: go-vet
         name: Run go vet
         entry: go vet ./...
         language: system
         pass_filenames: false
         always_run: true
      - id: golangci-lint
      - id: no-go-testing
      - id: go-unit-tests
```

First we to validate any TOML configuration files that might be present, after that:

* Run `go mod tidy`
* Make sure Go imports are defined correctly.
* Format the Go source code files are formatted correctly.
* Run `go vet` to check for common errors.
* Run linting on the Go files looking for common errors.
* Check if any Go unit test files are NOT using Go's testing package.
* Run Go unit tests.
  * These are files that have `//go:build unit` at the start of the testing modules.
  * This will also include files that do NOT have any `//go build:` lines defined.
  * For specific intergration test, use `go:build intergration` at the time of the file, they will be skipped for the unit testing.

### Dockerfile specific pre-commit checks

```yaml
  - repo: https://github.com/AleksaC/hadolint-py
    rev: v2.12.1b3
    hooks:
      - id: hadolint
        args: ["--ignore", "DL3008", "--ignore", "DL3007", "--ignore", "DL4006"]
```
We run Hadolint on the Dockerfiles, but we ignore two  rules:

* DL3007 - We we allo the use of the tag latest.
* DL3008 - We may need cases where we pin version using `apt-get`.
