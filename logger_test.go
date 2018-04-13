package golog

import (
	"testing"
	"fmt"
	"os"
	"log"
)

func TestNewLogger(t *testing.T) {

	func () {
		logger := NewLogger("testLogger", LogLevel_TRACE)
		appender := NewByteBufferAppender()
		logger.SetAppender(appender)
		logger.Debug("my first logger0")
		logger.Debug("my first logger1")
		logger.Debug("my first logger2")
		logger.Debug("my first logger3")
		fmt.Println(appender.String())
	}()

	func () {
		logger := NewLogger("testLogger", LogLevel_TRACE)
		logger.SetAppender(ConsoleAppender{destination:Destination_STDERR})
		logger.Debug("hogehoge")
	}()

	// os.Stdout
	func () {
		logger := NewLogger("testLogger", LogLevel_TRACE)
		logger.SetAppender(os.Stdout)
		logger.Debug("fugafuga\n")
	}()

	// os.File
	func () {
		f, err := os.Create("./log/dat2")

		if err != nil {
			fmt.Print(err)
		}

		logger := NewLogger("testLogger", LogLevel_TRACE)
		logger.SetAppender(f)
		logger.Debug("write to file1\n")
		logger.Debug("write to file2\n")
		logger.Debug("write to file3\n")
	}()

	func () {
		logger := NewLogger("testLogger", LogLevel_TRACE)
		appender := NewByteBufferAppender()
		logger.SetAppender(appender)
		logger.InfoS(JsonLogEvent{
			event: struct{
				Name string `json:"name"`
				Address string `json:"address"`
			}{
				Name:"name_value",
				Address:"address_value",
			},
		})
		fmt.Println(appender.String())
	}()

	func () {
		logger := NewLogger("testLogger",LogLevel_TRACE)
		logger.SetAppender(NewDefaultConsoleAppender(), NewDefaultConsoleAppender())
		logger.TraceF("value = %v", 10)
	}()
}

func TestNewAsyncFileAppender(t *testing.T) {
	appender,_ := NewAsyncFileAppender("./log/dat2")
	logger := NewDefaultLogger()
	logger.SetAppender(appender)
	logger.Info("x")
}

func BenchmarkLogger_Info_ontime(b *testing.B) {
	f, err := os.Create("./log/dat1")
	if err != nil {
		fmt.Print(err)
	}

	logger := NewDefaultLogger()
	logger.SetAppender(f)

	data := "abcdefghijklmn"
	for i := 0; i < b.N; i++ {
		logger.Info(data)
	}
}

func BenchmarkLogger_Json_to_buffer(b *testing.B) {
	appender,_ := NewAsyncFileAppender("./log/dat2")
	logger := NewDefaultLogger()
	logger.SetAppender(appender)
	logger.DisableLogEventMetadata()

	obj := struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}{
		Name:    "name",
		Address: "address",
	}

	for i := 0; i< b.N; i++ {
		logger.doAppendIfLevelEnabled(JsonLogEvent{
			event:obj,
		}, LogLevel_INFO)
	}
}

func BenchmarkLogger_text_to_buffer(b *testing.B) {
	appender,_ := NewAsyncFileAppender("./log/dat2")
	logger := NewDefaultLogger()
	logger.SetAppender(appender)
	logger.DisableLogEventMetadata()

	for i := 0; i< b.N; i++ {
		logger.Info("hoge")
	}
}

func BenchmarkLogger_text_to_buffer2(b *testing.B) {
	appender,_ := NewAsyncFileAppender("./log/dat2")
	logger := NewDefaultLogger()
	logger.SetAppender(appender)
	logger.DisableLogEventMetadata()

	for i := 0; i< b.N; i++ {
		logger.Info("hoge")
	}
}

func BenchmarkLogger_text_to_buffer3(b *testing.B) {
	appender,_ := NewAsyncFileAppender("./log/dat2")
	logger := NewDefaultLogger()
	logger.SetAppender(appender)

	for i := 0; i< b.N; i++ {
		logger.Info("hoge")
	}
}

func BenchmarkLogger_default(b *testing.B) {
	appender,_ := NewAsyncFileAppender("./log/dat_x")
	log.SetOutput(appender)

	for i :=0; i<b.N; i++ {
		log.Print("xxxxxxx")
	}
}


/*
func ExampleLogger_Info() {
	logger := NewDefaultLogger()
	logger.info(MyTextLogEvent {
		TextLogEvent: TextLogEvent{
			Event:"overwrite_format",
		},
	})

	// Output:
	// [overwrite_loglevel] [overwrite_time] [overwrite_loggername] [overwrite_sourcefile]([overwrite_sourceline]) overwrite_format
}*/
