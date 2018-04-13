package golog

import (
	"io"
	"fmt"
)

// FormatLogEvent
type FormatLogEvent struct {
	format string
	args   []interface{}
}

// Encode implements LogEvent.Encode
func (event FormatLogEvent) Encode(metadata *LogEventMetadata) io.Reader {
	// delegate to Text log Event
	return TextLogEvent{Event: fmt.Sprintf(event.format, event.args...),}.Encode(metadata)
}
