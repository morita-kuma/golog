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
logger.InfoF("%s", "message")
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
logger.InfoJ(message)
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
logger.InfoS(logEvent)
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

## 4.1. ConsoleAppender
## 4.2. BufferAppender
## 4.3. FileAppender
## 4.4. BufferedFileAppender

# 5. CustomLogAppender
ログアペンダーは、golangのio.WriteCloserのエイリアスです。サポートされているログアペンダーを拡張したい場合や、独自仕様を利用したい場合はこのインターフェースを実装してください。
```
type Appender = io.WriteCloser
```


# 5. Metadata
# 6. MetadataFormatter

# 6. Performance