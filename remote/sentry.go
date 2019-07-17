package remote

import (
	"github.com/getsentry/sentry-go"
	"sync"
)

var connected = false
var ErrorLog = &snio{ isErrLog: true }
var WarnLog = &snio{ false }
var wg sync.WaitGroup

func ConnectSentry(dsn string) error {
	if err := sentry.Init(sentry.ClientOptions{Dsn: dsn}); err != nil {
		return err
	} else {
		connected = true
		return nil
	}
}

type snio struct {
	isErrLog bool
}

func (sn *snio) Write(p []byte) (n int, err error) {
	if connected {
		wg.Add(1)
		go func() {
			sentry.WithScope(func(scope *sentry.Scope) {
				if (sn.isErrLog) {
					scope.SetLevel(sentry.LevelError);
				} else {
					scope.SetLevel(sentry.LevelWarning);
				}
				sentry.CaptureMessage(string(p))
			})
			wg.Done()
		}()
	}
	return 0,nil
}
