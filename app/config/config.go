package config

import (
	"encoding/json"
	"os"
)

var Conf = &Config{
	Http: Http{
		Addr: "127.0.0.1:1688",
	},
	System: System{
		Sleep: 10,
	},
}

type Config struct {
	Http     Http      `json:"http"`
	System   System    `json:"system"`
	Clients  []Client  `json:"clients"`
	Websites []Website `json:"websites"`
}

func Exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func Load(path string) error {
	if !Exist(path) {
		return nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, Conf)
}
