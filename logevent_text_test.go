package golog

import (
)

/*
func TestTextLogEvent_Encode(t *testing.T) {

	func() {
		metadata := NewDefaultLogEventMetadata("defaultLogger", LogLevel_TRACE)
		metadata.TimeFormatter = func(time Time) string {
			return "[timestamp]"
		}
		expected := `[TRACE] [timestamp] defaultLogger logevent_text_test.go(11) test`
		buf := TextLogEvent{Event: "test"}.Encode(metadata)
		assert.Equal(t, expected, string(buf))
	}()
}
*/

