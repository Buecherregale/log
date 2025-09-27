package log

import (
	"fmt"
	"strings"
	"time"
)

type LogEntry struct {
	Timestamp	string			`json:"timestamp"`
	Level			LogLevel		`json:"log_level"`
	Message		string			`json:"message"`
}

type LogLevel int

const (
	LEVEL_ERROR	LogLevel = 1
	LEVEL_WARN	LogLevel = 2
	LEVEL_INFO  LogLevel = 3
	LEVEL_DEBUG LogLevel = 4
)

func (level LogLevel) Name() string {
	switch level {
	case LEVEL_ERROR:
		return "ERROR"
	case LEVEL_WARN:
		return "WARN"
	case LEVEL_INFO:
		return "INFO"
	case LEVEL_DEBUG:
		return "DEBUG"
	}
	return "UNKNOWN"
}

func parseLevel(lvl string) LogLevel {
	upper := strings.ToUpper(lvl)
	switch upper {
	case "ERROR":
		return LEVEL_ERROR
	case "WARN":
		return LEVEL_WARN
	case "INFO":
		return LEVEL_INFO
	case "DEBUG":
		return LEVEL_DEBUG
	}
	return -1
}

func Init() {
	configureGlobalState(buildDefaultConfig())
}
func Configure(config LogConfig) {
	configureGlobalState(&config)
}

func Logf(level LogLevel, message string, args... any) {
	if level > activeConfig.Level {
		return
	}
	s := SLogf(level, message, args...)
	err := activePrinter.write(s)
	if err != nil {
		panic(err)
	}
}
func Logln(level LogLevel, args... any) {
	if level > activeConfig.Level {
		return
	}
	s := SLogln(level, args...)
	err := activePrinter.write(s)
	if err != nil {
		panic(err)
	}
}

func SLogf(level LogLevel, message string, args... any) string {
	entry := LogEntry {
		Timestamp: time.Now().Local().Format(activeConfig.Timeformat),
		Level: level,
		Message: fmt.Sprintf(message, args...),
	}
	s, err := activeSerializer.serialize(entry)
	if err != nil {
		panic(fmt.Sprintf("could not serialize log entry: %v\n", err))
	}
	return s
}
func SLogln(level LogLevel, args... any) string {
	entry := LogEntry {
		Timestamp: time.Now().Local().Format(activeConfig.Timeformat),
		Level: level,
		Message: fmt.Sprintln(args...),
	}
	s, err := activeSerializer.serialize(entry)
	if err != nil {
		panic(fmt.Sprintf("could not serialize log entry: %v\n", err))
	}
	return s
}

func Fatalf(message string, args... any) {
	panic(SLogf(LEVEL_ERROR, message, args...))
}
func Fatalln(args... any) {
	panic(SLogln(LEVEL_ERROR, args...))
}
func Fatal(err error) {
	panic(SLogf(LEVEL_ERROR, "%v\n", err)) 
}
func Errorf(message string, args... any) {
	Logf(LEVEL_ERROR, message, args...)
}
func Errorln(args... any) {
	Logln(LEVEL_ERROR, args...)
}
func Warnf(message string, args... any) {
	Logf(LEVEL_WARN, message, args...)
}
func Warnln(args... any) {
	Logln(LEVEL_WARN, args...)
}
func Infof(message string, args... any) {
	Logf(LEVEL_INFO, message, args...)
}
func Infoln(args... any) {
	Logln(LEVEL_INFO, args...)
}
func Debugf(message string, args... any) {
	Logf(LEVEL_DEBUG, message, args...)
}
func Debugln(args... any) {
	Logln(LEVEL_DEBUG, args...)
}

