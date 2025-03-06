# PRE-COMMIT

In order to make sure that processes are followed, we are using [pre-commit](https://pre-commit.com/)
during the software develop lifecycle to improve the CI/CD process.

For this to work correctly, the following will need to be installed on the developers workstation.

* pre-commit
* goimports
  * `go install golang.org/x/tools/cmd/goimports@latest`
* golangci-lint
  * `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

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
