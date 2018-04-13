package golog

import (
	"io"
	"strings"
)

type TextLogEvent struct {
	Event string
}

// Encode implements LogEvent.Encode
func (textLogEvent TextLogEvent) Encode(metadata *LogEventMetadata) io.Reader {
	if metadata != nil {

		/*
		return strings.NewReader(fmt.Sprintf("%s %s %s %s(%s) %s",
			metadata.GetLogLevel(),
			metadata.GetTime(),
			metadata.GetLoggerName(),
			metadata.GetSourceFile(),
			metadata.GetSourceLine(),
			textLogEvent.Event,))
		*/


		return 	strings.NewReader(
			metadata.GetLogLevel() + " " +
				metadata.GetTime() + " " +
					metadata.GetLoggerName() + " " +
						metadata.GetSourceFile() + "(" + metadata.GetSourceLine() +") " +
								textLogEvent.Event)
	}


	return strings.NewReader(textLogEvent.Event)
}
