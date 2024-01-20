// Package server
package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/webtoor/go-rest-api/internal/app"
	"github.com/webtoor/go-rest-api/router"
)

type httpServer struct {
	cfg  *app.Config
	appx *fiber.App
}

func NewHTTPServer(cfg *app.Config) Server {
	appx := fiber.New(app.NewFiberConfig())
	return &httpServer{
		cfg:  cfg,
		appx: appx,
	}
}

func (h *httpServer) Run(ctx context.Context) error {
	var err error

	router.NewRouter(h.cfg, h.appx).Route()

	go func() {
		err := h.appx.Listen(fmt.Sprintf(":%d", h.cfg.App.Port))
		if err != nil {
			logrus.Warn("starting http services failed error:", err)
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer func() {
		cancel()
	}()

	if err = h.appx.ShutdownWithContext(ctxShutDown); err != nil {
		logrus.Fatal("service http stopped failed error:", err)
	}

	logrus.Info("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}

func (h *httpServer) Done() {
	logrus.Info("service http stopped")
}
