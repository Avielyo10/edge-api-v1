#!/bin/bash
set -e

# check OS type
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    *)          machine="UNKNOWN:${unameOut}"
esac

# pprint message in bold blue
function pprint() {
    echo "\033[1;34m$1\033[0m"
}

# pprint error message in bold red
function perror() {
    echo "\033[1;31m$1\033[0m"
}

# check if go installed
if ! [ -x "$(command -v go)" ]; then
    pprint "go is not installed. Installing..."
    if [ "${machine}" == "Mac" ]; then
        brew install go
    elif [ "${machine}" == "Linux" ]; then
        yum install -y golang
    else
        perror "go is not installed. Please install go from https://golang.org/dl/ and add it to your PATH"
        exit 1
    fi
fi

# check if git installed
if ! [ -x "$(command -v git)" ]; then
    pprint "git is not installed. Installing..."
    if [ "${machine}" == "Mac" ]; then
        brew install git
    elif [ "${machine}" == "Linux" ]; then
        yum install -y git
    else
        perror "git is not installed. Please install git from https://git-scm.com/downloads/ and add it to your PATH"
        exit 1
    fi
fi

# check if oapi-codegen is installed
if ! command -v oapi-codegen > /dev/null; then
    pprint "oapi-codegen is not installed. Installing..."
    git clone https://github.com/deepmap/oapi-codegen.git
    cd oapi-codegen/cmd/oapi-codegen
    # build oapi-codegen
    GO111MODULE=on go build -o oapi-codegen
    mkdir -p ~/go/bin && mv oapi-codegen ~/go/bin/
    cd -
    rm -rf oapi-codegen
fi

# check if buf is installed
if ! command -v buf > /dev/null; then
    pprint "buf is not installed. Installing..."
    if [ "${machine}" == "Mac" ]; then
        brew install bufbuild/buf/buf
    elif [ "${machine}" == "Linux" ]; then
        # get the binary to ~/bin
        curl -sL https://github.com/bufbuild/buf/releases/download/v1.1.0/buf-Linux-x86_64 -o ~/bin/buf

        if [ "${SHELL}" == "/bin/bash" ]; then
            buf completion bash
        elif [ "${SHELL}" == "/bin/zsh" ]; then
            buf completion zsh
        else
            pprint "could not install buf completion"
        fi
    else
        perror "buf is not installed. Please install buf from https://docs.buf.build/installation and add it to your PATH"
        exit 1
    fi
fi

# check if golint is installed
if ! command -v golangci-lint > /dev/null; then
    pprint "golangci-lint is not installed. Installing..."
    if [ "$machine" == "Mac" ]; then
        brew install golangci-lint && brew upgrade golangci-lint
    elif [ "$machine" == "Linux" ]; then
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.44.2
    else
        perror "golangci-lint is not installed. Please install golangci-lint from https://golangci-lint.run/usage/install/ and add it to your PATH"
   fi
fi

# check if gotests is installed
if ! command -v gotests > /dev/null; then
    pprint "gotests is not installed. Installing..."
    go get -u github.com/cweill/gotests/...
fi
