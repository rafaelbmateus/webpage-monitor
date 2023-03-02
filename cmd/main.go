package main

import (
	"context"
	"os"

	"monitor/config"
	"monitor/monitor"

	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	cfg, err := config.Load(".config.yaml")
	if err != nil {
		log.Fatal().Msgf("error on load config file %q", err)

		return
	}

	notify := monitor.NewSlackNotify(cfg.Title, os.Getenv("SLACK_WEBHOOK_URL"))
	monitor.NewMonitor(ctx, &log, notify).Run(cfg)

	select {} // block forever
}
