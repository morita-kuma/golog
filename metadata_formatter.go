package golog

import (
	"strconv"
	"path/filepath"

	"time"
)

// TimeFormatter
type TimeFormatter = func(UnixTime) string

// LogLevelFormatter
type LogLevelFormatter = func(level LogLevel) string

// SourceLineFormatter
type SourceLineFormatter = func(line SourceLine) string

// SourceFileFormatter
type SourceFileFormatter = func(packageName SourceFile) string

// LoggerNameFormatter
type LoggerNameFormatter = func(loggerName LoggerName) string

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
	// DefaultLogLevelFormatter
	var defaultLogLevelFormatter = func(logLevel LogLevel) string {
		return logLevel.String()
	}

	// DefaultTimeFormatter
	var defaultTimeFormatter = func(unixTime UnixTime) string {
		return time.Unix(unixTime, 0).Format(time.RFC3339)
	}

	// DefaultLineFormatter
	var defaultSourceLineFormatter = func(sourceLine SourceLine) string {
		return strconv.FormatInt(int64(sourceLine), 10)
	}

	// DefaultSourceFileFormatter
	var defaultSourceFileFormatter = func(sourceFile SourceFile) string {
		_, fileName := filepath.Split(sourceFile)
		return fileName
	}

	// DefaultLoggerNameFormatter
	var defaultLoggerNameFormatter = func(loggerName LoggerName) string {
		return loggerName
	}

	return  MetadataFormatter{
		LogLevelFormatter:   defaultLogLevelFormatter,
		TimeFormatter:       defaultTimeFormatter,
		SourceLineFormatter: defaultSourceLineFormatter,
		SourceFileFormatter: defaultSourceFileFormatter,
		LoggerNameFormatter: defaultLoggerNameFormatter,
	}
}

