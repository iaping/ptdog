module ptdog

go 1.21.2

require (
	github.com/anacrolix/torrent v1.53.1
	github.com/autobrr/go-qbittorrent v1.5.0
	github.com/gookit/cache v0.4.0
	github.com/hekmon/transmissionrpc/v2 v2.0.1
	github.com/rs/zerolog v1.31.0
)

require (
	github.com/anacrolix/missinggo v1.3.0 // indirect
	github.com/anacrolix/missinggo/v2 v2.7.2 // indirect
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/bradfitz/iter v0.0.0-20191230175014-e8f45d346db8 // indirect
	github.com/gookit/gsr v0.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hekmon/cunits/v2 v2.1.0 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)

replace github.com/autobrr/go-qbittorrent v1.5.0 => github.com/iaping/go-qbittorrent v0.0.0-20231106074650-9991b94e4419

replace github.com/gookit/cache v0.4.0 => github.com/iaping/cache v0.0.0-20231106113618-edded85d0f13
