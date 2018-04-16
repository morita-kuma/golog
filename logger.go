package golog

import (
	"os"
	"reflect"
	"fmt"
	"io"
)

// Logger
type Logger struct {
	// LoggerName
	// Public Required
	//
	// It can be used as an identifier when outputting logs.
	// If you specify an empty string, the default logger name is substituted
	Name string

	// levelAppender
	// Private Required
	levelAppender map[LogLevel][]Appender

	// enabledMetadata
	// Private Required
	enabledMetadata bool

	// metadataFormatter
	// Private Option
	//
	// If not specified, the default formatter will be used
	metadataFormatter *MetadataFormatter
}

// doAppendIfLevelEnabled
func (logger *Logger) doAppendIfLevelEnabled(event []byte, level LogLevel) {

	// recover
	defer func(writer io.Writer) {
		if err := recover(); err != nil {
			fmt.Fprintln(writer, "Error: golog exit appending error:", err)
		}
	}(os.Stderr)

	if appenders, ok := logger.levelAppender[level]; ok {
		for _, appender := range appenders {
			appender.Write(event)
		}
	}
}

func (logger *Logger) Trace(string string) {

	var metadata *LogEventMetadata
	if logger.enabledMetadata {
		metadata = NewLogEventMetadata(logger.Name, LogLevel_TRACE, logger.metadataFormatter, 4)
	}
	logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(metadata), LogLevel_TRACE)
}

func (logger *Logger) Debug(string string) {

	var metadata *LogEventMetadata
	if logger.enabledMetadata {
		metadata = NewLogEventMetadata(logger.Name, LogLevel_TRACE, logger.metadataFormatter, 4)
	}
	logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(metadata), LogLevel_DEBUG)
}

func (logger *Logger) Info(string string) {

	var metadata *LogEventMetadata
	if logger.enabledMetadata {
		metadata = NewLogEventMetadata(logger.Name, LogLevel_TRACE, logger.metadataFormatter, 4)
	}
	logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(metadata), LogLevel_INFO)
}

func (logger *Logger) Warn(string string) {

	var metadata *LogEventMetadata
	if logger.enabledMetadata {
		metadata = NewLogEventMetadata(logger.Name, LogLevel_WARN, logger.metadataFormatter, 4)
	}
	logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(metadata), LogLevel_WARN)
}

func (logger *Logger) Error(string string) {

	var metadata *LogEventMetadata
	if logger.enabledMetadata {
		metadata = NewLogEventMetadata(logger.Name, LogLevel_ERROR, logger.metadataFormatter, 4)
	}
	logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(metadata), LogLevel_ERROR)
}

func (logger *Logger) Fatal(string string) {

	var metadata *LogEventMetadata
	if logger.enabledMetadata {
		metadata = NewLogEventMetadata(logger.Name, LogLevel_FATAL, logger.metadataFormatter, 4)
	}
	logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(metadata), LogLevel_FATAL)
	os.Exit(1)
}

// TraceF
func (logger *Logger) TraceF(format string, args ...interface{}) {
	logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_TRACE)
}

// DebugF
func (logger *Logger) DebugF(format string, args ...interface{}) {
	logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_DEBUG)
}

// InfoF
func (logger *Logger) InfoF(format string, args ...interface{}) {
	logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_INFO)
}

// WarnF
func (logger *Logger) WarnF(format string, args ...interface{}) {
	logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_WARN)
}

// ErrorF
func (logger *Logger) ErrorF(format string, args ...interface{}) {
	logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_ERROR)
}

// FatalF
func (logger *Logger) FatalF(format string, args ...interface{}) {
	logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_FATAL)
}

// TraceJ
func (logger *Logger) TraceJ(obj interface{}) {
	logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_TRACE)
}

// DebugJ
func (logger *Logger) DebugJ(obj interface{}) {
	logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_DEBUG)
}

// InfoJ
func (logger *Logger) InfoJ(obj interface{}) {
	logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_INFO)
}

// WarnJ
func (logger *Logger) WarnJ(obj interface{}) {
	logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_WARN)
}

// ErrorJ
func (logger *Logger) ErrorJ(obj interface{}) {
	logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_ERROR)
}

// FatalJ
func (logger *Logger) FatalJ(obj interface{}) {
	logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_FATAL)
}

func (logger *Logger) TraceS(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_TRACE)
}

func (logger *Logger) DebugS(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_DEBUG)
}

func (logger *Logger) InfoS(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_INFO)
}

func (logger *Logger) WarnS(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_WARN)
}

func (logger *Logger) ErrorS(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_ERROR)
}

func (logger *Logger) FatalS(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_FATAL)
}

// SetAppender
func (logger *Logger) SetAppender(appender ...Appender) {
	for k := range logger.levelAppender {
		logger.levelAppender[k] = appender
	}
}

// DisableLogEventMetadata
// If metadata is unnecessary, please disable it.
// It is possible to prevent unnecessary allocation.
// It is enabled by default.
func (logger *Logger) DisableLogEventMetadata() {
	logger.enabledMetadata = false
}

// SetMetadataFormatter
func (logger *Logger) SetMetadataFormatter(formatter *MetadataFormatter) {
	logger.metadataFormatter = formatter
}

// SetLogLevel enables the specified log level
func (logger *Logger) SetAppenderWithLevel(logLevel LogLevel, appender ...Appender) {
	logger.levelAppender[logLevel] = appender
}

// SetLogLevel enables the specified log level
func (logger *Logger) SetAppenderWithLevels(logLevels []LogLevel, appender ...Appender) {
	for _, v := range logLevels {
		logger.levelAppender[v] = appender
	}
}

// NewLogger
func NewLogger(loggerName string, logLevel LogLevel, appender ...Appender) Logger {
	levelAppender := map[LogLevel][]Appender{}
	logLevels := NewDefaultLevelFilter().DoFilter(logLevel)
	for _, logLevel := range logLevels {
		levelAppender[logLevel] = appender
	}

	// Print Filter and Appender settings
	/*
	fmt.Println("golog logger is initialized.")
	for logLevel, appenders := range levelAppender {

		var appenderNames []string
		for _, appender := range appenders {
			appenderNames = append(appenderNames, reflect.TypeOf(appender).Name())
		}

		info := fmt.Sprintf("LoggerName: %s LEVEL: %s appenders: %s", loggerName, logLevel.String(), strings.Join(appenderNames,","))
		fmt.Println(info)
	}
	*/

	return Logger{
		Name:            loggerName,
		levelAppender:   levelAppender,
		enabledMetadata: true,
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

// NewDefaultLogger
func NewDefaultLogger() Logger {
	return NewLogger("defaultLogger", LogLevel_TRACE, NewDefaultConsoleAppender())
}
