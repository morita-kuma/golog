package golog

import (
	"io"
)

// LogEvent
type LogEvent interface {

	/*
	// Read
	io.Reader

	// WriterTo
	io.WriterTo

	AppendMetadata(metadata *LogEventMetadata)
	*/

	Encode(metadata *LogEventMetadata) io.Reader
}
