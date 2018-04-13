package golog

type FluentAppender struct {
}

func (FluentAppender) Write(data []byte) (n int, err error)  {
	return 0, nil
}
