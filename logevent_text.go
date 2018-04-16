package golog

import (
	"bytes"
	"sync"
)

var bufferPool *sync.Pool

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// TextLogEvent
type TextLogEvent struct {
	Event string
}

// Encode implements LogEvent.Encode
func (logEvent TextLogEvent) Encode(metadata *LogEventMetadata) []byte {
	if metadata != nil {

		data := metadata.GetLogLevel() + " " +
			metadata.GetTime() + " " +
			metadata.GetLoggerName() + " " +
			/*
			metadata.GetSourceFile() + "(" +
				metadata.GetSourceLine() + ") " +
			*/
					logEvent.Event


		// get buffer from bufferPool
		buffer := bufferPool.Get().(*bytes.Buffer)
		buffer.WriteString(data)

		// release
		defer func() {
			buffer.Reset()
			bufferPool.Put(buffer)
		}()

		return buffer.Bytes()
	}

	// get buffer from bufferPool
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.WriteString(logEvent.Event)

	// release
	defer func() {
		buffer.Reset()
		bufferPool.Put(buffer)
	}()

	return buffer.Bytes()
}