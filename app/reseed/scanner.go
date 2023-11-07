package reseed

import (
	"os"
	"path"
	"ptdog/app/client"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Scanner struct {
	client client.IClient
	dir    string
	sleep  time.Duration
}

func NewScanner(client client.IClient, dir string, sleep time.Duration) *Scanner {
	return &Scanner{
		client: client,
		dir:    dir,
		sleep:  sleep,
	}
}

func (s *Scanner) Run() {
	go func() {
		ticker := time.NewTicker(s.sleep)
		for {
			s.scan()
			<-ticker.C
		}
	}()
}

func (s *Scanner) scan() {
	torrents, err := s.torrents()
	if err != nil {
		s.log(log.Err(err)).Msg("扫描错误")
		return
	}

	s.log(log.Info()).Int("匹配种子", len(torrents)).Msg("扫描完成")

	if len(torrents) > 0 {
		querier.Push(&Batch{
			client:   s.client,
			torrents: torrents,
		})
	}
}

func (s *Scanner) torrents() (map[string]*client.Torrent, error) {
	hashes, err := s.load()
	if err != nil {
		return nil, err
	}

	var queries []string
	for hash := range hashes {
		queries = append(queries, hash)
	}

	data, err := s.client.Torrents(queries)
	if err != nil {
		return nil, err
	}

	var torrents = make(map[string]*client.Torrent)
	for _, t := range data {
		if !t.IsFinished {
			continue
		}

		// if !t.Seeding() {
		// 	continue
		// }

		if piecesHash, ok := hashes[t.InfoHash]; ok {
			t.PiecesHash = piecesHash
			torrents[t.PiecesHash] = t
		}
	}

	return torrents, nil
}

func (s *Scanner) load() (map[string]string, error) {
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		return nil, err
	}

	var hashes = make(map[string]string)
	for _, entry := range entries {
		path := path.Join(s.dir, entry.Name())

		meta, err := metainfo.LoadFromFile(path)
		if err != nil {
			continue
		}

		info, err := meta.UnmarshalInfo()
		if err != nil {
			continue
		}

		hash := meta.HashInfoBytes().HexString()
		piecesHash := metainfo.HashBytes(info.Pieces).HexString()
		hashes[hash] = piecesHash
	}

	return hashes, nil
}

func (s *Scanner) log(event *zerolog.Event) *zerolog.Event {
	event.Str("等待", s.sleep.String())
	event.Str("目录", s.dir)
	event.Str("下载器", s.client.String())
	return event
}
