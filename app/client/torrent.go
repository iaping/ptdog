package client

type TorrentStatus int

const (
	// TorrentStatusStopped represents a stopped torrent
	TorrentStatusStopped TorrentStatus = 0
	// TorrentStatusCheckWait represents a torrent queued for files checking
	TorrentStatusCheckWait TorrentStatus = 1
	// TorrentStatusCheck represents a torrent which files are currently checked
	TorrentStatusCheck TorrentStatus = 2
	// TorrentStatusDownloadWait represents a torrent queue to download
	TorrentStatusDownloadWait TorrentStatus = 3
	// TorrentStatusDownload represents a torrent currently downloading
	TorrentStatusDownload TorrentStatus = 4
	// TorrentStatusSeedWait represents a torrent queued to seed
	TorrentStatusSeedWait TorrentStatus = 5
	// TorrentStatusSeed represents a torrent currently seeding
	TorrentStatusSeed TorrentStatus = 6
	// TorrentStatusIsolated represents a torrent which can't find peers
	TorrentStatusIsolated TorrentStatus = 7
)

type Torrent struct {
	Name         string
	InfoHash     string
	PiecesHash   string
	DownloadPath string
	Status       TorrentStatus
	IsFinished   bool
}

func (t Torrent) Seeding() bool {
	return t.Status == TorrentStatusSeed
}
