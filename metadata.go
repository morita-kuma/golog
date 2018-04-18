package golog

import (
	"strconv"
	"runtime"
	"path/filepath"
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

type MetadataConfig struct {
	EnabledLogLevel   bool
	EnabledTime       bool
	EnabledSourceFile bool
	EnabledSourceLine bool
	EnabledLoggerName bool
}

func NewDefaultMetadataConfig() MetadataConfig {
	return MetadataConfig{
		EnabledLoggerName: true,
		EnabledLogLevel:   true,
		EnabledSourceFile: true,
		EnabledSourceLine: true,
		EnabledTime:       true,
	}
}

// GetLogLevel returns log level formatted by LogLevelFormatter
func (metadata *LogEventMetadata) GetLogLevel() string {
	if metadata.EnabledLogLevel == false {
		return ""
	}

	return metadata.LogLevelFormatter(metadata.LogLevel)
}

// GetTime returns time formatted by TimeFormatter
func (metadata *LogEventMetadata) GetTime() string {
	if metadata.EnabledTime == false {
		return ""
	}

	return metadata.TimeFormatter(metadata.Time)
}

// GetSourceFile returns package name formatted by SourceFileFormatter
func (metadata *LogEventMetadata) GetSourceFile() string {
	if metadata.EnabledSourceFile == false {
		return ""
	}

	return metadata.SourceFileFormatter(metadata.SourceFile)
}

// GetSourceFile returns package name formatted by SourceFileFormatter
func (metadata *LogEventMetadata) GetLoggerName() string {
	if metadata.EnabledLoggerName == false {
		return ""
	}

	return metadata.LoggerNameFormatter(metadata.LoggerName)
}

// GetSourceLine returns line formatted by SourceLineFormatter
func (metadata *LogEventMetadata) GetSourceLine() string {
	if metadata.EnabledSourceLine == false {
		return ""
	}

	return metadata.SourceLineFormatter(metadata.SourceLine)
}

// setLoggerName
func (metadata *LogEventMetadata) setLoggerName(loggerName string) {
	if metadata == nil {
		return
	}

	if metadata.EnabledLoggerName == true {
		metadata.LoggerName = loggerName
	}
}

// setLogLevel
func (metadata *LogEventMetadata) setLogLevel(logLevel LogLevel) {
	if metadata == nil {
		return
	}

	if metadata.EnabledLogLevel == true {
		metadata.LogLevel = logLevel
	}
}

// setSource
func (metadata *LogEventMetadata) setSource(skip int) {
	if metadata == nil {
		return
	}

	if metadata.EnabledSourceLine == true || metadata.EnabledSourceFile == true {

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

	if metadata.EnabledTime == true {
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

// NewLogEventMetadata
func NewDefaultLogEventMetadata(loggerName string, logLevel LogLevel) *LogEventMetadata {
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