BUILD_NAME ?= mncplay
BUILD_BIN_DIR ?= ./build/bin/
BUILD_WEBVIEW2 ?= embed

# Uncomment it if you need UPX compressed.
# It cases segmentation fault in my case.
# UPX="--upx --upxflags -1"

# Uncomment it if you need NSIS installer.
# NSIS="--nsis"

# Built-in App Parameters
BUILD_TIME ?= $(shell date '+%d.%m.%y %H:%M:%S %z')
BUILD_ORIGIN ?= $(shell git remote get-url origin)
BUILD_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TAG ?= $(shell git tag --points-at HEAD)
BUILD_COMMIT ?= $(shell git rev-parse HEAD)
BUILD_COMMIT_AUTHOR ?= $(shell git log -1 --pretty=format:'%an')
BUILD_COMMIT_EMAIL ?= $(shell git log -1 --pretty=format:'%ae')
BUILD_COMPILER ?= $(shell go version)
BUILD_WAILS ?= $(shell wails version | head -n 1)
BUILD_OS ?= $(shell go env GOOS)
BUILD_ARCH ?= $(shell go env GOARCH)
BUILD_LOG_LEVEL ?= info
BUILD_LOGS_DIR ?= ./logs
BUILD_APP_ID ?= ${BUILD_NAME}
BUILD_WINDOW_TITLE ?= MNCPlay
BUILD_WINDOW_WIDTH ?= 1000
BUILD_WINDOW_HEIGHT ?= 600
BUILD_WINDOW_WIDTH_MIN ?= 1000
BUILD_WINDOW_WIDTH_MAX ?= 1000
BUILD_WINDOW_WIDTH_MIN ?= 600
BUILD_WINDOW_WIDTH_MAX ?= 600
BUILD_WINDOW_DISABLE_RESIZE ?= true
BUILD_WINDOW_FRAMELESS ?= true
BUILD_WINDOW_DISABLE_MAXIMIZE ?= true
BUILD_WINDOW_DISABLE_MINIMIZE ?= 
BUILD_WINDOW_DISABLE_CLOSE ?= 

LDFLAGS='\
	-X "github.com/mncred/mncplay/internal/build.Time=${BUILD_TIME}" \
	-X "github.com/mncred/mncplay/internal/build.Origin=${BUILD_ORIGIN}" \
	-X "github.com/mncred/mncplay/internal/build.Branch=${BUILD_BRANCH}" \
	-X "github.com/mncred/mncplay/internal/build.Commit=${BUILD_COMMIT}" \
	-X "github.com/mncred/mncplay/internal/build.CommitAuthor=${BUILD_COMMIT_AUTHOR}" \
	-X "github.com/mncred/mncplay/internal/build.CommitEmail=${BUILD_COMMIT_EMAIL}" \
	-X "github.com/mncred/mncplay/internal/build.CommitTag=${BUILD_TAG}" \
\
	-X "github.com/mncred/mncplay/internal/build.Compiler=${BUILD_COMPILER}" \
	-X "github.com/mncred/mncplay/internal/build.Wails=${BUILD_WAILS}" \
	-X "github.com/mncred/mncplay/internal/build.OS=${BUILD_OS}" \
	-X "github.com/mncred/mncplay/internal/build.Arch=${BUILD_ARCH}" \
\
	-X "github.com/mncred/mncplay/internal/build.LogLevel=${BUILD_LOG_LEVEL}" \
	-X "github.com/mncred/mncplay/internal/build.LogsDir=${BUILD_LOGS_DIR}" \
\
	-X "github.com/mncred/mncplay/internal/build.AppId=${BUILD_APP_ID}" \
	-X "github.com/mncred/mncplay/internal/build.WindowTitle=${BUILD_WINDOW_TITLE}" \
	-X "github.com/mncred/mncplay/internal/build.WindowWidth=${BUILD_WINDOW_WIDTH}" \
	-X "github.com/mncred/mncplay/internal/build.WindowHeight=${BUILD_WINDOW_HEIGHT}" \
	-X "github.com/mncred/mncplay/internal/build.WindowWidthMin=${BUILD_WINDOW_WIDTH_MIN}" \
	-X "github.com/mncred/mncplay/internal/build.WindowWidthMax=${BUILD_WINDOW_WIDTH_MAX}" \
	-X "github.com/mncred/mncplay/internal/build.WindowHeightMin=${BUILD_WINDOW_HEIGHT_MIN}" \
	-X "github.com/mncred/mncplay/internal/build.WindowHeightMax=${BUILD_WINDOW_HEIGHT_MAX}" \
	-X "github.com/mncred/mncplay/internal/build.WindowDisableResize=${BUILD_WINDOW_DISABLE_RESIZE}" \
	-X "github.com/mncred/mncplay/internal/build.WindowFrameless=${BUILD_WINDOW_FRAMELESS}" \
	-X "github.com/mncred/mncplay/internal/build.WindowDisableMaximize=${BUILD_WINDOW_DISABLE_MAXIMIZE}" \
	-X "github.com/mncred/mncplay/internal/build.WindowDisableMinimize=${BUILD_WINDOW_DISABLE_MINIMIZE}" \
	-X "github.com/mncred/mncplay/internal/build.WindowDisableClose=${BUILD_WINDOW_DISABLE_CLOSE}" \
'

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
	wails build ${NSIS} --ldflags ${LDFLAGS} ${UPX} --webview2 ${BUILD_WEBVIEW2} --platform windows/386 -o ${BUILD_NAME}.exe
	mv ${BUILD_BIN_DIR}${BUILD_NAME}.exe ${BUILD_BIN_DIR}${BUILD_NAME}-windows-386.exe
