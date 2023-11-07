package client

import (
	"fmt"

	"github.com/autobrr/go-qbittorrent"
)

type Qbittorrent struct {
	conf Config
	c    *qbittorrent.Client
}

func NewQbittorrent(conf Config) (*Qbittorrent, error) {
	scheme := "http"
	if conf.Https {
		scheme = "https"
	}

	host := fmt.Sprintf("%s://%s:%d", scheme, conf.Host, conf.Port)

	c := qbittorrent.NewClient(qbittorrent.Config{
		Host:          host,
		Username:      conf.Username,
		Password:      conf.Password,
		TLSSkipVerify: !conf.Https,
	})

	return &Qbittorrent{
		conf: conf,
		c:    c,
	}, nil
}

func (q *Qbittorrent) Type() Type {
	return TypeQbittorrent
}

func (q *Qbittorrent) String() string {
	return fmt.Sprintf("%s: %s:%d", q.Type().Name(), q.conf.Host, q.conf.Port)
}

func (q *Qbittorrent) Torrents(hashes []string) ([]*Torrent, error) {
	if err := q.c.Login(); err != nil {
		return nil, err
	}

	data, err := q.c.GetTorrents(qbittorrent.TorrentFilterOptions{
		Hashes: hashes,
	})
	if err != nil {
		return nil, err
	}

	var torrents []*Torrent
	for _, t := range data {
		torrents = append(torrents, &Torrent{
			Name:         t.Name,
			InfoHash:     t.Hash,
			DownloadPath: t.SavePath,
			Status:       q.transform(t.State),
			IsFinished:   t.CompletionOn > 0,
		})
	}

	return torrents, nil
}

func (q *Qbittorrent) TorrentAdd(url, dir string) error {
	if err := q.c.Login(); err != nil {
		return err
	}

	skipChecking := "false"
	if q.conf.SkipChecking {
		skipChecking = "true"
	}

	return q.c.AddTorrentFromUrl(url, map[string]string{
		"savepath":      dir,
		"skip_checking": skipChecking,
	})
}

func (q *Qbittorrent) transform(status qbittorrent.TorrentState) TorrentStatus {
	switch status {
	//"uploading", "stalledUP", "forcedUP"
	case qbittorrent.TorrentStateUploading, qbittorrent.TorrentStateStalledUp, qbittorrent.TorrentStateForcedUp:
		return TorrentStatusSeed
	}
	return TorrentStatus(-1)
}
