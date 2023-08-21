package logger

import "strings"

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelTrace
	LevelPrint
)

// ParseLevel returns log level from string.
func ParseLevel(level string) Level {
	switch strings.ToLower(level) {
	case "debug", "dbg", "d":
		return LevelDebug
	case "info", "inf", "i":
		return LevelInfo
	case "warning", "warn", "war", "wrn", "w":
		return LevelWarning
	case "error", "err", "e":
		return LevelError
	case "fatal", "fat", "ftl", "f":
		return LevelFatal
	case "trace", "trc", "t":
		return LevelTrace
	case "print", "prn", "p":
		return LevelPrint
	default:
		return LevelInfo
	}
}