.PHONY: build_windows_386

build_windows_amd64:
	wails build ${NSIS} --ldflags ${LDFLAGS} ${UPX} --webview2 ${BUILD_WEBVIEW2} --platform windows/amd64 -o ${BUILD_NAME}.exe
	mv ${BUILD_BIN_DIR}${BUILD_NAME}.exe ${BUILD_BIN_DIR}${BUILD_NAME}-windows-amd64.exe
.PHONY: build_windows_amd64

build_windows_arm64:
# UPX can't compress this platform binaries
	wails build ${NSIS} --ldflags ${LDFLAGS} --webview2 ${BUILD_WEBVIEW2} --platform windows/arm64 -o ${BUILD_NAME}.exe
	mv ${BUILD_BIN_DIR}${BUILD_NAME}.exe ${BUILD_BIN_DIR}${BUILD_NAME}-windows-arm64.exe
.PHONY: build_windows_arm64

build_darwin_amd64:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${BUILD_WEBVIEW2} --platform darwin/amd64 -o ${BUILD_NAME}
	mv ${BUILD_BIN_DIR}${BUILD_NAME}.app ${BUILD_BIN_DIR}${BUILD_NAME}-darwin-amd64.app
.PHONY: build_darwin_amd64

build_darwin_arm64:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${BUILD_WEBVIEW2} --platform darwin/arm64 -o ${BUILD_NAME}
	mv ${BUILD_BIN_DIR}${BUILD_NAME}.app ${BUILD_BIN_DIR}${BUILD_NAME}-darwin-arm64.app
.PHONY: build_darwin_arm64

build_darwin_universal:
# UPX can't compress this platform binaries
	wails build --ldflags ${LDFLAGS} --webview2 ${BUILD_WEBVIEW2} --platform darwin/universal -o ${BUILD_NAME}
	mv ${BUILD_BIN_DIR}${BUILD_NAME}.app ${BUILD_BIN_DIR}${BUILD_NAME}-darwin-universal.app
.PHONY: build_darwin_universal

build_linux_amd64:
	wails build --ldflags ${LDFLAGS} ${UPX} --webview2 ${BUILD_WEBVIEW2} --platform linux/amd64 -o ${BUILD_NAME}
	mv ${BUILD_BIN_DIR}${BUILD_NAME} ${BUILD_BIN_DIR}${BUILD_NAME}-linux-amd64
.PHONY: build_linux_amd64

build_linux_arm:
# UPX can't compress this platform binaries
	wails build --ldflags ${LDFLAGS} --webview2 ${BUILD_WEBVIEW2} --platform linux/arm -o ${BUILD_NAME}
	mv ${BUILD_BIN_DIR}${BUILD_NAME} ${BUILD_BIN_DIR}${BUILD_NAME}-linux-arm
.PHONY: build_linux_arm

build_linux_arm64:
# UPX can't compress this platform binaries
	wails build --ldflags ${LDFLAGS} --webview2 ${BUILD_WEBVIEW2} --platform linux/arm64 -o ${BUILD_NAME}
	mv ${BUILD_BIN_DIR}${BUILD_NAME} ${BUILD_BIN_DIR}${BUILD_NAME}-linux-arm64
.PHONY: build_linux_arm64
