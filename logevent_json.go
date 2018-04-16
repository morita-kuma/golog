package golog

import (
	"fmt"
	"os"
	"encoding/json"
)

type EventData interface{}

type JsonLogEvent struct {
	event EventData
}

// Encode is implementation of LogEvent.Encode
func (jsonLogEvent JsonLogEvent) Encode(data *LogEventMetadata) []byte {

	if data == nil {

		// new event data object
		eventData := struct {
			EventData
		} {
			EventData: jsonLogEvent.event,
		}

		// encode json
		encoded, err := json.Marshal(eventData)
		if err != nil {
			fmt.Fprintf(os.Stdout, err.Error())
		}

		return encoded
	}


	eventData := struct {
		EventData

		LogLevel   string `json:"logLevel"`
		Time       string `json:"timestamp"`
		SourceLine string `json:"sourceLine"`
		SourceFile string `json:"sourceFile"`
		LoggerName string `json:"loggerName"`
	}{
		EventData: jsonLogEvent.event,

		LogLevel:   data.GetLogLevel(),
		Time:       data.GetTime(),
		SourceLine: data.GetSourceLine(),
		SourceFile: data.GetSourceFile(),
		LoggerName: data.GetLoggerName(),
	}

	encoded, err := json.Marshal(eventData)

	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
	}

	return encoded
}
