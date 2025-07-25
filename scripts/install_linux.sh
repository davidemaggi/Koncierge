#!/bin/bash

# allow specifying different destination directory
DIR="${DIR:-"$HOME/.local/bin"}"

ARCH=$(uname -m)

# prepare the download URL
GITHUB_LATEST_VERSION=$(curl -L -s -H 'Accept: application/json' https://github.com/davidemaggi/koncierge/releases/latest | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
GITHUB_FILE="Koncierge_${GITHUB_LATEST_VERSION//v/}_Linux_${ARCH}.tar.gz"

GITHUB_URL="https://github.com/davidemaggi/Koncierge/releases/download/${GITHUB_LATEST_VERSION}/${GITHUB_FILE}"


# install/update the local binary
curl -L -o koncierge.tar.gz $GITHUB_URL
tar xzvf koncierge.tar.gz koncierge
install -Dm 755 koncierge -t "$DIR"
rm koncierge koncierge.tar.gz