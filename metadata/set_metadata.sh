#!/bin/bash

work_dir=$1

echo "Generating metadata..."

# add static metadata
tr '\n\r\t' ' ' < "${work_dir}/metadata.txt" | tr -s ' ' | sed 's/^ *//;s/ *$//' > "${work_dir}/metadata/base.txt"

# add build info
tag=$(git describe --exact-match --tags HEAD 2>/dev/null)
if [ $? -gt 0 ]; then
  tag="none"
fi

repo=$(basename `git rev-parse --show-toplevel` | tr '\n' ' ')

echo "${repo}" > ${work_dir}/metadata/name.txt

echo "repository: ${repo}
branch: $(git rev-parse --abbrev-ref HEAD)
commit: $(git rev-parse --short HEAD)
release: ${tag}
time: $(date +"%Y.%m.%d %T (%z)")" > ${work_dir}/metadata/build.txt
