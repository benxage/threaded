#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

CWD=$(pwd)
PROJECT_DIR=$CWD/..
BUILD_DIR=$PROJECT_DIR/build
CLIENT_DIR=$PROJECT_DIR/client
SERVER_DIR=$PROJECT_DIR/server
SERVER=server.go

IGNORE_DIR=vendor,build

EnsureDep() {
    echo "checking dependencies..."

    if [[ "$(uname)" == "Linux" ]]; then
        apt-get update && apt-get install -y cloc dep
    elif [[ "$(uname)" == "Darwin" ]]; then
        command -v brew >/dev/null 2>&1 || { echo >&2 "Installing Homebrew Now"; \
            /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"; }
        
        brew list cloc || brew install cloc
        brew list dep || brew install dep
    fi

    pushd $SERVER_DIR
    dep ensure
    popd
}

ProjectInfo() {
    cloc --exclude-ext=toml --exclude-dir=$IGNORE_DIR $PROJECT_DIR
}

Launch() {
    pushd $SERVER_DIR
    go run $SERVER
    popd
}

LaunchBackground() {
    echo "LaunchBackground"
    pushd $SERVER_DIR
    nohup go run $SERVER &
    popd
}

# always ensure dependency
EnsureDep

if [[ $# -eq 0 ]]; then
    Launch
elif [[ $# -eq 1 ]]; then
    if [[ "$1" == "launch" ]]; then
        Launch
    elif [[ "$1" == "info" ]]; then
        ProjectInfo
        exit
    else
        echo "unknown command"
        exit 1
    fi
elif [[ $# -eq 2 ]]; then
    if [[ "$1" == "launch" && "$2" == "background" ]]; then
        LaunchBackground
    else
        echo "unknown command"
        exit 1
    fi
else
    echo "too many arguments"
    exit 1
fi