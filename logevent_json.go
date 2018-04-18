package golog

import (
	"fmt"
	"os"
	"encoding/json"
)

type EventData interface {

}

type JsonLogEvent struct {
	event EventData
}

// Encode is implementation of LogEvent.Encode
func (jsonLogEvent JsonLogEvent) Encode(data *LogEventMetadata) []byte {

	if data == nil {
		// encode json
		encoded, err := json.Marshal(jsonLogEvent.event)
		if err != nil {
			fmt.Fprintf(os.Stdout, err.Error())
		}

		return encoded
	} else {

		eventData := struct {
			EventData

			LogLevel   string `json:"logLevel,omitempty"`
			Time       string `json:"timestamp,omitempty"`
			SourceLine string `json:"sourceLine,omitempty"`
			SourceFile string `json:"sourceFile,omitempty"`
			LoggerName string `json:"loggerName,omitempty"`
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
}
