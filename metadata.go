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
	EnabledLogLevel bool
	EnabledTime bool
	EnabledSourceFile bool
	EnabledSourceLine bool
	EnabledLoggerName bool
}

func NewMetadataConfig() MetadataConfig {
	return MetadataConfig{
		EnabledLoggerName:true,
		EnabledLogLevel:true,
		EnabledSourceFile:true,
		EnabledSourceLine:true,
		EnabledTime:true,
	}
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

func (metadata *LogEventMetadata) setLoggerName(loggerName string) {
	if metadata != nil {

		if metadata.EnabledLoggerName {
			metadata.LoggerName = loggerName
		}
	}
}

func (metadata *LogEventMetadata) setLogLevel(logLevel LogLevel) {
	if metadata != nil {

		if metadata.EnabledLogLevel {
			metadata.LogLevel = logLevel
		}
	}
}

func (metadata *LogEventMetadata) SetSource(skip int) {
	if metadata != nil {
		_, file, line, _ := runtime.Caller(skip)
		_, fileName := filepath.Split(file)

		if metadata.EnabledSourceFile {
			metadata.SourceFile = fileName
		}

		if metadata.EnabledSourceLine {
			metadata.SourceLine = line
		}
	}
}

// NewLogEventMetadata
func NewLogEventMetadata(config *MetadataConfig, formatter *MetadataFormatter) *LogEventMetadata {

	metadata := LogEventMetadata{}
	metadata.MetadataConfig = NewMetadataConfig()
	if config != nil {
		metadata.MetadataConfig = *config
	}

	metadata.MetadataFormatter = NewDefaultMetadataFormatter()
	if formatter != nil {
		metadata.MetadataFormatter = *formatter
	}

	if config.EnabledTime {
		metadata.Time = time.Now().UnixNano()
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

	defaultLogEventMetadataFormatter := MetadataFormatter{
		LogLevelFormatter:   defaultLogLevelFormatter,
		TimeFormatter:       defaultTimeFormatter,
		SourceLineFormatter: defaultSourceLineFormatter,
		SourceFileFormatter: defaultSourceFileFormatter,
		LoggerNameFormatter: defaultLoggerNameFormatter,
	}
	return defaultLogEventMetadataFormatter
}


