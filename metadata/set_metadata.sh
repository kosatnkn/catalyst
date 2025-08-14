#!/bin/bash

work_dir=$1

echo "Generating metadata..."

# add static metadata
cat ${work_dir}/metadata.txt > ${work_dir}/metadata/base.txt

# add build info
tag=$(git describe --exact-match --tags HEAD 2>/dev/null)
if [ $? -eq 0 ]; then
  release=", release [${tag}]"
fi

echo "built from [$(git rev-parse --abbrev-ref HEAD)] branch, commit [$(git rev-parse --short HEAD)]${release}" > ${work_dir}/metadata/build.txt
echo "built on $(date +"%Y.%m.%d %T (%z)")" >> ${work_dir}/metadata/build.txt
