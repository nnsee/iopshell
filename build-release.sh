#!/bin/sh

OUT="./build/iopshell"
mkdir build/

# linux specific
for arch in {386,amd64,arm,arm64,mips,mipsle,mips64,mips64le,riscv64}; do
    echo "linux/$arch"
    GOOS=linux GOARCH=$arch go build -ldflags="-s -w" -o "$OUT.$arch" iopshell.go
done

# windows specific
for arch in {386,amd64,arm}; do
    echo "windows/$arch"
    GOOS=windows GOARCH=$arch go build -ldflags="-s -w" -o "$OUT.$arch.exe" iopshell.go
done
