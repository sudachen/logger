package internal

import (
	"github.com/getsentry/sentry-go"
	"sync"
	"time"
)

const flashTimeout = 3*time.Second

var connected = false
var ErrorLog = &snio{ sentry.LevelError }
var WarnLog = &snio{ sentry.LevelWarning }
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
	level sentry.Level
}

func (sn *snio) Write(p []byte) (n int, err error) {
	if connected {
		wg.Add(1)
		go func() {
			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetLevel(sn.level);
				sentry.CaptureMessage(string(p))
			})
			sentry.Flush(flashTimeout)
			wg.Done()
		}()
	}
	return 0,nil
}

func Info(text string) {
	if connected {
		wg.Add(1)
		go func() {
			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelInfo);
				sentry.CaptureMessage(text)
			})
			sentry.Flush(flashTimeout)
			wg.Done()
		}()
		Wait()
	}
}

func Wait() {
	wg.Wait()
}
