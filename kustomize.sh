#!/bin/sh

TEMP_CWD=$(pwd)
TEMP_GO111MODULE=$GO111MODULE
export GO111MODULE=off
cd ~
go get golang.org/dl/go1.12beta1
$GOPATH/bin/go1.12beta1 download
export GO111MODULE=on
$GOPATH/bin/go1.12beta1 get github.com/FigureTechnologies/kustomize@v2.0.0-figure
export GO111MODULE=$TEMP_GO111MODULE
cd $TEMP_CWD