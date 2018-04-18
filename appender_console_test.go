package golog

func ExampleConsoleAppender_Write() {

	appender := NewDefaultConsoleAppender()
	appender.Write([]byte("test1"))
	appender.Write([]byte("test2"))
	appender.Write([]byte("test3"))

	// Output:
	//test1
	//test2
	//test3
}