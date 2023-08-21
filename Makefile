NAME=mncplay
BIN_DIR=./build/bin/
WEBVIEW2="embed"

# Uncomment it if you need UPX compressed.
# It cases segmentation fault in my case.
# UPX="--upx --upxflags -1"

# Uncomment it if you need NSIS installer.
# NSIS="--nsis"

BUILD_TIME=$(shell date '+%d.%m.%y %H:%M:%S %z')
BUILD_ORIGIN=$(shell git remote get-url origin)
BUILD_COMMIT=$(shell git rev-parse HEAD)
BUILD_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT_AUTHOR=$(shell git log -1 --pretty=format:'%an')
BUILD_COMMIT_EMAIL=$(shell git log -1 --pretty=format:'%ae')
BUILD_TAG=$(shell git tag --points-at HEAD)
BUILD_COMPILER=$(shell go version)
BUILD_WAILS=$(shell wails version | head -n 1)
LDFLAGS='-X "github.com/mncred/mncplay/internal/build.BuildTime=${BUILD_TIME}" -X "github.com/mncred/mncplay/internal/build.BuildOrigin=${BUILD_ORIGIN}" -X "github.com/mncred/mncplay/internal/build.BuildCommit=${BUILD_COMMIT}" -X "github.com/mncred/mncplay/internal/build.BuildBranch=${BUILD_BRANCH}" -X "github.com/mncred/mncplay/internal/build.BuildCommitAuthor=${BUILD_COMMIT_AUTHOR}" -X "github.com/mncred/mncplay/internal/build.BuildCommitEmail=${BUILD_COMMIT_EMAIL}" -X "github.com/mncred/mncplay/internal/build.BuildTag=${BUILD_TAG}" -X "github.com/mncred/mncplay/internal/build.BuildCompiler=${BUILD_COMPILER}" -X "github.com/mncred/mncplay/internal/build.BuildWails=${BUILD_WAILS}"'

dev:
	wails dev --ldflags ${LDFLAGS}
.PHONY: dev

build_all: build_windows build_darwin build_linux
.PHONY: build_all

build_windows: build_windows_amd64 build_windows_386 build_windows_arm64
.PHONY: build_windows

build_darwin: build_darwin_amd64 build_darwin_arm64 build_darwin_universal
.PHONY: build_darwin

build_linux: build_linux_amd64 build_linux_arm build_linux_arm64
.PHONY: build_linux

build_windows_386:
	wails build ${NSIS} --ldflags ${LDFLAGS} ${UPX} --webview2 ${WEBVIEW2} --platform windows/386 -o ${NAME}.exe
	mv ${BIN_DIR}${NAME}.exe ${BIN_DIR}${NAME}-windows-386.exe
.PHONY: build_windows_386

build_windows_amd64:
	wails build ${NSIS} --ldflags ${LDFLAGS} ${UPX} --webview2 ${WEBVIEW2} --platform windows/amd64 -o ${NAME}.exe
	mv ${BIN_DIR}${NAME}.exe ${BIN_DIR}${NAME}-windows-amd64.exe
.PHONY: build_windows_amd64

build_windows_arm64:
# UPX can't compress windows/arm64
	wails build ${NSIS} --ldflags ${LDFLAGS} --webview2 ${WEBVIEW2} --platform windows/arm64 -o ${NAME}.exe
	mv ${BIN_DIR}${NAME}.exe ${BIN_DIR}${NAME}-windows-arm64.exe
.PHONY: build_windows_arm64

build_darwin_amd64:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${WEBVIEW2} --platform darwin/amd64 -o ${NAME}
	mv ${BIN_DIR}${NAME}.app ${BIN_DIR}${NAME}-darwin-amd64.app
.PHONY: build_darwin_amd64

build_darwin_arm64:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${WEBVIEW2} --platform darwin/arm64 -o ${NAME}
	mv ${BIN_DIR}${NAME}.app ${BIN_DIR}${NAME}-darwin-arm64.app
.PHONY: build_darwin_arm64

build_darwin_universal:
# UPX can't compress darwin/universal
	wails build --ldflags ${LDFLAGS} --webview2 ${WEBVIEW2} --platform darwin/universal -o ${NAME}
	mv ${BIN_DIR}${NAME}.app ${BIN_DIR}${NAME}-darwin-universal.app
.PHONY: build_darwin_universal

build_linux_amd64:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${WEBVIEW2} --platform linux/amd64 -o ${NAME}
	mv ${BIN_DIR}${NAME} ${BIN_DIR}${NAME}-linux-amd64
.PHONY: build_linux_amd64

build_linux_arm:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${WEBVIEW2} --platform linux/arm -o ${NAME}
	mv ${BIN_DIR}${NAME} ${BIN_DIR}${NAME}-linux-arm
.PHONY: build_linux_arm

build_linux_arm64:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${WEBVIEW2} --platform linux/arm64 -o ${NAME}
	mv ${BIN_DIR}${NAME} ${BIN_DIR}${NAME}-linux-arm64
.PHONY: build_linux_arm64
