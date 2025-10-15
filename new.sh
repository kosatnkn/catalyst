#!/bin/bash

HELP="
Help here
yes

mulvjv

"

MODULE_NAME=$1
if [ -z ${MODULE_NAME} ]; then
  echo "error: a go module name is not provided for the new mocroservice"
  echo -e ${HELP}
  exit 1
fi

CATALYST_BASE="https://github.com/kosatnkn/catalyst.git"
CATALYST_VER="v2.9.0" # must be the release version TODO: need to enforce checking this when trying to create new tag

git clone --branch="${CATALYST_VER}" --depth=1 "${CATALYST_BASE}"

# remove following dirs and files
# - .git/
# - .github/
# - docs/img/*.drawio.*
# - LICENSE
# - README.md

# replace following
# github.com/kosatnkn/catalyst/v3 with <new_module_name>

# create following dirs and files
