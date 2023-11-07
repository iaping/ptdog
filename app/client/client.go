package client

type IClient interface {
	Type() Type
	String() string
	Torrents([]string) ([]*Torrent, error)
	TorrentAdd(filename, dir string) error
}

type Type uint8

const (
	TypeTransmission Type = iota
	TypeQbittorrent
)

func (t Type) Name() string {
	switch t {
	case TypeTransmission:
		return "Transmission"
	case TypeQbittorrent:
		return "qBittorrent"
	}
	return "unknown"
}

func (t Type) Supported() bool {
	switch t {
	case TypeTransmission, TypeQbittorrent:
		return true
	}
	return false
}
