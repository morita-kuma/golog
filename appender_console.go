package golog

import (
	"os"
)

type Destination string
const Destination_STDOUT = "STDOUT"
const Destination_STDERR = "STDERR"

type ConsoleAppender struct {
	destination Destination
}

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

func NewDefaultConsoleAppender() ConsoleAppender {
	return NewConsoleAppender(Destination_STDOUT)
}

func NewConsoleAppender(destination Destination) ConsoleAppender {
	return ConsoleAppender{
		destination: destination,
	}
}