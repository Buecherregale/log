package log

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SerilizationStrategy int

const (
	SERIALIZATION_JSON		SerilizationStrategy = 1 // if used with Target = FILE the resulting file will NOT be valid json. 
	SERIALIZATION_SIMPLE	SerilizationStrategy = 2
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

func parseSerializer(strategy SerilizationStrategy) (serializer, error) {
	switch strategy {
	case SERIALIZATION_JSON:
		return &jsonSerializer{}, nil
	case SERIALIZATION_SIMPLE:
		return &simpleSerializer{}, nil
	default:
		return nil, errors.New("unknown serialization strategy")
	}
}

