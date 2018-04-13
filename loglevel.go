package golog

import "sort"

// LogLevel
type LogLevel int32

// LogLevel Constants
const (
	LogLevel_TRACE LogLevel = 0
	LogLevel_DEBUG LogLevel = 1
	LogLevel_INFO  LogLevel = 2
	LogLevel_WARN  LogLevel = 3
	LogLevel_ERROR LogLevel = 4
	LogLevel_FATAL LogLevel = 5
)

// TypeVal
func (logLevel LogLevel) TypeVal() int32 {
	return int32(logLevel)
}

// String
func (logLevel LogLevel) String() string {
	switch logLevel {
	case LogLevel_TRACE :
		return "[TRACE]"
	case LogLevel_DEBUG:
		return "[DEBUG]"
	case LogLevel_INFO:
		return "[INFO]"
	case LogLevel_WARN:
		return "[WARN]"
	case LogLevel_ERROR:
		return "[ERROR]"
	case LogLevel_FATAL:
		return "[FATAL]"
	default:
		return "[UNKNOWN]"
	}
}

// TODO go generate
var logLevelMap = map[int32]LogLevel {
	0 : LogLevel_TRACE,
	1 : LogLevel_DEBUG,
	2 : LogLevel_INFO,
	3 : LogLevel_WARN,
	4 : LogLevel_ERROR,
	5 : LogLevel_FATAL,
}

// LogLevels
type LogLevels []LogLevel

// Contains
func (levels LogLevels) Contains (logLevel LogLevel) bool {
	for _, level := range levels {
		if level == logLevel {
			return true
		}
	}
	return false
}

// Len implements the Sort interface
func (levels LogLevels) Len() int {
	return len(levels)
}

// Swap implements the Sort interface
func (levels LogLevels) Swap(i, j int) {
	levels[i], levels[j] = levels[j], levels[i]
}

// Less implements the Sort interface
func (levels LogLevels) Less(i, j int) bool {
	return  levels[i] < levels[j]
}

// Sort by ID ASC
func (levels LogLevels) SortAsc() LogLevels {
	sort.Sort(levels)
	return levels
}