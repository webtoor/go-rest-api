package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/webtoor/go-rest-api/internal/app"
)

func Setup(cfg *app.Config) {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)

	if cfg.App.Debug {
		logrus.SetLevel(logrus.TraceLevel)
		return
	}

	lvl, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		lvl = logrus.InfoLevel
	}

	logrus.SetLevel(lvl)
}
