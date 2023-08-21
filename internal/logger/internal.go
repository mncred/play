package logger

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// renderFields converts fields to string
func (l Logger) renderFields() string {
	pairs := make([]string, 0, len(l.fields))
	for i := range l.fields {
		str := fmt.Sprintf("%v", l.fields[i].Value)
		val := strconv.Quote(str)
		pairs = append(pairs, fmt.Sprintf("%v=%v", l.fields[i].Key, val))
	}
	return strings.Join(pairs, " ")
}

// renderLine renders log line
func (l Logger) renderLine(level Level, message string) string {
	lvl := ""
	switch level {
	case 0:
		lvl = "DBG"
	case 1:
		lvl = "INF"
	case 2:
		lvl = "WRN"
	case 3:
		lvl = "ERR"
	case 4:
		lvl = "FTL"
	case 5:
		lvl = "TRC"
	case 6:
		lvl = "PRN"
	}

	return strings.TrimSpace(fmt.Sprintf("%s\n", strings.Join([]string{
		time.Now().Format("15:04:05"),
		lvl,
		l.renderFields(),
		message,
	}, " | ")))
}

// getField returns field value
func (l Logger) getField(field string) any {
	for i := range l.fields {
		if l.fields[i].Key == field {
			return l.fields[i].Value
		}
	}
	return nil
}
