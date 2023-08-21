package system

import (
	"runtime"
	"strings"
)

type Info struct{}

// Platform returns GoLang's runtime.GOOS value.
// It might be something like: windows, linux, darwin, etc.
func (info *Info) Platform() string {
	return runtime.GOOS
}

func (info *Info) PlatformIs(name string) bool {
	name = strings.ToLower(name)
	platform := runtime.GOOS

	switch platform {
	case "windows":
		switch name {
		case "win", "windows":
			return true
		}
	case "linux":
		switch name {
		case "unix", "linux":
			return true
		}
	case "darwin":
		switch name {
		case "mac", "osx", "macos", "darwin":
			return true
		}
	}
	return false
}

func (info *Info) PlatformIsWindows() bool {
	return info.PlatformIs("windows")
}

func (info *Info) PlatformIsLinux() bool {
	return info.PlatformIs("linux")
}

func (info *Info) PlatformIsDarwin() bool {
	return info.PlatformIs("darwin")
}

// Arch returns GoLang's runtime.GOARCH value.
// It might be something like: 386, amd64, arm, arm64, etc.
func (info *Info) Arch() string {
	return runtime.GOARCH
}

func (info *Info) ArchIs386() bool {
	return runtime.GOARCH == "386"
}

func (info *Info) ArchIsAmd64() bool {
	return runtime.GOARCH == "amd64"
}

func (info *Info) ArchIsArm() bool {
	return runtime.GOARCH == "arm"
}

func (info *Info) ArchIsArm64() bool {
	return runtime.GOARCH == "arm64"
}

func (info *Info) ArchIs32Bit() bool {
	return runtime.GOARCH == "arm" || runtime.GOARCH == "386"
}

func (info *Info) ArchIs64Bit() bool {
	return runtime.GOARCH == "arm64" || runtime.GOARCH == "amd64"
}
