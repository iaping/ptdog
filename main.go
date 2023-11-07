package main

import (
	"os"
	"ptdog/app"
	"ptdog/app/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stdout,
		// TimeFormat: time.RFC3339,
	})
}

func main() {
	if err := config.Load("config.json"); err != nil {
		log.Panic().Err(err).Msg("配置加载失败")
	}

	if err := app.New().Run(); err != nil {
		log.Panic().Err(err).Send()
	}
}
