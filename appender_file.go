package golog

import (
	"bufio"
	"os"
	"io"
	"sync"
)

// FileAppender
type FileAppender struct {
	file         *os.File
	buffedWriter *bufio.Writer

	activated bool
}

// Write implements io.Writer
func (appender *FileAppender) Write(data []byte) (n int, err error) {
	return appender.buffedWriter.Write(data)
}

// Close implements io.Closer
func (appender *FileAppender) Close() error {

	if appender.activated {

		// Flush Buffer
		if err := appender.buffedWriter.Flush(); err != nil {

			return err
		}

		if err := appender.file.Close(); err != nil {

			return err
		}

		appender.activated = false
	}

	// Close File
	return appender.file.Close()
}

// NewFileAppender returns new FileAppender
func NewFileAppender(fileName string) (fileAppender *FileAppender, err error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return &FileAppender{
		buffedWriter: bufio.NewWriterSize(f, 4096),
		activated:    true,
	}, nil
}

// BufferedFileAppender
type BufferedFileAppender struct {
	file           *os.File
	bufferedWriter *bufio.Writer
	mu             *sync.Mutex
}

// NewBufferedFileAppender returns new BufferedFileAppender
func NewBufferedFileAppender(fileName string) (asyncFileAppender *BufferedFileAppender, err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return &BufferedFileAppender{
		file:           file,
		bufferedWriter: bufio.NewWriterSize(file, 4096),
		mu:             new(sync.Mutex),
	}, nil
}

// Write implements io.Write
func (appender *BufferedFileAppender) Write(data []byte) (n int, err error) {
	appender.mu.Lock()
	defer appender.mu.Unlock()
	return appender.bufferedWriter.Write(data)
}

// ReadFrom implements io.ReadFrom
func (appender *BufferedFileAppender) ReadFrom(r io.Reader) (n int64, err error) {
	return appender.bufferedWriter.ReadFrom(r)
}

// Close implements io.Closer
func (appender *BufferedFileAppender) Close() error {
	appender.mu.Lock()
	defer appender.mu.Unlock()

	if err := appender.bufferedWriter.Flush(); err != nil {
		return err
	}

	return appender.file.Close()
}