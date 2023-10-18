#!/usr/bin/env bash

set -euo pipefail

export SCRIPT_DIR=$(dirname "${BASH_SOURCE[0]}")
cd "$SCRIPT_DIR" && go run ./generate/main.go > objects_tmp.go && mv objects_tmp.go objects.go
