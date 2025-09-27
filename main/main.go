package main

import (
	"time"

	"github.com/Buecherregale/log"
)

func main() {
	config := log.LogConfig {
		Level: log.LEVEL_DEBUG,
		Timeformat: time.RFC3339,
		SerializationStrategy: log.SERIALIZATION_SIMPLE,
		TargetMode: log.TARGET_STDOUT,
		Logfile: "log.txt",
	}
	log.Configure(config)
	log.Infof("Config: %+v\n", config)
	
	print()

	log.Infoln("Now as json")
	config.SerializationStrategy = log.SERIALIZATION_JSON
	log.Configure(config)

	print()

	log.Infoln("Now to file")
	config.TargetMode = log.TARGET_FILE
	log.Configure(config)

	print()

	log.Infoln("Now simply to file")
	config.SerializationStrategy = log.SERIALIZATION_SIMPLE
	log.Configure(config)

	print()
}

func print() {
	log.Debugf("Debug f: %d\n", 0)
	log.Debugln("Debug ln:", 0)
	log.Infof("Info f: %d\n", 0)
	log.Infoln("Info ln:", 0)
	log.Warnf("Warn f: %d\n", 0)
	log.Warnln("Warn ln:", 0)
	log.Errorf("Error f: %d\n", 0)
	log.Errorln("Error ln:", 0)
	log.Logf(log.LEVEL_INFO, "Log f: %d\n", 0)
	log.Logln(log.LEVEL_INFO, "Log ln:", 0)
}

