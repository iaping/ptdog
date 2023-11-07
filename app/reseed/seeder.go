package reseed

import (
	"ptdog/app/client"

	"github.com/gookit/cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	cache.Register(cache.DvrFile, cache.NewFileCache(".cache"))
	cache.DefaultUse(cache.DvrFile)
}

type Seed struct {
	client  client.IClient
	website *Website
	id      int
	torrent *client.Torrent
}

func (s Seed) url() string {
	return s.website.FormatDownload(s.id)
}

func (s Seed) log(event *zerolog.Event) *zerolog.Event {
	event.Str("下载器", s.client.String())
	event.Str("站点", s.website.String())
	event.Any("ID", s.id)
	event.Any("资源", s.torrent.DownloadPath)
	return event
}

var seeder = &Seeder{
	queue: make(chan Seed, 8),
}

type Seeder struct {
	queue chan Seed
}

func (s *Seeder) Push(seed Seed) {
	s.queue <- seed
}

func (s *Seeder) Run() {
	go func() {
		for seed := range s.queue {
			go s.handler(seed)
		}
	}()
}

func (s *Seeder) handler(seed Seed) {
	url := seed.url()
	if cache.Has(url) {
		return
	}

	if err := seed.client.TorrentAdd(url, seed.torrent.DownloadPath); err != nil {
		seed.log(log.Err(err)).Msg("辅种失败")
		return
	}

	if err := cache.Set(url, true, cache.Forever); err != nil {
		log.Err(err).Msg("缓存失败")
	}

	seed.log(log.Info()).Msg("辅种成功")
}
