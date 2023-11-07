package client

import (
	"time"
)

type Config struct {
	Https        bool
	Host         string
	Username     string
	Password     string
	Port         uint16
	Timeout      time.Duration
	SkipChecking bool
}
