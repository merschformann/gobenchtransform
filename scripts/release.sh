#!/bin/bash

set -e

# This script is used to release a new version of the project.

# Check that the current branch is main.
if [[ $(git branch --show-current) != "main" ]]; then
  echo "You must be on the main branch to release a new version."
  exit 1
fi

# Check that the working directory is clean.
if [[ $(git status --porcelain) ]]; then
  echo "You have uncommitted changes. Please commit or stash them before releasing a new version."
  exit 1
fi

# Check that the current branch is up to date with the remote.
if [[ $(git rev-list HEAD...origin/main --count) != "0" ]]; then
  echo "Your local main branch is not up to date with the remote. Please pull the latest changes before releasing a new version."
  exit 1
fi

# Get the current version.
VERSION="$(cat VERSION)"

# Check that the tag does not already exist.
if git rev-parse "$VERSION" >/dev/null 2>&1; then
  echo "Version $VERSION already exists. Please update the VERSION file and try again."
  exit 1
fi

# Get notes, if provided.
NOTES="$1"
if [[ -z "$NOTES" ]]; then
  NOTES="--generate-notes"
else
  NOTES="--notes $NOTES"
fi

# Ask for confirmation.
read -p "Releasing version $VERSION. Are you sure? (y/n) " -n 1 -r

# Release.
gh release create "$VERSION" $NOTES --title "$VERSION"

