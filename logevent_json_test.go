package golog

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestJsonLogEvent_Encode(t *testing.T) {

	func() {

		logEvent := JsonLogEvent{
			event: struct {
				Name    string `json:"name"`
				Address string `json:"address"`
			}{
				Name:    "name_value",
				Address: "address_value",
			},
		}

		metadata := NewDefaultLogEventMetadata("defaultLogger", LogLevel_TRACE)
		metadata.TimeFormatter = func(time Time) string {
			return "[timestamp]"
		}
		buf := logEvent.Encode(metadata)

		expected := `{"EventData":{"name":"name_value","address":"address_value"},"logLevel":"[TRACE]","timestamp":"[timestamp]","sourceLine":"22","sourceFile":"logevent_json_test.go","loggerName":"defaultLogger"}`
		assert.Equal(t, expected, string(buf))
	}()

}
