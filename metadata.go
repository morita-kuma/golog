package golog

import (
	"time"
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

// TimeFormatter
type TimeFormatter = func(Time) string

// LogLevelFormatter
type LogLevelFormatter = func(level LogLevel) string

// SourceLineFormatter
type SourceLineFormatter = func(line SourceLine) string

// SourceFileFormatter
type SourceFileFormatter = func(packageName SourceFile) string

// LoggerNameFormatter
type LoggerNameFormatter = func(loggerName LoggerName) string

// LogEventMetadata
type LogEventMetadata struct {
	LogLevel   LogLevel
	Time       Time
	SourceFile SourceFile
	SourceLine SourceLine
	LoggerName LoggerName

	MetadataFormatter
}

// MetadataFormatter
type MetadataFormatter struct {
	LogLevelFormatter   LogLevelFormatter
	TimeFormatter       TimeFormatter
	SourceFileFormatter SourceFileFormatter
	SourceLineFormatter SourceLineFormatter
	LoggerNameFormatter LoggerNameFormatter
}

// NewDefaultMetadataFormatter
func NewDefaultMetadataFormatter() MetadataFormatter {
	defaultLogEventMetadataFormatter := MetadataFormatter{
		LogLevelFormatter:   defaultLogLevelFormatter,
		TimeFormatter:       defaultTimeFormatter,
		SourceLineFormatter: defaultSourceLineFormatter,
		SourceFileFormatter: defaultSourceFileFormatter,
		LoggerNameFormatter: defaultLoggerNameFormatter,
	}
	return defaultLogEventMetadataFormatter
}


// GetLogLevel returns log level formatted by LogLevelFormatter
func (metadata *LogEventMetadata) GetLogLevel() string {
	return metadata.LogLevelFormatter(metadata.LogLevel)
	}

// GetTime returns time formatted by TimeFormatter
func (metadata *LogEventMetadata) GetTime() string {
	return metadata.TimeFormatter(metadata.Time)
}

// GetSourceFile returns package name formatted by SourceFileFormatter
func (metadata *LogEventMetadata) GetSourceFile() string {
	return metadata.SourceFileFormatter(metadata.SourceFile)
}

// GetSourceFile returns package name formatted by SourceFileFormatter
func (metadata *LogEventMetadata) GetLoggerName() string {
	return metadata.LoggerNameFormatter(metadata.LoggerName)
}

// GetSourceLine returns line formatted by SourceLineFormatter
func (metadata *LogEventMetadata) GetSourceLine() string {
	return metadata.SourceLineFormatter(metadata.SourceLine)
}

// DefaultLogLevelFormatter
var defaultLogLevelFormatter = func(logLevel LogLevel) string {
	return logLevel.String()
}

// DefaultTimeFormatter
var defaultTimeFormatter = func(time Time) string {
	return string(time)
}

// DefaultLineFormatter
var defaultSourceLineFormatter = func(sourceLine SourceLine) string {
	return strconv.FormatInt(int64(sourceLine), 10)
}

// DefaultSourceFileFormatter
var defaultSourceFileFormatter = func(sourceFile SourceFile) string {
	return sourceFile
}

// DefaultLoggerNameFormatter
var defaultLoggerNameFormatter = func(loggerName LoggerName) string {
	return loggerName
}

// NewLogEventMetadata
func NewLogEventMetadata(loggerName string, logLevel LogLevel, formatter *MetadataFormatter, skip int) *LogEventMetadata {

	 _, file, line, _ := runtime.Caller(skip)
	_, fileName := filepath.Split(file)

	metadataFormatter := NewDefaultMetadataFormatter()
	if formatter != nil {
		metadataFormatter = *formatter
	}

	metadata := LogEventMetadata{
		Time:              time.Now().UnixNano(),
		SourceLine:        line,
		SourceFile:        fileName,
		LogLevel:          logLevel,
		LoggerName:        loggerName,
		MetadataFormatter: metadataFormatter,
	}
	return &metadata
}

// NewLogEventMetadata
func NewDefaultLogEventMetadata(loggerName string, logLevel LogLevel) *LogEventMetadata {
	_, file, line, _ := runtime.Caller(1)
	_, fileName := filepath.Split(file)

	metadata := LogEventMetadata{
		Time:              time.Now().UnixNano(),
		SourceLine:        line,
		SourceFile:        fileName,
		LogLevel:          logLevel,
		LoggerName:        loggerName,
		MetadataFormatter: NewDefaultMetadataFormatter(),
	}
	return &metadata
}


