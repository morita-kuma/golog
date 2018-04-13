package golog

import (
	"bytes"
	"sync"
)

// ByteBufferAppender
type ByteBufferAppender struct {
	bufferPool *sync.Pool
}

// Write implements Appender interface
func (appender *ByteBufferAppender) Write(data []byte) (n int, err error) {

	if appender != nil {

		// get buffer from bufferPool
		buffer := appender.bufferPool.Get().(*bytes.Buffer)

		// release
		defer func() {
			appender.bufferPool.Put(buffer)
		}()

		// write
		buffer.Write(data)
		buffer.WriteString("\n")
	}

	return 0, nil
}

// String is implementation of Appender interface
func (appender *ByteBufferAppender) String() string {

	if appender != nil {
		buffer := appender.bufferPool.Get().(*bytes.Buffer)
		return buffer.String()
	}

	return ""
}

// NewByteBufferAppender
func NewByteBufferAppender() *ByteBufferAppender {

	bufferPool := &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	return &ByteBufferAppender{
		bufferPool: bufferPool,
	}
}
