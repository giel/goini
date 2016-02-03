#!/usr/bin/env bash

# Format, build, test and install script for ini. 
# Meant to run from this directory.
# NOTE: Temporarily sets GOPATH variable !!!

errorFound=""

GOPATH="$PWD"

function build {
    echo "Building:" $1
    cd $1
    echo "Formatting..."
    go fmt
    echo "Testing..."
    go test
    
    if [ "$?" = "0" ]; then
        echo "Installing package..."
        go install
    else
       echo "*** Package" $1 "not installed ***"
       errorFound=$1
    fi
    cd ..
    echo "Done building."
}  

build src/ini

if [ "$errorFound" = "" ]; then
    echo "No errors found"
else
    echo "*** Errors found in test, ini not build ***"
fi
echo "Build done."

