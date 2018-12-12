#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

CWD=$(pwd)
PROJECT_DIR=$CWD/..
BUILD_DIR=$PROJECT_DIR/build
CLIENT_DIR=$PROJECT_DIR/client
SERVER_DIR=$PROJECT_DIR/server
IGNORE_DIR=vendor,build

ProjectInfo() {
    if [[ "$(uname)" == "Linux" ]]; then
        apt-get update && apt-get install -y cloc
    elif [[ "$(uname)" == "Darwin" ]]; then
        command -v brew >/dev/null 2>&1 || { echo >&2 "Installing Homebrew Now"; \
            /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"; }
        
        brew list cloc || brew install cloc
    fi

    cloc --exclude-ext=toml --exclude-dir=$IGNORE_DIR $PROJECT_DIR
}

ProjectInfo