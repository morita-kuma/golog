package golog

type LogLevelFilter interface {
	DoFilter(logLevel LogLevel) LogLevels
}

// DefaultLogLevelFilter
type DefaultLogLevelFilter struct {
}

func (DefaultLogLevelFilter) DoFilter(logLevel LogLevel) LogLevels {
	var filteredLogLevels LogLevels
	for k, v := range logLevelMap {
		if k >= logLevel.TypeVal() {
			filteredLogLevels = append(filteredLogLevels, v)
		}
	}
	return filteredLogLevels
}

func NewDefaultLevelFilter () DefaultLogLevelFilter {
	return DefaultLogLevelFilter{}
}