package golog

import (
	"os"
	"sync"

)

// defaultBufferSize
const defaultBufferSize = 4096

// FileAppender
type FileAppender struct {
	file           *os.File
	bufferedWriter *bufferedWriter
	mu             *sync.Mutex
	activated      bool
}

// NewFileAppender returns new FileAppender
func NewFileAppender(fileName string) (asyncFileAppender *FileAppender, err error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileAppender{
		file:           file,
		bufferedWriter: newBufferedWriter(file, withBufferSize(defaultBufferSize)),
		mu:             new(sync.Mutex),
		activated:      true,
	}, nil
}

// NewFileAppender returns new FileAppender
func NewFileAppenderWithBufferSize(fileName string, bufferSize int) (asyncFileAppender *FileAppender, err error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	size := defaultBufferSize
	if bufferSize > 0 {
		size = bufferSize
	}

	return &FileAppender{
		file:           file,
		bufferedWriter: newBufferedWriter(file, withBufferSize(size)),
		mu:             new(sync.Mutex),
		activated:      true,
	}, nil
}

// Write implements io.Write
func (appender *FileAppender) Write(data []byte) (n int, err error) {
	appender.mu.Lock()
	defer appender.mu.Unlock()
	data = append(data, '\n')
	return appender.bufferedWriter.Write(data)
}

// Close implements io.Closer
func (appender *FileAppender) Close() error {
	appender.mu.Lock()
	defer func() {
		appender.mu.Unlock()
		appender.activated = false
	}()

	if appender.activated {

		if err := appender.bufferedWriter.Flush(); err != nil {
			return err
		}

		return appender.file.Close()
	}

	return nil
}