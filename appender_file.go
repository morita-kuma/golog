package golog

import (
	"bufio"
	"os"
	"bytes"
	"io"
)

type FileAppender struct {
	file         *os.File
	buffedWriter *bufio.Writer
}

func (appender *FileAppender) Write(data []byte) (n int, err error) {
	return appender.buffedWriter.Write(data)
}

func (appender *FileAppender) Close() error {

	// Flush Buffer
	if err := appender.buffedWriter.Flush(); err != nil {
		return err
	}

	// Close File
	if err := appender.file.Close(); err != nil {
		return err
	}

	return nil
}

func NewFileAppender(fileName string) (fileAppender *FileAppender, err error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	return &FileAppender{
		buffedWriter: bufio.NewWriterSize(f, 4096),
	}, nil
}

type AsyncFileAppender struct {
	file *os.File
	buffedWriter *bufio.Writer
}

func NewAsyncFileAppender(fileName string) (asyncFileAppender *AsyncFileAppender, err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return &AsyncFileAppender{
		file:         file,
		buffedWriter: bufio.NewWriterSize(file, 4096),
	}, nil
}

func (appender *AsyncFileAppender) Write(data []byte) (n int, err error) {
	return appender.buffedWriter.Write(data)
}

func (appender *AsyncFileAppender) Flush() {

}

func (appender *AsyncFileAppender) ReadFrom(r io.Reader) (n int64, err error) {
	return appender.buffedWriter.ReadFrom(r)
}

func (appender *AsyncFileAppender) Close() {
	appender.file.Close()
}

// AsyncFileAppender
type AsyncFileAppender2 struct {
	// file
	file *os.File

	// buffer
	buffer *bytes.Buffer

	// readBuffer
	readBuffer chan []byte

	// capacity
	capacity int
}

func NewAsyncFileAppender2(fileName string) (asyncFileAppender *AsyncFileAppender2, err error) {

	file, err := os.Create(fileName)
	//file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	appender := &AsyncFileAppender2{
		file:       file,
		buffer:     new(bytes.Buffer),
		readBuffer: make(chan []byte, 1024),
		capacity:   8192,
	}

	appender.Listener()

	return appender, nil
}

func (appender *AsyncFileAppender2) Write(data []byte) (n int, err error) {

	b := make([]byte, len(data))

	copy(b, data)

	if appender != nil {
		appender.readBuffer <- b
	}

	return len(b), nil
}

func (appender *AsyncFileAppender2) Listener() {

	go func(appender2 *AsyncFileAppender2) {

		for {

			select {
			case buf := <-appender2.readBuffer:

				if len(buf)+appender2.buffer.Len() > appender2.capacity {

					appender.Flush()
					appender.buffer.Reset()
				} else {

					appender.buffer.Write(buf)
				}
			}
		}

	}(appender)
}

func (appender *AsyncFileAppender2) Flush() {
	appender.file.Write(appender.buffer.Bytes())
}

func (appender *AsyncFileAppender2) Close() {
	appender.file.Close()
}
