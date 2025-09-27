package log

import (
	"os"
	"time"
)

type LogConfig struct {
	Level 								LogLevel
	Timeformat						string
	SerializationStrategy	SerilizationStrategy
	TargetMode 						Target
	Logfile								string
}

var activeConfig *LogConfig
var activePrinter printer
var activeSerializer serializer

// builds a config with env and default values
func buildDefaultConfig() *LogConfig {
	cfg := LogConfig {
		Timeformat: time.RFC3339,
		SerializationStrategy: SERIALIZATION_SIMPLE,
		TargetMode: TARGET_STDOUT,
	}

	file := os.Getenv("GO_LOG_FILE")
	cfg.Logfile = file 
 
	level, set := os.LookupEnv("GO_LOG_LEVEL")
	if set {
		cfg.Level = parseLevel(level)
	} else {
		cfg.Level = LEVEL_INFO
	}

	return &cfg
}

// sets the active config, printer and serializer
func configureGlobalState(config *LogConfig) {
	printer, err := parsePrinter(config.TargetMode)
	if err != nil {
		panic(err)
	}
	serializer, err := parseSerializer(config.SerializationStrategy)
	if err != nil {
		panic(err)
	}
	activePrinter = printer
	activeSerializer = serializer
	activeConfig = config
}

