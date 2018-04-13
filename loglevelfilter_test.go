package golog

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewLogLevelFilter(t *testing.T) {

	cases := []struct {
		input LogLevel
		expected LogLevels
	} {
		// trace
		{
			input: LogLevel_TRACE,
			expected: LogLevels {
				LogLevel_TRACE,
				LogLevel_DEBUG,
				LogLevel_INFO,
				LogLevel_WARN,
				LogLevel_ERROR,
				LogLevel_FATAL,
			},
		},

		// debug
		{
			input: LogLevel_DEBUG,
			expected: LogLevels {
				LogLevel_DEBUG,
				LogLevel_INFO,
				LogLevel_WARN,
				LogLevel_ERROR,
				LogLevel_FATAL,
			},
		},

		// info
		{
			input: LogLevel_INFO,
			expected: LogLevels {
				LogLevel_INFO,
				LogLevel_WARN,
				LogLevel_ERROR,
				LogLevel_FATAL,
			},
		},

		// warn
		{
			input: LogLevel_WARN,
			expected: LogLevels {
				LogLevel_WARN,
				LogLevel_ERROR,
				LogLevel_FATAL,
			},
		},

		// error
		{
			input: LogLevel_ERROR,
			expected: LogLevels {
				LogLevel_ERROR,
				LogLevel_FATAL,
			},
		},

		// fatal
		{
			input: LogLevel_FATAL,
			expected: LogLevels {
				LogLevel_FATAL,
			},
		},
	}

	for _, c := range cases {
		logLevelFilter := NewDefaultLevelFilter()
		actual := logLevelFilter.DoFilter(c.input)
		assert.Equal(t, actual.SortAsc(), c.expected.SortAsc())
	}
}