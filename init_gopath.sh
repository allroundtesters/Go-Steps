#!/usr/bin/env bash

echo `pwd`
WORKSPACE=`pwd`
echo "export GOPATH=\"${WORKSPACE}\"" >> /.zshrc

export GOPATH=$(go env GOPATH)
export PATH=${GOPATH}/bin:$PATH
source ~/.zshrc
