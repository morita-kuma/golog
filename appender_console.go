package golog

import (
	"os"
)

type Destination string
const Destination_STDOUT = "STDOUT"
const Destination_STDERR = "STDERR"

// ConsoleAppender
type ConsoleAppender struct {
	destination Destination
}

// Write implements io.Writer
func (appender ConsoleAppender) Write(data []byte) (n int, err error) {
	data = append(data, []byte("\n")...)

	switch appender.destination {
	case Destination_STDERR :
		os.Stderr.Write(data)

	case Destination_STDOUT :
		os.Stdout.Write(data)

	default:
		os.Stdout.Write(data)
	}
	return 0,nil
}

// Close implements io.Closer
func (appender ConsoleAppender) Close() error {
	return nil
}

// NewDefaultConsoleAppender returns new ConsoleAppender
// default io.writer is used to os.Stdout
func NewDefaultConsoleAppender() ConsoleAppender {
	return NewConsoleAppender(Destination_STDOUT)
}

// NewConsoleAppender returns new ConsoleAppender
func NewConsoleAppender(destination Destination) ConsoleAppender {
	return ConsoleAppender{
		destination: destination,
	}
}