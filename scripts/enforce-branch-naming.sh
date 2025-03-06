#!/bin/bash

# a script to enforce branch naming conventions
#
# this script is meant to be run as a pre-commit hook
# exit 1 if ther is an error
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
if [[ -z "$BRANCH_NAME" ]]; then
  echo "no branch name specified"
  exit 1
fi

ALLOWED_BRANCH_PREFIXES="feature/* hotfix/* bugfix/* ci/* release/*"

if [[ -z "$ALLOWED_BRANCH_PREFIXES" ]]; then
	echo "ERROR: Invalid branch name prefix list '$ALLOWED_BRANCH_PREFIXES'"
	exit 1
fi
REGEX=$(echo "$ALLOWED_BRANCH_PREFIXES" | sed 's/ /|/g' | sed 's/\*/[a-z0-9._-]*/g')

# Check if the branch name matches any of the allowed patterns
if [[ ! "$BRANCH_NAME" =~ ^($REGEX)$ ]]; then
    echo "ERROR: Invalid branch name '$BRANCH_NAME'"
    echo "Allowed branch name patterns: $ALLOWED_BRANCH_PREFIXES"
    echo ""
    echo "Rename your branch using:"
    echo "  git branch -m <new-valid-branch-name>"
    exit 1
fi

echo "Branch name '$BRANCH_NAME' is valid."
exit 0
