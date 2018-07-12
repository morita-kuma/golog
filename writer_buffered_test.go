package golog

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"fmt"
	"io"
)

func BenchmarkWriter_Write(b *testing.B) {
	bufferedFileWriter := newBufferedWriter(os.Stdout, withBufferSize(100))
	data := []byte(`x5tT0dAb6ntcZIgn9danMXABNt8MfowerG87UgxjLODNibqt3H3zfzNZvzJr9DiLoRMqePUzot0UsgPM63mzcSQkgVOexrCpWyFR`)
	for i := 0; i < b.N; i++ {
		bufferedFileWriter.Write(data)
	}
}

func TestNewWriter(t *testing.T) {
	t.Run("expected initial value " , func(t *testing.T) {
		writer := newBufferedWriter(os.Stdout)
		assert.NotNil(t, writer)
		assert.Equal(t, 0, writer.numberOfWrittenBytes)
		assert.Equal(t, defaultBufferSize, writer.capacity())
	})

	t.Run("returns own instance " , func(t *testing.T) {
		writer := newBufferedWriter(os.Stdout)
		newWriter := newBufferedWriter(writer)
		assert.Equal(t, fmt.Sprintf("%p", writer),  fmt.Sprintf("%p", newWriter))
	})

	t.Run("returns buffer size overwritten by option parameter", func(t *testing.T) {
		capacity := 100
		writer := newBufferedWriter(os.Stdout, withBufferSize(capacity))
		expected := capacity
		assert.Equal(t, expected, writer.capacity())
	})

	t.Run("if capacity setted negative integer then default value will be used", func(t *testing.T) {
		writer := newBufferedWriter(os.Stdout, withBufferSize(-1))
		assert.Equal(t, defaultBufferSize, writer.capacity())
	})
}

func TestWriter_Write(t *testing.T) {
	
	t.Run("returns error if writer has error", func(t *testing.T) {
		writer := bufferedWriter{
			err: io.ErrShortWrite,
		}
		_, err := writer.Write([]byte(""))
		assert.NotNil(t, err)
	})

	t.Run("not flushed for the same input as capacity", func(t *testing.T) {
		input := []byte("test")
		capacity := len(input)
		bufferedFileWriter := newBufferedWriter(os.Stdout, withBufferSize(capacity))
		bufferedFileWriter.Write(input)
		expected := len(input)
		assert.Equal(t, expected, bufferedFileWriter.buffered())
	})

	t.Run("If it exceeds the capacity, it will be output directly", func(t *testing.T) {
		capacity := 4
		bufferedFileWriter := newBufferedWriter(os.Stdout, withBufferSize(capacity))
		n, err := bufferedFileWriter.Write([]byte("test123"))
		assert.Nil(t, err)
		assert.Equal(t, 0, bufferedFileWriter.buffered())
		assert.Equal(t, 7, n)
	})

	t.Run("If it exceeds the available size of buffer(capacity >= input > available), it will be flush before buffering", func(t *testing.T) {
		capacity := 10
		bufferedFileWriter := newBufferedWriter(os.Stdout, withBufferSize(capacity))
		n, err := bufferedFileWriter.Write([]byte("test123"))
		n, err = bufferedFileWriter.Write([]byte("test"))
		assert.Nil(t, err)
		assert.Equal(t, 4, bufferedFileWriter.buffered())
		assert.Equal(t, 4, n)
	})

	t.Run("If length of input is less than available, it will be buffering without flush call", func(t *testing.T) {
		capacity := 10
		bufferedFileWriter := newBufferedWriter(os.Stdout, withBufferSize(capacity))
		n, err := bufferedFileWriter.Write([]byte("test123"))
		assert.Nil(t, err)
		assert.Equal(t, 7, bufferedFileWriter.buffered())
		assert.Equal(t, 7, n)

	})

	t.Run("Buffers are gradually accumulated", func(t *testing.T) {
		input := []byte("test")
		bufferedFileWriter := newBufferedWriter(os.Stdout, withBufferSize(15))
		bufferedFileWriter.Write(input)
		assert.Equal(t, 4, bufferedFileWriter.buffered())
		assert.Equal(t, "test", string(bufferedFileWriter.buffer[0:bufferedFileWriter.buffered()]))

		bufferedFileWriter.Write(input)
		assert.Equal(t, 8, bufferedFileWriter.buffered())
		assert.Equal(t, "testtest", string(bufferedFileWriter.buffer[0:bufferedFileWriter.buffered()]))

		bufferedFileWriter.Write(input)
		assert.Equal(t, 12, bufferedFileWriter.buffered())
		assert.Equal(t, "testtesttest", string(bufferedFileWriter.buffer[0:bufferedFileWriter.buffered()]))

		bufferedFileWriter.Write(input)
		assert.Equal(t, 4, bufferedFileWriter.buffered())
		assert.Equal(t, "test", string(bufferedFileWriter.buffer[0:bufferedFileWriter.buffered()]))
	})

}

func TestWriter_Flush(t *testing.T) {

	t.Run("returns error if writer has error", func(t *testing.T) {
		writer := bufferedWriter{
			err: io.ErrShortWrite,
		}
		err := writer.Flush()
		assert.NotNil(t, err)
	})
}

func TestWriter_buffered(t *testing.T) {
	
	t.Run("returns zero", func(t *testing.T) {
		bufferedFileWriter := newBufferedWriter(os.Stdout)
		expected := 0
		actual := bufferedFileWriter.buffered()
		assert.Equal(t, expected, actual)
	})

	t.Run("returns length of inputting", func(t *testing.T) {
		input := []byte("test")
		bufferedFileWriter := newBufferedWriter(os.Stdout)
		bufferedFileWriter.Write(input)
		expected := 4
		actual := bufferedFileWriter.buffered()
		assert.Equal(t, expected, actual)
	})

	t.Run("returns zero after flush", func(t *testing.T) {
		input := []byte("test")
		bufferedFileWriter := newBufferedWriter(os.Stdout)
		bufferedFileWriter.Write(input)
		bufferedFileWriter.Flush()
		expected := 0
		actual := bufferedFileWriter.buffered()
		assert.Equal(t, expected, actual)
	})

	t.Run("returns length of third input after flush automatically", func(t *testing.T) {
		input := []byte("test")
		bufferedFileWriter := newBufferedWriter(os.Stdout, withBufferSize(10))
		bufferedFileWriter.Write(input)
		bufferedFileWriter.Write(input)
		bufferedFileWriter.Write(input) // flush and input
		expected := 4
		actual := bufferedFileWriter.buffered()
		assert.Equal(t, expected, actual)
	})
}