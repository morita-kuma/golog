package golog

import (
)

/*
func TestFormatLogEvent_Encode(t *testing.T) {

	formatLogEvent := FormatLogEvent{
		format: "%d %d %d",
		args:   []interface{}{1, 2, 3},
	}

	metadata := NewDefaultLogEventMetadata("defaultLogger", LogLevel_TRACE)
	metadata.TimeFormatter = func(time Time) string {
		return "[timestamp]"
	}
	formatted :=  formatLogEvent.Encode(metadata)
	assert.Equal(t, "[TRACE] [timestamp] defaultLogger logevent_format_test.go(15) 1 2 3", string(formatted))
}
*/