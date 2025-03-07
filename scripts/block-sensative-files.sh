#!/bin/bash
# Block files that are sensitive to be committed

FILES=$(git diff --cached --name-only| grep  -E '(\.env$|config\.json$|secrets\.json$|id_rsa$)')

if [ -n "$FILES" ]; then
    echo "The following files are not allowed in the git repo:"
    echo "$FILES"
    exit 1
fi

exit 0
