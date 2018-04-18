package golog

import (
	"testing"
)


func BenchmarkFileAppender_Write(b *testing.B) {
	appender, _ := NewFileAppender("./log/dat2")
	data := []byte("abcdefghijklmn")
	for i := 0; i < b.N; i++ {
		appender.Write(data)
	}
}

func BenchmarkAsyncFileAppender_Write(b *testing.B) {
	appender, _ := NewBufferedFileAppender("./log/dat1")
	data := []byte("abcdefghijklmn")
	for i := 0; i < b.N; i++ {
		appender.Write(data)
	}

}