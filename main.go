package main

import (
	"fmt"
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
	defer func() {
		fmt.Println("Press 'Enter' to continue...")
		fmt.Scanln()
	}()

	if err := config.Load("config.json"); err != nil {
		log.Err(err).Str("配置", "config.json").Msg("配置加载失败")
		return
	}

	if err := app.New().Run(); err != nil {
		log.Err(err).Send()
	}
}
