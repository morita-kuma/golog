package golog

import (
	"testing"
)


func BenchmarkAsyncFileAppender_Write(b *testing.B) {
	appender, _ := NewFileAppender("./log/dat1")
	data := []byte("abcdefghijklmn")
	for i := 0; i < b.N; i++ {
		appender.Write(data)
	}

}