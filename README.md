# Go Web Test

This is a sample repository to illustrate best practices
for using pre-commit hooks and setup pipeline actions to
ensure code quality and security.

This applicaiton is a simple go web server that runs as a perl script.
The web server will display excuses, umm, I am oppritunities that
make be causing problems.

The applicaition will be containerized and could be desployed using
docker, docker-compose, or kubernetes.

## Pre-commit Hooks

We are using pre-commit hooks to make sure things are done correctly.
Some of the things are want to make sure is:

* Code is formatted correctly
* Code is linted correctly
* Code is secure
* Code is tested

### Code Formatting

We are using gofmt to format the code. As a TODO, we should do
formatting and checking on the perl code. We shouldl replace the
perl code with Go at some point.

First we do some basic formatting checks on files:

* Trailing whitespaces are removed.
* The end of the file has a new line.
* Make sure YAML files are formatted correctly.

Then we do specific checks on go files:

* gofmt
* goimports
* We make sure the go.* TOML files are formatted correctly.

## Linting

Here we dont have any business specific linting rules to check,
but we are using golang-ci-lint to check the code for common
problems.

* golangci-lint

## Security

Here we check for common security issues:

* check for AWS credentials
* check for secrets
* go vet and go security checks should be run (TODO)

## Testing

Go testing is done with go test. We are using the following
tools to test the code:

* go test
* go vet
* go-acc

## Pipeline Actions

Here we are using pipeline actions to make sure the code is
formatted, linted, and tested.

### Code Formatting

We are using the following pipeline actions to make sure the code
is formatted correctly:

* golangci-lint
* gofmt

### Linting

We are using the following pipeline actions to make sure the code
is linted correctly:

* golangci-lint

### Testing

We are using the following pipeline actions to make sure the code
is tested correctly:

* go test

At the present time, we do NOT check that test are created (TOD0)
and we do not check the overall coverage of the tests.

We also have a local test action to test the perl code.

* cgi-test
