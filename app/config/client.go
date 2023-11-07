package config

import (
	"ptdog/app/client"
)

type Client struct {
	Enable       bool        `json:"enable"`
	Https        bool        `json:"https"`
	Type         client.Type `json:"type"`
	Host         string      `json:"host"`
	Username     string      `json:"username"`
	Password     string      `json:"password"`
	Port         uint16      `json:"port"`
	Timeout      int64       `json:"timeout"`
	Dir          string      `json:"dir"`
	SkipChecking bool        `json:"skip_checking"`
}
