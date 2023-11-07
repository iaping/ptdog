package config

import "time"

type System struct {
	Sleep int64 `json:"sleep"`
}

func (c System) SleepDuration() time.Duration {
	if c.Sleep <= 0 {
		c.Sleep = 60
	}
	return time.Duration(c.Sleep) * time.Minute
}
