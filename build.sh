#!/bin/bash
echo "Removing old binaries from ./bin"
rm -rf ./bin/*

echo "Compiling Windows binaries into ./bin/"
GOOS=windows GOARCH=arm64 go build -o ./bin/smartpass-link-gen-windows-arm-64.exe ./src/
GOOS=windows GOARCH=amd64 go build -o ./bin/smartpass-link-gen-windows-x86-64.exe ./src/

echo "Compiling Linux binaries into ./bin/"
GOOS=linux GOARCH=arm64 go build -o ./bin/smartpass-link-gen-linux-arm-64 ./src/
GOOS=linux GOARCH=amd64 go build -o ./bin/smartpass-link-gen-linux-x86-64 ./src/

echo "Compiling OSX binaries into ./bin/"
GOOS=darwin GOARCH=arm64 go build -o ./bin/smartpass-link-gen-osx-apple-silicon ./src/
GOOS=darwin GOARCH=amd64 go build -o ./bin/smartpass-link-gen-osx-intel-x86-64 ./src/

echo "Done"