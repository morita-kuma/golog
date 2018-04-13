package golog

import "io"

// Appender
type Appender interface {
	io.Writer
}
