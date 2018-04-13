package golog

import (
	"testing"
	"io"
	"bytes"
)

func TestConsoleAppender(t *testing.T) {

	// Append
	func() {
		appender := NewDefaultConsoleAppender()
		io.Copy(appender, bytes.NewBufferString("test1"))
		io.Copy(appender, bytes.NewBufferString("test2"))
		io.Copy(appender, bytes.NewBufferString("test3"))
		/*
		expected := "test1\ntest2\ntest3\n"
		assert.Equal(t, expected, appender)
		*/
	}()

}
