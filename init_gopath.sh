#!/usr/bin/env bash

echo `pwd`
WORKSPACE=`pwd`
echo "export GOPATH=\"${WORKSPACE}\"" >> /.zshrc
source ~/.zshrc
