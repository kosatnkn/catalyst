#!/bin/bash

echo "Generating metadata..."

# add static metadata
cat metadata.txt > metadata/base.txt

# add build info
tag=$(git describe --exact-match --tags HEAD)
if [ $? -eq 0 ]; then
  release=", release [${tag}]"
fi

echo "built from [$(git rev-parse --abbrev-ref HEAD)] branch, commit [$(git rev-parse --short HEAD)]${release}" > metadata/build.txt
echo "build date $(date +"%Y.%m.%d %T")" >> metadata/build.txt
