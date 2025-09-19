package log

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SerilizationStrategy int

const (
	JSON		SerilizationStrategy = 1 // if used with Target = FILE the resulting file will NOT be valid json. 
	SIMPLE	SerilizationStrategy = 2
)

type serializer interface {
	serialize(entry LogEntry) (string, error)
}

type jsonSerializer struct {}

func (serializer *jsonSerializer) serialize(entry LogEntry) (string, error) {
	b, err := json.MarshalIndent(entry, "", "    ")
	if err != nil {
		return "", err
	}
	return string(b) + "\n", nil
}

const simple_formatting string = "%s [%-5s] | %s"

type simpleSerializer struct {}

func (serializer *simpleSerializer) serialize(entry LogEntry) (string, error) {
	return fmt.Sprintf(simple_formatting, entry.Timestamp, entry.Level.Name(), entry.Message), nil
}

func getSerializer(strategy SerilizationStrategy) (serializer, error) {
	switch strategy {
	case JSON:
		return &jsonSerializer{}, nil
	case SIMPLE:
		return &simpleSerializer{}, nil
	default:
		return nil, errors.New("unknown serialization strategy")
	}
}

