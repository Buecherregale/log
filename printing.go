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
	FILE	Target = iota
	STDOUT
)

func getPrinter(target Target) (printer, error) {
	switch target {
	case FILE:
		f, err := os.OpenFile(config.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err 
		}
		return &fileWriter{ file: f }, nil
	case STDOUT:
		return &stdwriter{}, nil
	default:
		return nil, errors.New("unknown target")
	}
}

