package app

import (
	"ptdog/app/http"
	"ptdog/app/reseed"

	"github.com/rs/zerolog/log"
)

const (
	Version = "v1.0.0"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (app *App) Run() (err error) {
	app.info()
	if err := reseed.New().Run(); err != nil {
		return err
	}
	return http.NewServer().Run()
}

func (app *App) info() {
	log.Info().Msgf("PTDog %s", Version)
	log.Info().Msg("PT站点自动辅种工具，支持开放pieces_hash查询的站点。未适配的站点请联系！")
	log.Info().Msg("QQ群: 881030035")
	log.Info().Msg("Telegram: https://t.me/+ibBCW1uE4Zs5ODk1")
}
