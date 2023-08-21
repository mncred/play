package logger

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"text/template"
	"time"

	"github.com/fatih/color"
)

// Logger provides an instance of the app logging system
type Logger struct {
	fields []Field
	level  Level
	// middleware modifies the logs into output and do whatever other things
	middleware func(level Level, line string) string
	// contains log file mask to copy logs
	output string
}

// New creates new logger
func New() Logger {
	return Logger{
		fields: []Field{},
		level:  LevelInfo,
		middleware: func(level Level, line string) string {
			return line
		},
	}
}

// WithField returns new instance with appended or replaced field
func (l Logger) WithField(field string, value any) Logger {
	// make deep copy of fields to manipulate as a new instance of logger
	fields := make([]Field, len(l.fields))
	copy(fields, l.fields)

	// search if field already exists and save index
	fieldIndex := -1
	for i := range fields {
		if fields[i].Key == field {
			fieldIndex = i
			break
		}
	}

	// check if the field exist - set new value
	if fieldIndex > -1 {
		fields[fieldIndex].Value = value
	} else {
		fields = append(fields, Field{Key: field, Value: value})
	}

	l.fields = fields
	return l
}

// WithName returns logger with specific name.
// Unlike WithField it will append new names separated with slashes.
func (l Logger) WithName(name string) Logger {
	line := ""
	if l.getField(FieldKeyName) != nil {
		line = fmt.Sprintf("%v/%s", l.getField(FieldKeyName), name)
	} else {
		line = name
	}
	return l.WithField(FieldKeyName, line)
}

// WithLogger returns new logger instance with specific log level.
func (l Logger) WithLevel(level Level) Logger {
	l.level = level
	return l
}

// WithMiddleware returns new logger instance with specific middleware.
func (l Logger) WithMiddleware(middleware func(level Level, line string) string) Logger {
	l.middleware = middleware
	return l
}

// WithOutput returns new logger instance with specific output file by template.
func (l Logger) WithOutput(output string) Logger {
	// reset output if empty
	if len(output) == 0 {
		l.output = ""
		return l
	}
	// template data
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("[LOGGER] unable to get user info, logs won't be saved: %v\n", err)
		return l
	}
	data := map[string]any{
		"stamp":  time.Now().Unix(),
		"home":   usr.HomeDir,
		"user":   usr.Username,
		"year":   time.Now().Format("06"),
		"month":  time.Now().Format("01"),
		"day":    time.Now().Format("02"),
		"hour":   time.Now().Format("15"),
		"minute": time.Now().Format("04"),
		"second": time.Now().Format("05"),
	}
	// template this output
	tpl, err := template.New("output").Parse(output)
	if err != nil {
		fmt.Printf("[LOGGER] unable to parse output template, logs won't be saved: %s: %v\n", output, err)
		return l
	}
	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, data); err != nil {
		fmt.Printf("[LOGGER] unable to execute output template, logs won't be saved: %s: %v\n", output, err)
		return l
	}
	filename := buffer.String()
	l.output = filename
	return l
}

// Pointer returns pointer to the logger instance.
func (l Logger) Pointer() *Logger {
	return &l
}

// Trace level logging.
func (l Logger) Trace(message string) {
	if l.level <= LevelTrace {
		line := l.middleware(LevelTrace, l.renderLine(LevelTrace, message))
		color.HiBlack("%s", line)
	}
}

// Debug level logging.
func (l Logger) Debug(message string) {
	if l.level <= LevelDebug {
		line := l.middleware(LevelDebug, l.renderLine(LevelDebug, message))
		color.WhiteString("%s", line)
		outputWrite(l.output, line)
	}
}

// Info level logging.
func (l Logger) Info(message string) {
	if l.level <= LevelInfo {
		line := l.middleware(LevelInfo, l.renderLine(LevelInfo, message))
		color.HiBlue("%s", line)
		outputWrite(l.output, line)
	}
}

// Warning level logging.
func (l Logger) Warning(message string) {
	if l.level <= LevelWarning {
		line := l.middleware(LevelWarning, l.renderLine(LevelWarning, message))
		color.HiYellow("%s", line)
		outputWrite(l.output, line)
	}
}

// Error level logging.
func (l Logger) Error(message string) {
	if l.level <= LevelError {
		line := l.middleware(LevelError, l.renderLine(LevelError, message))
		color.HiRed("%s", line)
		outputWrite(l.output, line)
	}
}

// Fatal level logging.
func (l Logger) Fatal(message string) {
	if l.level <= LevelFatal {
		line := l.middleware(LevelFatal, l.renderLine(LevelFatal, message))
		color.HiRed("%s", line)
		outputWrite(l.output, line)
	}
	os.Exit(1)
}

// Print level logging.
func (l Logger) Print(message string) {
	if l.level <= LevelPrint {
		line := l.middleware(LevelPrint, l.renderLine(LevelPrint, message))
		fmt.Println(line)
		outputWrite(l.output, line)
	}
}
