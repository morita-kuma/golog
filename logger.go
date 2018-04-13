package golog

import (
	"os"
	"fmt"
	"io"
	"reflect"
)

// Logger
type Logger struct {

	// LoggerName
	// Public Required
	//
	// It can be used as an identifier when outputting logs.
	// If you specify an empty string, the default logger name is substituted
	Name              string

	// levelAppender
	// Private Required
	levelAppender     map[LogLevel][]Appender

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
func (logger *Logger) doAppendIfLevelEnabled(event LogEvent, level LogLevel) {

	// recover
	defer func(writer io.Writer) {
		if err := recover(); err != nil {
			fmt.Fprintln(writer, "Error: golog exit appending error:", err)
		}
	}(os.Stderr)

	var metadata *LogEventMetadata
	if logger.enabledMetadata {
		metadata = NewLogEventMetadata(logger.Name, level, logger.metadataFormatter,4)
	}

	if appenders, ok := logger.levelAppender[level]; ok {
		for _, appender := range appenders {
			io.Copy(appender, event.Encode(metadata))
		}
	}
}

func (logger *Logger) trace(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent, LogLevel_TRACE)
}

func (logger *Logger) debug(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent, LogLevel_DEBUG)
}

func (logger *Logger) info(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent, LogLevel_INFO)
}

func (logger *Logger) warn(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent, LogLevel_WARN)
}

func (logger *Logger) error(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent, LogLevel_ERROR)
}

func (logger *Logger) fatal(logEvent LogEvent) {
	logger.doAppendIfLevelEnabled(logEvent, LogLevel_FATAL)
	os.Exit(1)
}

func (logger *Logger) Trace(string string) {
	logger.trace(TextLogEvent{Event: string})
}

func (logger *Logger) Debug(string string) {
	logger.debug(TextLogEvent{Event: string})
}

func (logger *Logger) Info(string string) {
	logger.info(TextLogEvent{Event: string})
}

func (logger *Logger) Warn(string string) {
	logger.warn(TextLogEvent{Event: string})
}

func (logger *Logger) Error(string string) {
	logger.error(TextLogEvent{Event: string})
}

func (logger *Logger) Fatal(string string) {
	logger.fatal(TextLogEvent{Event: string})
}

func (logger *Logger) TraceS(logEvent LogEvent) {
	logger.trace(logEvent)
}

func (logger *Logger) DebugS(logEvent LogEvent) {
	logger.debug(logEvent)
}

func (logger *Logger) InfoS(logEvent LogEvent) {
	logger.info(logEvent)
}

func (logger *Logger) WarnS(logEvent LogEvent) {
	logger.warn(logEvent)
}

func (logger *Logger) ErrorS(logEvent LogEvent) {
	logger.error(logEvent)
}

func (logger *Logger) FatalS(logEvent LogEvent) {
	logger.fatal(logEvent)
}

// TraceF
func (logger *Logger) TraceF(format string, args ...interface{}) {
	logger.trace(FormatLogEvent{format: format, args: args,})
}

// DebugF
func (logger *Logger) DebugF(format string, args ...interface{}) {
	logger.debug(FormatLogEvent{format: format, args: args,})
}

// InfoF
func (logger *Logger) InfoF(format string, args ...interface{}) {
	logger.info(FormatLogEvent{format: format, args: args,})
}

// WarnF
func (logger *Logger) WarnF(format string, args ...interface{}) {
	logger.warn(FormatLogEvent{format: format, args: args,})
}

// ErrorF
func (logger *Logger) ErrorF(format string, args ...interface{}) {
	logger.error(FormatLogEvent{format: format, args: args,})
}

// FatalF
func (logger *Logger) FatalF(format string, args ...interface{}) {
	logger.fatal(FormatLogEvent{format: format, args: args,})
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
	for _,v := range logLevels {
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
		Name : loggerName,
		levelAppender:levelAppender,
		enabledMetadata:true,
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

