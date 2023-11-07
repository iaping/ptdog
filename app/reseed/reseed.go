package reseed

import (
	"fmt"
	"ptdog/app/client"
	"ptdog/app/config"
	"time"

	"github.com/rs/zerolog/log"
)

type Reseed struct {
}

func New() *Reseed {
	return &Reseed{}
}

func (r *Reseed) Run() error {
	scanners, err := r.scanners()
	if err != nil {
		return err
	}
	if len(scanners) == 0 {
		log.Warn().Msg("未配置下载器")
	} else {
		for _, scanner := range scanners {
			scanner.Run()
		}
	}

	if websites := r.websites(); len(websites) == 0 {
		log.Warn().Msg("未配置辅种站点")
	} else {
		querier.Websites(websites).Run()
		seeder.Run()
	}

	return nil
}

func (r *Reseed) websites() []*Website {
	var websites []*Website

	for _, web := range config.Conf.Websites {
		if !web.Enable {
			continue
		}
		if web.Limit <= 0 {
			web.Limit = 100
		}
		websites = append(websites, &Website{
			name:    web.Name,
			api:     web.Api,
			passkey: web.Passkey,
			limit:   web.Limit,
			domain:  web.Domain,
		})
	}

	return websites
}

func (r *Reseed) scanners() ([]*Scanner, error) {
	var scanners []*Scanner
	for _, conf := range config.Conf.Clients {
		if !conf.Enable {
			continue
		}
		client, err := r.client(conf)
		if err != nil {
			return nil, err
		}
		scanner := NewScanner(client, conf.Dir, config.Conf.System.SleepDuration())
		scanners = append(scanners, scanner)
	}
	return scanners, nil
}

func (r *Reseed) client(config config.Client) (client.IClient, error) {
	if config.Timeout <= 0 {
		config.Timeout = 60
	}
	timeout := time.Duration(config.Timeout) * time.Second

	conf := client.Config{
		Https:        config.Https,
		Host:         config.Host,
		Username:     config.Username,
		Password:     config.Password,
		Port:         config.Port,
		Timeout:      timeout,
		SkipChecking: config.SkipChecking,
	}

	switch config.Type {
	case client.TypeTransmission:
		return client.NewTransmission(conf)
	case client.TypeQbittorrent:
		return client.NewQbittorrent(conf)
	}

	return nil, fmt.Errorf("the type of the client is %d, not supported", config.Type)
}
