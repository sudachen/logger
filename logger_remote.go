package logger

import (
	"fmt"
	"github.com/google/logger/internal"
)

func Rinfo(a string) {
	Info(a)
	internal.Info(a)
}

func Rinfof(a string, b... interface{}) {
	Rinfo(fmt.Sprintf(a, b...))
}

func ConnectSentry(dsn string) error {
	return internal.ConnectSentry(dsn)
}
