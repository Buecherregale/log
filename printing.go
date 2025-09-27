package log

import (
	"errors"
	"fmt"
	"os"
)

type printer interface {
	write(s string) error
}

type fileWriter struct {
	file 	*os.File
}

func (writer *fileWriter) write(s string) error {
	_, err := writer.file.WriteString(s)
	return err
} 

type stdwriter struct {}

func (writer *stdwriter) write(s string) error {
	fmt.Fprint(os.Stdout, s)
	return nil
}

type Target int

const (
	TARGET_FILE	Target = iota
	TARGET_STDOUT
)

func parsePrinter(target Target) (printer, error) {
	switch target {
	case TARGET_FILE:
		f, err := os.OpenFile(activeConfig.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err 
		}
		return &fileWriter{ file: f }, nil
	case TARGET_STDOUT:
		return &stdwriter{}, nil
	default:
		return nil, errors.New("unknown target")
	}
}

