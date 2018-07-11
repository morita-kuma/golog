package bufio

import (
	"io"
)

var defaultBufferSize = 4096

// Writer implements buffering for an io.Writer object.
// If an error occurs writing to a Writer, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.Writer.
type Writer struct {
	err    error
	buffer []byte
	writer io.Writer
	numberOfWrittenBytes int
}

type WriterOption func(writer *Writer)

// NewWriterSize returns a new Writer whose buffer has at least the specified
// size. If the argument io.Writer is already a Writer with large enough
// size, it returns the underlying Writer.
func NewWriter(writer io.Writer, writerOptions ...WriterOption) *Writer {

	if bufferedWriter, ok := writer.(*Writer); ok {
		return bufferedWriter
	}

	bufferedWriter := &Writer{
		buffer: make([]byte, defaultBufferSize),
		writer: writer,
	}

	for _, option := range writerOptions {
		option(bufferedWriter)
	}

	return bufferedWriter
}

func WithBufferSize(size int) WriterOption {
	return func(writer *Writer) {
		if size <= 0 {
			writer.buffer = make([]byte, defaultBufferSize)
			return
		}
		writer.buffer = make([]byte, size)
	}
}

// Write writes the contents of p into the buffer.
// It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining
// why the write is short.
func (w *Writer) Write(data []byte) (n int, err error) {

	if w.hasError() {
		return n, w.err
	}

	if len(data) > w.capacity() {
		n, w.err = w.writer.Write(data)
		return
	}

	if len(data) > w.available() {
		err = w.Flush()
		n = copy(w.buffer[0:len(data)], data)
		w.numberOfWrittenBytes += n
		return
	}

	n = copy(w.buffer[w.buffered():], data)
	w.numberOfWrittenBytes += n
	return
}

// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush() error {

	if w.hasError() {
		return w.err
	}

	if w.buffered() == 0 {
		return nil
	}

	n, err := w.writer.Write(w.buffer[0:w.buffered()])
	if n < w.buffered() && err == nil {
		err = io.ErrShortWrite
	}

	if err != nil {
		if n > 0 && n < w.buffered() {
			copy(w.buffer[0:w.buffered()-n], w.buffer[n:w.buffered()])
		}
		w.numberOfWrittenBytes -= n
		w.err = err
		return err
	}

	w.numberOfWrittenBytes = 0
	return nil
}

func (w *Writer) capacity() int {
	return cap(w.buffer)
}

// buffered returns the number of bytes that have been written into the current buffer.
func (w *Writer) buffered() int {
	return w.numberOfWrittenBytes
}

// available returns how many bytes are unused in the buffer.
func (w *Writer) available() int {
	return w.capacity() - w.buffered()
}

func (w *Writer) hasError() bool {
	return w.err != nil
}