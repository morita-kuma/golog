package golog

import (
	"runtime"
)

// SourceLine
type SourceLine = int

// SourceFile
type SourceFile = string

// LoggerName
type LoggerName = string

// Time
type Time = int64

// LogEventMetadata
type LogEventMetadata struct {
	LogLevel   LogLevel
	Time       Time
	SourceFile SourceFile
	SourceLine SourceLine
	LoggerName LoggerName

	MetadataFormatter
	MetadataConfig
}

// MetadataConfig
type MetadataConfig struct {
	IsEnabledLogLevel   bool
	IsEnabledTime       bool
	IsEnabledSourceFile bool
	IsEnabledSourceLine bool
	IsEnabledLoggerName bool
}

// NewDefaultMetadataConfig
func NewDefaultMetadataConfig() MetadataConfig {
	return MetadataConfig{
		IsEnabledLoggerName: true,
		IsEnabledLogLevel:   true,
		IsEnabledSourceFile: true,
		IsEnabledSourceLine: true,
		IsEnabledTime:       true,
	}
}

// GetLogLevel returns log level formatted by LogLevelFormatter
// if metadata config is disabled, returns empty string
func (metadata *LogEventMetadata) GetLogLevel() string {
	if metadata.IsEnabledLogLevel == false {
		return ""
	}

	return metadata.LogLevelFormatter(metadata.LogLevel)
}

// GetTime returns time formatted by TimeFormatter
// if metadata config is disabled, returns empty string
func (metadata *LogEventMetadata) GetTime() string {
	if metadata.IsEnabledTime == false {
		return ""
	}

	return metadata.TimeFormatter(metadata.Time)
}

// GetSourceFile returns package name formatted by SourceFileFormatter
// if metadata config is disabled, returns empty string
func (metadata *LogEventMetadata) GetSourceFile() string {
	if metadata.IsEnabledSourceFile == false {
		return ""
	}

	return metadata.SourceFileFormatter(metadata.SourceFile)
}

// GetSourceFile returns package name formatted by SourceFileFormatter
// if metadata config is disabled, returns empty string
func (metadata *LogEventMetadata) GetLoggerName() string {
	if metadata.IsEnabledLoggerName == false {
		return ""
	}

	return metadata.LoggerNameFormatter(metadata.LoggerName)
}

// GetSourceLine returns line formatted by SourceLineFormatter
// if metadata config is disabled, returns empty string
func (metadata *LogEventMetadata) GetSourceLine() string {
	if metadata.IsEnabledSourceLine == false {
		return ""
	}

	return metadata.SourceLineFormatter(metadata.SourceLine)
}

// setLoggerName
func (metadata *LogEventMetadata) setLoggerName(loggerName string) {
	if metadata == nil {
		return
	}

	if metadata.IsEnabledLoggerName == true {
		metadata.LoggerName = loggerName
	}
}

// setLogLevel
func (metadata *LogEventMetadata) setLogLevel(logLevel LogLevel) {
	if metadata == nil {
		return
	}

	if metadata.IsEnabledLogLevel == true {
		metadata.LogLevel = logLevel
	}
}

// setSource
func (metadata *LogEventMetadata) setSource(skip int) {
	if metadata == nil {
		return
	}

	if metadata.IsEnabledSourceLine == true || metadata.IsEnabledSourceFile == true {

		_, file, line, _ := runtime.Caller(skip)

		metadata.SourceLine = line

		metadata.SourceFile = file

	}
}

// setTime
func (metadata *LogEventMetadata) setTime() {
	if metadata == nil {
		return
	}

	if metadata.IsEnabledTime == true {
	}
}

// NewLogEventMetadata
func NewLogEventMetadata(config *MetadataConfig, formatter *MetadataFormatter) LogEventMetadata {

	metadata := LogEventMetadata{
		MetadataFormatter: NewDefaultMetadataFormatter(),
		MetadataConfig:NewDefaultMetadataConfig(),
	}

	if config != nil {
		metadata.MetadataConfig = *config
	}

	if formatter != nil {
		metadata.MetadataFormatter = *formatter
	}

	return metadata
}

// newLogEventMetadata
func newDefaultLogEventMetadata(loggerName string, logLevel LogLevel) *LogEventMetadata {
	metadata := LogEventMetadata{
		MetadataFormatter: NewDefaultMetadataFormatter(),
		MetadataConfig:NewDefaultMetadataConfig(),
	}
	metadata.setLogLevel(logLevel)
	metadata.setLoggerName(loggerName)
	metadata.setSource(2)
	metadata.setTime()

	return &metadata
}