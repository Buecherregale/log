package log

import (
	"fmt"
	"time"
)

type LogEntry struct {
	Timestamp	string			`json:"timestamp"`
	Level			LogLevel		`json:"log_level"`
	Message		string			`json:"message"`
}

type LogLevel int

const (
	ERROR	LogLevel = 1
	WARN	LogLevel = 2
	INFO  LogLevel = 3
	DEBUG LogLevel = 4
)

func (level LogLevel) Name() string {
	switch level {
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	}
	return "UNKNOWN"
}

func Init(logConfig LogConfig) {
	config = &logConfig
	printer, err := getPrinter(logConfig.TargetMode)
	if err != nil {
		panic(err)
	}
	serializer, err := getSerializer(logConfig.SerializationStrategy)
	if err != nil {
		panic(err)
	}
	singletons = &instances{
		printer: printer,
		serializer: serializer,
	}
}

func Logf(level LogLevel, message string, args... any) {
	if level > config.Level {
		return
	}
	s := SLogf(level, message, args...)
	err := singletons.printer.write(s)
	if err != nil {
		panic(err)
	}
}

func SLogf(level LogLevel, message string, args... any) string {
	entry := LogEntry {
		Timestamp: time.Now().Local().Format(config.Timeformat),
		Level: level,
		Message: fmt.Sprintf(message, args...),
	}
	s, err := singletons.serializer.serialize(entry)
	if err != nil {
		panic(fmt.Sprintf("could not serialize log entry: %v\n", err))
	}
	return s
}
func Logln(level LogLevel, message string) {
	Logf(level, message + "\n")
}
func Fatalf(message string, args... any) {
	panic(SLogf(ERROR, message, args...))
}
func Fatalln(message string) {
	panic(SLogf(ERROR, message + "\n"))
}
func Errorf(message string, args... any) {
	Logf(ERROR, message, args...)
}
func Errorln(message string) {
	Logln(ERROR, message)
}
func Warnf(message string, args... any) {
	Logf(WARN, message, args...)
}
func Warnln(message string) {
	Logln(WARN, message)
}
func Infof(message string, args... any) {
	Logf(INFO, message, args...)
}
func Infoln(message string) {
	Logln(INFO, message)
}
func Debugf(message string, args... any) {
	Logf(DEBUG, message, args...)
}
func Debugln(message string) {
	Logln(DEBUG, message)
}

