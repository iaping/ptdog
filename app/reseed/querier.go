package reseed

import (
	"ptdog/app/client"
	"sync"

	"github.com/rs/zerolog/log"
)

type Batch struct {
	client   client.IClient
	torrents map[string]*client.Torrent
	pieces   []string
}

func (b *Batch) init() {
	for hash := range b.torrents {
		b.pieces = append(b.pieces, hash)
	}
}

var querier = &Querier{
	queue: make(chan *Batch, 8),
}

type Querier struct {
	websites []*Website
	queue    chan *Batch
	wg       sync.WaitGroup
}

func (q *Querier) Push(batch *Batch) {
	q.queue <- batch
}

func (q *Querier) Websites(websites []*Website) *Querier {
	q.websites = websites
	return q
}

func (q *Querier) Run() {
	go func() {
		for batch := range q.queue {
			q.handler(batch)
		}
	}()
}

func (q *Querier) handler(batch *Batch) {
	if len(q.websites) == 0 {
		return
	}

	batch.init()

	q.wg.Add(len(q.websites))
	for _, website := range q.websites {
		go q.batch(batch, website)
	}
	q.wg.Wait()
}

func (q *Querier) batch(batch *Batch, website *Website) {
	defer q.wg.Done()

	length := len(batch.pieces)

	for i := 0; i < length; i++ {
		s := i * website.limit
		e := s + website.limit
		if e >= length {
			q.query(batch, website, batch.pieces[s:])
			break
		}
		q.query(batch, website, batch.pieces[s:e])
	}
}

func (q *Querier) query(batch *Batch, website *Website, hashes []string) {
	log.Info().Str("站点", website.String()).Int("Hashes", len(hashes)).Msg("开始查询Hash数据")

	data, err := website.Query(hashes)
	if err != nil {
		log.Err(err).Str("站点", website.String()).Msg("查询失败")
		return
	}

	for hash, id := range data {
		if torrent, ok := batch.torrents[hash]; ok {
			seeder.Push(Seed{
				client:  batch.client,
				website: website,
				id:      id,
				torrent: torrent,
			})
		}
	}
}
