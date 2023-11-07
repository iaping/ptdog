package client

import (
	"context"
	"fmt"

	"github.com/hekmon/transmissionrpc/v2"
)

type Transmission struct {
	conf Config
	c    *transmissionrpc.Client
}

func NewTransmission(conf Config) (*Transmission, error) {
	c, err := transmissionrpc.New(conf.Host, conf.Username, conf.Password, &transmissionrpc.AdvancedConfig{
		HTTPS:       conf.Https,
		Port:        conf.Port,
		HTTPTimeout: conf.Timeout,
	})
	if err != nil {
		return nil, err
	}

	return &Transmission{
		conf: conf,
		c:    c,
	}, nil
}

func (t *Transmission) Type() Type {
	return TypeTransmission
}

func (t *Transmission) String() string {
	return fmt.Sprintf("%s: %s:%d", t.Type().Name(), t.conf.Host, t.conf.Port)
}

func (t *Transmission) Torrents(hashes []string) ([]*Torrent, error) {
	ctx, cencel := context.WithTimeout(context.Background(), t.conf.Timeout)
	defer cencel()

	data, err := t.c.TorrentGetAllForHashes(ctx, hashes)
	if err != nil {
		return nil, err
	}

	var torrents []*Torrent
	for _, t := range data {
		if t.HashString == nil || t.DownloadDir == nil || t.Status == nil {
			continue
		}
		torrents = append(torrents, &Torrent{
			InfoHash:     *t.HashString,
			DownloadPath: *t.DownloadDir,
			Status:       TorrentStatus(*t.Status),
			Name:         *t.Name,
			IsFinished:   *t.IsFinished,
		})
	}

	return torrents, nil
}

func (t *Transmission) TorrentAdd(filename, dir string) error {
	ctx, cencel := context.WithTimeout(context.Background(), t.conf.Timeout)
	defer cencel()

	_, err := t.c.TorrentAdd(ctx, transmissionrpc.TorrentAddPayload{
		Filename:    &filename,
		DownloadDir: &dir,
	})
	return err
}
