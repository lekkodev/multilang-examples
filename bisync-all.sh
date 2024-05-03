#!/usr/bin/env bash
set -Eeuo pipefail

pushd golang >/dev/null
lekko bisync --repo-path .. > /dev/null
if [ -n "$(git status --porcelain)" ]; then
  echo "Go: detected changes!"
else
  echo "Go: no changes"
fi
popd >/dev/null

pushd ts >/dev/null
lekko bisync --repo-path .. > /dev/null
if [ -n "$(git status --porcelain)" ]; then
  echo "TypeScript: detected changes!"
else
  echo "TypeScript: no changes"
fi
popd >/dev/null
