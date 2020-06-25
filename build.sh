#!/bin/bash
echo "Removing old binaries from ./bin"
rm -rf ./bin/*

echo "Compiling Windows binaries into ./bin/"
GOOS=windows GOARCH=386 go build -o ./bin/encrypted-link-generator-windows-32-bit.exe ./src/
GOOS=windows GOARCH=amd64 go build -o ./bin/encrypted-link-generator-windows-64-bit.exe ./src/

echo "Compiling Linux binaries into ./bin/"
GOOS=linux GOARCH=386 go build -o ./bin/encrypted-link-generator-linux-32-bit ./src/
GOOS=linux GOARCH=amd64 go build -o ./bin/encrypted-link-generator-linux-64-bit ./src/

echo "Compiling OSX binaries into ./bin/"
GOOS=darwin GOARCH=386 go build -o ./bin/encrypted-link-generator-osx-32-bit ./src/
GOOS=darwin GOARCH=amd64 go build -o ./bin/encrypted-link-generator-osx-64-bit ./src/

echo "Done"