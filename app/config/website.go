package config

type Website struct {
	Enable   bool   `json:"enable"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Api      string `json:"api"`
	Passkey  string `json:"passkey"`
	Download string `json:"download"`
	Limit    int    `json:"limit"`
}
