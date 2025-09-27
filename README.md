# Logging
This module aims at delivering an improved experience from the std::log module.

## Usage
Import the go modue:
```golang
import github.com/buecherregale/log
``` 
Create a config and initialize logging module:
```golang
func main() {
    config := log.LogConfig {
        // your configuration here
    }
    log.Init(&config)
}
```
Start Logging!

## Configuration
A note on the additional configuration: 
`Logfile` is the path with filename of the file used for logging. If this file does not exist, its created. Else it is **appended**.

`Timeformat` is a string to be used with stds `time.Format()` method. For ease of use utilize the `time.RFC...` standards.

### Environment
Controlling the logging level and target file via environment variables is supported. 
They will automatically be read during the `log.Init()` call. 
The variables are: 
- `GO_LOG_LEVEL`: Expecting the `Name()` of the level, e.g. `DEBUG` (Case-Insensitive)
- `GO_LOG_FILE`: Expecting the path to the file
Setting level or file in the code configuration will overwrite the environment values.

### Default
Without calling `Init()` a standard configuration is used:
```golang
var config *LogConfig = &LogConfig{
	Level: INFO,
	Timeformat: time.RFC3339,
	SerializationStrategy: SIMPLE,
	TargetMode: STDOUT,
	Logfile: "",
} 
```

## Improvements
Improvements include:
1. Logging Levels: Users are able to configure a level impacting which logs are actually printed, e.g. `DEBUG` for development or `INFO` for production.
1. Serialization Strategy: This module implements different serialization strategies, for example classic strings or a `json` format.
1. Target Mode: The user can configure the target of the print to files or stdout.
1. Central Configuration: The central `Init()` function allows simple configuration.

### Logging Levels
The default logging levels are:
1. `ERROR`
1. `WARN`
1. `INFO`
1. `DEBUG`
Setting the `Level` in the configuration means only Logs at or **above** the level will be printed. 

### Serialization
Following Serialization Strategies are supported:
1. `JSON`: Serialize a [LogEntry](src/log.go) as prettyfied json
1. `SIMPLE`: Normal serialization to a formatted string, containing the same info as the json but on a single line. 

### Printing
Printing is supported with 2 target modes:
1. `STDOUT`: Prints with `fmt.Printf()` to standard out.
1. `FILE`: Creates or appends the file located at `LogConfig.Filename` instead.

