package http

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/webtoor/go-rest-api/internal/app"
	"github.com/webtoor/go-rest-api/server"
)

func Start(ctx context.Context, cfg *app.Config) {
	serve := server.NewHTTPServer(cfg)
	defer serve.Done()
	logrus.Info("starting http services...")

	if err := serve.Run(ctx); err != nil {
		logrus.Info("service http stopped error: ", err.Error())
	}
}
