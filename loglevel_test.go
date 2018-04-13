package golog

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLogLevels_Contains(t *testing.T) {

	// is contains = true
	func () {
		logLevels := LogLevels{LogLevel_INFO, LogLevel_ERROR, LogLevel_FATAL,}
		assert.Equal(t, true, logLevels.Contains(LogLevel_ERROR))

	}()

	// is contains = false
	func () {
		logLevels := LogLevels{LogLevel_FATAL}
		assert.Equal(t, false, logLevels.Contains(LogLevel_DEBUG))
	}()

}

func TestLogLevels_Len(t *testing.T) {

	// Length is 1
	func () {
		logLevels := LogLevels{LogLevel_FATAL}
		assert.Equal(t, 1, logLevels.Len())
	}()
}


func TestLogLevels_Sort(t *testing.T) {

	// is Sorted LogLevels
	func () {
		logLevels := LogLevels {
			LogLevel_FATAL, // 5
			LogLevel_TRACE, // 0
			LogLevel_DEBUG, // 1
		}

		expected := LogLevels {
			LogLevel_TRACE, // 0
			LogLevel_DEBUG, // 1
			LogLevel_FATAL, // 5
		}

		assert.Equal(t, expected, logLevels.SortAsc())
	}()
}