package golog

// LogEvent
type LogEvent interface {

	Encode(metadata *LogEventMetadata) []byte
}
