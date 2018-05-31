# 0. Overview

以下の特徴を持つgolang用のロガー実装です.

- 拡張可能な設計 (カスタムアペンダー)
- アロケーションフリー(特定の条件下で)で高速に動作します
- ログレベル別出力
- ログイベントの複数出力

# 1. LogEvent
ログイベントは、対象のログイベントに対するエンコーディング方式を定義しています。デフォルトで定義されているログイベントは、以下の通りです。

## 1.1. TextLogEvent
```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
logger.Info("message")
```

## 1.2. FormatLogEvent
```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
logger.Infof("%s", "message")
```


## 1.3. JsonLogEvent

```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
message := struct{
	Name string `json:"name"`
	Address string `json:"address"`
}{
    Name:"name_value",
	Address:"address_value",
}
logger.Infof(message)
```


# 2. CustomLogEvent
デフォルトのログイベントに必要な実装が無くても、多くの場合はstringerを実装することで要件を満たせるはずです。
しかしながら、メタデータのようにロガー内部で生成される値をハンドリングすることは難しいです。この場合は、
ログイベント自体を以下のようにLogEventMetadataを受け取り、[]byteを返却するインターフェースを実装してください。
```
type LogEvent interface {
	Encode(metadata *LogEventMetadata) []byte
}
```

以下は、sliceをフィールドに持つstructにログイベントを実装した例です。
フィールドをcsv形式に変換し、golang.TextLogEventにmetadataをフォーマットする処理を移譲します。
```
type CustomLogEvent struct {
	fields []string
}

func (e CustomLogEvent) Encode(metadata *golog.LogEventMetadata) []byte {
	if len(e.fields) > 0 {
		logEvent := golog.TextLogEvent{
			Event:strings.Join(e.fields, ","),
		}
		return logEvent.Encode(metadata)
	}
	return []byte{}
}
```

Example: 
```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
logEvent := CustomLogEvent{
	fields:[]string{"a", "b", "c"},
}
logger.SInfo(logEvent)
```

Result:
```
[INFO] 2018-05-06T22:01:14+09:00 defaultLogger test.go(141) a,b,c
```


# 3. LogEvent
ログレベルは、TRACE, DEBUG, INFO, WARN, ERROR, FATALの6通りに対応しています。デフォルトのロガーでは、全てのレベルが出力されます。

## 3.1. デフォルト設定

| LogLevel | OUTPUT |
| :---: | :---: |
| TRACE | ○ |
| DEBUG | ○ |
| INFO | ○ |
| WARN | ○ |
| ERROR | ○ |
| FATAL | ○ |

## 3.2. 特定のレベル以上を出力する
環境によってログ出力レベルを抑制したい場合には、ロガーの初期化時にログレベルを指定してください。
例えば、WARNレベル以上のログを出力したい場合には、以下のように初期化をします。

| LogLevel | OUTPUT |
| :---: | :---: |
| TRACE | - |
| DEBUG | - |
| INFO | - |
| WARN | ○ |
| ERROR | ○ |
| FATAL | ○ |

Example: 
```
logger := golog.NewLogger("testLogger", golog.LogLevel_WARN)
logger.SetAppender(golog.NewDefaultConsoleAppender())
logger.Trace("message")
logger.Debug("message")
logger.Info("message")
logger.Warn("message")
logger.Error("message")
logger.Fatal("message")
```

Result:
```
[WARN] 2018-05-06T22:14:47+09:00 testLogger test.go(166) message
[ERROR] 2018-05-06T22:14:47+09:00 testLogger test.go(166) message
[FATAL] 2018-05-06T22:14:47+09:00 testLogger test.go(166) message
```

# 4. LogAppender
LogAppenderは、LogEventの出力先を実装します。
1つのLogEventに対して複数の出力先が必要な場合は、以下のように実装することも可能です。

```
logger := golog.NewDefaultLogger()
logger.SetAppender(
	golog.NewDefaultConsoleAppender(),
	golog.NewByteBufferAppender(), )
logger.Info("message")
```


## 4.1. ConsoleAppender
LogEventを標準出力もしくは標準エラー出力に出力します。

Example:
```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
logger.Info("message")
```

