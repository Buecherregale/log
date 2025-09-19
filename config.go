package log

import "time"

type LogConfig struct {
	Level 								LogLevel
	Timeformat						string
	SerializationStrategy	SerilizationStrategy
	TargetMode 						Target
	Logfile								string
}

type instances struct {
	printer			printer
	serializer 	serializer
}

var config *LogConfig = &LogConfig{
	Level: LEVEL_INFO,
	Timeformat: time.RFC3339,
	SerializationStrategy: SIMPLE,
	TargetMode: STDOUT,
	Logfile: "",
} 
var singletons *instances = &instances{
	printer: &stdwriter{},
	serializer: &simpleSerializer{},
}

