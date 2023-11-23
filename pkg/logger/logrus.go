package logger

import (
	"context"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
)

type Logg struct {
	*logrus.Entry
}

func GetLogger() *Logg {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.JSONFormatter{
		PrettyPrint: true,
	}

	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	return &Logg{logrus.NewEntry(l)}
}

func (l *Logg) Log(ctx context.Context, level logging.Level, s string, a ...any) {
	l.Info(s)
}