Result:
```
[INFO] 2018-05-07T12:14:20+09:00 defaultLogger test.go(204) message
```

## 4.2. BufferAppender
LogEventをバッファに出力します。 LogEventのテストなどに利用してください。
バッファの内容は、String()で取得することができます。

Example:
```
logger := golog.NewDefaultLogger()
appender := golog.NewByteBufferAppender()
logger.SetAppender(appender)
logger.Info("message")
fmt.Print(appender.String())
```

Result:
```
[INFO] 2018-05-07T12:17:52+09:00 defaultLogger test.go(231) message
```

## 4.3. FileAppender
LogEventをファイルに出力します。内部で4096Byteをデフォルトでバッファし、容量超えた場合にFlushを行います。
また、バッファの内容はAppenderもしくはLoggerのCloserによってFlushされます。
アプリケーションの終了時やpanicリカバリーのタイミングでcloseを実行し、バッファが確実に出力されるように実装してください。

Example:
```
logger := golog.NewDefaultLogger()
appender, _ := golog.NewFileAppender("./log/dat2")
logger.SetAppender(appender)
logger.Info("message1")
logger.Info("message2")
logger.Info("message3")
logger.Close()
```

Result:
```
[INFO] 2018-05-07T12:19:00+09:00 defaultLogger test.go(215) message1
[INFO] 2018-05-07T12:19:00+09:00 defaultLogger test.go(215) message2
[INFO] 2018-05-07T12:19:00+09:00 defaultLogger test.go(215) message3
```

# 5. CustomLogAppender
LogAppenderは、golangのio.WriteCloserのエイリアスとして実装されています。
従って、このインターフェースを満たす既存の実装はそのまま利用することができます。
加えて、サポートされているLogAppenderを拡張したい場合や、独自仕様を利用したい場合はこのインターフェースを実装してください。
```
type Appender = io.WriteCloser
```


# 6. Metadata
Library内で生成されるログのMetadataで、LogEventで自由に整形することができます。
サポートされているのは、以下の通りです。

| metadata | type | expla | example |
| :--- | :--- | :--- | :--- |
| LogLevel | string | ログレベル | \[INFO\] |
| Time | int64 | unixtime seconds | 2018-05-07T12:19:00+09:00 |
| SourceFile | string | ログを出力したファイル名 | test.go |
| SourceLine | int | ログを出力したソースのLine |  (100) |
| LoggerName | string | Logger生成時に指定したロガー名 | defaultLogger |

```
[INFO] 2018-05-07T12:19:00+09:00 defaultLogger test.go(215) message3
```

## 6.1. Metadataを無効にする
Metadataが不要な場合や、Loggerのperformanceを向上させたい場合は以下のように無効にしてください。

Example:
```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
logger.DisableLogEventMetadata()
logger.Info("message")
```
Result:
```
message
```

## 6.2. 指定のMetadataを無効にする
指定のMetadataのみ必要な場合は、以下のように設定してください。

Example:
```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
logger.SetMetadataConfig(&golog.MetadataConfig{
	IsEnabledLogLevel:   false,
	IsEnabledTime:       true,
	IsEnabledSourceFile: true,
	IsEnabledSourceLine: true,
	IsEnabledLoggerName: false,
})
logger.Info("message")
```

Result:
```
 2018-05-07T12:37:19+09:00  test.go(250) message
```

# 6.3. Metadataのデフォルトフォーマットを上書きする
Metadataのfieldがどのようにレイアウトされるかは、LogEventの実装依存ですが、
フォーマットは何も指定しない場合はdefaultFormatterが使用されます。
フォーマッターを上書きしたい場合には、以下のように実装してください。

Example: ISO3166 -> UnixTime
```
logger := golog.NewDefaultLogger()
logger.SetAppender(golog.NewDefaultConsoleAppender())
formatter := golog.NewDefaultMetadataFormatter()
formatter.TimeFormatter = func(unixTime golog.UnixTime) string {
	return strconv.FormatInt(int64(unixTime),10)
}
logger.SetMetadataFormatter(&formatter)
logger.Info("message")
```

Result:
```
[INFO] 1525665335 defaultLogger test.go(273) message
```



# 7. Performance