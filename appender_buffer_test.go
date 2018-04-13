package golog

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io"
	"bytes"
)

func TestByteBufferAppender(t *testing.T) {

	// append Event three times
	func(){
		appender := NewByteBufferAppender()
		io.Copy(appender, bytes.NewBufferString("test1"))
		io.Copy(appender, bytes.NewBufferString("test2"))
		io.Copy(appender, bytes.NewBufferString("test3"))
		assert.Equal(t, "test1\ntest2\ntest3\n", appender.String())
	}()
}
