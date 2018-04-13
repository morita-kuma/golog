package golog

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestJsonLogEvent_Encode(t *testing.T) {

	func () {

		logEvent := JsonLogEvent{
			event: struct{
				Name string `json:"name"`
				Address string `json:"address"`
			}{
				Name:"name_value",
				Address:"address_value",
			},
		}

		metadata := NewDefaultLogEventMetadata("defaultLogger", LogLevel_TRACE)
		metadata.TimeFormatter = func(time Time) string {
			return "[timestamp]"
		}
		eventReader := logEvent.Encode(metadata)
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(eventReader)

		expected := `{"EventData":{"name":"name_value","address":"address_value"},"logLevel":"[TRACE]","timestamp":"[timestamp]","sourceLine":"23","sourceFile":"logevent_json_test.go","loggerName":"defaultLogger"}`
		assert.Equal(t, expected, buffer.String())
	}()

}
