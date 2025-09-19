package main

import "time"

func main() {
	Infof("Logging without prior initialization!\n")
	Warnln("PURE ANARCHY")
	config := LogConfig {
		SerializationStrategy: SIMPLE,
		Timeformat: time.RFC3339,
		TargetMode: STDOUT,
		Logfile: "",
		Level: INFO,
	}
	Init(config)

	Errorf("Logging an error: %s\n", "ayayayaa")
	Errorln("This is just a single line error")
	Errorf("Another fromatted log arg1: %s, arg2: %d\n", "log1", 4)

	file := "./log.txt"
	Logf(INFO, "Updating configuration...\n")
	Logf(INFO, "Now writing to file: %s\n", file)
	config.Logfile = file
	config.TargetMode = FILE
	Init(config)
	Errorln("First line written to file?")
	Errorf("Formatted %s and wrote to file: %s\n", "xd", file)

	Errorln("Now printing json")
	config.SerializationStrategy = JSON
	Init(config)
	Errorf("First time formatting %s as json message: %s\n", "first", "second")
	Errorln("And a simple line")
	Logf(WARN, "Warning: %s\n", "adsada")

	Errorln("And lastly, json to STDOUT")
	config.TargetMode = STDOUT
	Init(config)
	Errorln("This should be a json written to STDOUT")

	config.SerializationStrategy = SIMPLE
	Init(config)
	Infof("Testing the log level. Currently set to: %s\n", config.Level.Name())
	Warnln("Warning should work as well.")
	Errorln("Error always works with predefined level")
	Debugln("Debug however is lower and should not be printed")
}
