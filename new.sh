#!/bin/bash
set -e

CATALYST_BASE="https://github.com/kosatnkn/catalyst"
CATALYST_REF="<ph_ref>"
CATALYST_MODULE="github.com/kosatnkn/catalyst/v3"

show_help() {
  echo -e "
Usage: $0 --module=<go_module_name> [--dir=<working_dir>]
Options:
  --module     (Required) Go module path (e.g., example.com/sampler)
  --dir        (Optional) Target directory (default: current directory)
  -y, --yes    Answer 'yes' to confirmation prompt
  -h, --help   Display this help message
Example:
  $0 --module example.com/dummyuser/sampler --dir ./projects"
}

# messages
msg_info() {
  echo -e "⦁ $1"
}

msg_err() {
  echo -e "⨉ $1"
}

msg_success() {
  echo -e "✔ $1"
}

msg_ongoing() {
  echo -e "⨠ $1"
}

msg_done() {
  echo -e "✔ Done"
}

msg_complete() {
  echo -e "✔✔ Complete!"
}

prompt_yes_no() {
  local msg="∆ Are you sure?"
  local decline="⦸ Declined, Stopped!"
  local proceed="⨠ Proceeding..."

  while true; do
    read -r -p "${msg} [Y/n]: " input

    case ${input} in
    [yY][eE][sS] | [yY])
      echo -e "${proceed}"
      return 0
      ;;
    [nN][oO] | [nN])
      echo -e "${decline}"
      exit 1
      ;;
    esac
  done
}

# parse args
CONFIRM=false
while [[ $# -gt 0 ]]; do
  case "$1" in
    --module)
      MODULE="$2"
      shift 2
      ;;
    --module=*)
      MODULE="${1#*=}"
      shift
      ;;
    --dir)
      WORKING_DIR="$2"
      shift 2
      ;;
    --dir=*)
      WORKING_DIR="${1#*=}"
      shift
      ;;
    -y|--yes)
      CONFIRM=true
      shift
      ;;
    -h|--help)
      show_help
      exit 0
      ;;
    *)
      msg_err "Unknown option '$1'"
      show_help
      exit 1
      ;;
  esac
done

# validate
if [[ -z "${MODULE}" ]]; then
  msg_err "'--module' is required"
  show_help
  exit 1
fi

# set defaults
if [[ -z "${WORKING_DIR}" ]]; then
  WORKING_DIR="$(pwd)"
fi

# normalize path (remove trailing slash)
WORKING_DIR="${WORKING_DIR%/}"

#infer project dir from go module
PROJECT_DIR=$(basename "$MODULE")
# If it ends with /vN (where N is a number), remove that suffix
if [[ "$PROJECT_DIR" =~ ^v[0-9]+$ ]]; then
  # If last part is version, get the previous segment
  PROJECT_DIR=$(basename "$(dirname "$MODULE")")
fi

# confirm
msg_info "Go module:          ${MODULE}"
msg_info "Working directory:  ${WORKING_DIR}"
msg_info "Project directory:  ${PROJECT_DIR}"
if [[ ${CONFIRM} = false ]]; then
  prompt_yes_no
fi

msg_ongoing "Switching to working directory '${WORKING_DIR}'..."
cd ${WORKING_DIR}
msg_done

msg_ongoing "Creating project directory '${PROJECT_DIR}' in '${WORKING_DIR}'..."
if [[ -d "${PROJECT_DIR}" ]]; then
  msg_err "A directory with the name '${PROJECT_DIR}' already exists in '${WORKING_DIR}'"
  exit 1
fi
mkdir ${PROJECT_DIR}
msg_done

# NOTE: it is very important that you change dir to project dir
# safety gates are implemented to halt operation if it detects that you are not in the correct dir
cd ${PROJECT_DIR}

# get source
msg_ongoing "Copying files from '${CATALYST_BASE}' to '$(pwd)'..."
[[ $(basename "$(pwd)") == "${PROJECT_DIR}" ]] || { msg_err "Not in '${PROJECT_DIR}', exiting!"; exit 1; } # NOTE: safety gate
curl --silent --show-error --location "${CATALYST_BASE}/archive/${CATALYST_REF}.tar.gz" | tar --extract --gzip --strip-components=1
msg_done

# remove dirs and files
msg_ongoing "Cleaning up..."
[[ $(basename "$(pwd)") == "${PROJECT_DIR}" ]] || { msg_err "Not in '${PROJECT_DIR}', exiting!"; exit 1; } # NOTE: safety gate
REMOVE_LIST=(
  ".github"
  "docs/img/*.drawio.*"
  "LICENSE"
  "README.md"
  "NOTES.md"
  "new.sh"
)
for item in "${REMOVE_LIST[@]}"; do
  rm -rf ${item} 2>/dev/null || true
  msg_success "Removed, '${item}'"
done

# replace following
# github.com/kosatnkn/catalyst/v3 with <new_module_name>
msg_ongoing "Updating go module and references..."
[[ $(basename "$(pwd)") == "${PROJECT_DIR}" ]] || { msg_err "Not in '${PROJECT_DIR}', exiting!"; exit 1; } # NOTE: safety gate
file -i $(find . -type f ! -path "*/.git/*") | grep "text/" | cut -d: -f1 | while read -r file; do # NOTE: checks whether the file contains text
  before=$(md5sum ${file})
  sed -i -e "s#${CATALYST_MODULE}#${MODULE}#g" "${file}"
  after=$(md5sum ${file})
  [[ "${before}" != "${after}" ]] && msg_success "Updated, '${file}'"
done

# create necessary files
msg_ongoing "Creating 'README.md'..."
[[ $(basename "$(pwd)") == "${PROJECT_DIR}" ]] || { msg_err "Not in '${PROJECT_DIR}', exiting!"; exit 1; } # NOTE: safety gate
echo -e "# ${PROJECT_DIR}

TODO: Add content

---
Powered by [https://github.com/kosatnkn/catalyst](kosatnkn/catalyst) (${CATALYST_REF})
" > README.md
msg_done

msg_complete
