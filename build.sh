#!/usr/bin/env bash
set -ex

GOOS=linux   GOARCH=386   go build -o builds/lgtm_linux_386
GOOS=linux   GOARCH=amd64 go build -o builds/lgtm_linux_amd64
GOOS=linux   GOARCH=arm   go build -o builds/lgtm_linux_arm7
GOOS=linux   GOARCH=arm64 go build -o builds/lgtm_linux_arm64
GOOS=darwin  GOARCH=amd64 go build -o builds/lgtm_mac_amd64
GOOS=windows GOARCH=386   go build -o builds/lgtm_windows_386.exe
GOOS=windows GOARCH=amd64 go build -o builds/lgtm_windows_amd64.exe
GOOS=freebsd GOARCH=386   go build -o builds/lgtm_freebsd_386
GOOS=freebsd GOARCH=amd64 go build -o builds/lgtm_freebsd_amd64
GOOS=freebsd GOARCH=arm   go build -o builds/lgtm_freebsd_arm7
GOOS=openbsd GOARCH=386   go build -o builds/lgtm_openbsd_386
GOOS=openbsd GOARCH=amd64 go build -o builds/lgtm_openbsd_amd64
