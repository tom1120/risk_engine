package log

import (
	"context"
	"fmt"
	"io"
	"log"
)

type defaultLogger struct {
	*log.Logger
	writer [LevelMax]outputFn
	ctx    context.Context
}

func NewDefaultLogger(writer io.Writer, prefix string, flag int, level Level) *defaultLogger {
	l := &defaultLogger{}
	l.Logger = log.New(writer, prefix, flag)
	for i := int(LevelError); i < int(LevelMax); i++ {
		if i <= int(level) {
			l.writer[i] = l.Output
		} else {
			l.writer[i] = dropOutput
		}
	}
	return l
}

type outputFn func(calldepth int, s string) error

func dropOutput(calldepth int, s string) error {
	return nil
}

func header(prefix, msg string) string {
	return fmt.Sprintf("%s: %s", prefix, msg)
}

func (l *defaultLogger) Error(v ...interface{}) {
	l.writer[int(LevelError)](calldepth, header(LevelError.String(), fmt.Sprint(v...)))
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	l.writer[int(LevelError)](calldepth, header(LevelError.String(), fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) Warn(v ...interface{}) {
	l.writer[int(LevelWarn)](calldepth, header(LevelWarn.String(), fmt.Sprint(v...)))
}

func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	l.writer[int(LevelWarn)](calldepth, header(LevelWarn.String(), fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) Info(v ...interface{}) {
	l.writer[int(LevelInfo)](calldepth, header(LevelInfo.String(), fmt.Sprint(v...)))
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	l.writer[int(LevelInfo)](calldepth, header(LevelInfo.String(), fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) Debug(v ...interface{}) {
	l.writer[int(LevelDebug)](calldepth, header(LevelDebug.String(), fmt.Sprint(v...)))
}

func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	l.writer[int(LevelDebug)](calldepth, header(LevelDebug.String(), fmt.Sprintf(format, v...)))
}
