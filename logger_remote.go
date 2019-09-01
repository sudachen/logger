package logger

import (
	"fmt"
	"github.com/sudachen/logger/internal"
)

func Rinfo(a string) {
	defaultLogger.output(sInfo, 0, a)
	internal.Info(a)
}

func Rinfof(a string, b ...interface{}) {
	t := fmt.Sprintf(a, b...)
	defaultLogger.output(sInfo, 0, t)
	internal.Info(a)
}

func ConnectSentry(dsn string) error {
	return internal.ConnectSentry(dsn)
}
