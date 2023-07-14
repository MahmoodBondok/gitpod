#!/usr/bin/env bash

set -euo pipefail

# shellcheck disable=SC2155
export PATH="$PATH:$HOME/bin"

mkdir "$HOME/bin"

gcloud auth activate-service-account --key-file "${PREVIEW_ENV_DEV_SA_KEY_PATH}"

leeway run dev/preview/previewctl:download

previewctl get-credentials --gcp-service-account "${PREVIEW_ENV_DEV_SA_KEY_PATH}"

export TF_INPUT=0
export TF_IN_AUTOMATION=true
TF_VAR_preview_name="$(previewctl get-name --branch "${INPUT_NAME}")"
export TF_VAR_preview_name

leeway run dev/preview:delete-preview
