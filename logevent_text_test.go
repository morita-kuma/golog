package golog

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestTextLogEvent_Encode(t *testing.T) {

	func () {
		metadata := NewDefaultLogEventMetadata("defaultLogger", LogLevel_TRACE)
		metadata.TimeFormatter = func(time Time) string {
			return "[timestamp]"
		}
		expected := `[TRACE] [timestamp] defaultLogger logevent_text_test.go(12) test`
		reader := TextLogEvent{Event:"test"}.Encode(metadata)
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(reader)
		assert.Equal(t, expected, buffer.String())
	}()


}