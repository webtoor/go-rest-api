package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/webtoor/go-rest-api/cmd/http"
	"github.com/webtoor/go-rest-api/internal/app"
	"github.com/webtoor/go-rest-api/pkg/logger"
)

func Start() {
	cfg := app.NewConfig()
	logger.Setup(cfg)
	rootCmd := &cobra.Command{}

	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	cmd := []*cobra.Command{
		{
			Use:   "http",
			Short: "Run HTTP Service",
			Run: func(cmd *cobra.Command, args []string) {
				http.Start(ctx, cfg)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
