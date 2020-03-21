package tg_logger

import (
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/config"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/tags"
	"time"
)

type Log struct {
	Username      string
	Exercise      string
	Repeats       string
	Time_of_start time.Time
	Time_of_end   time.Time
}

func NewLog(t *tags.Tags, c *config.Config) (l *Log) {
	l = &Log{
		Username:      t.User,
		Exercise:      t.Exercise,
		Repeats:       t.Repeats,
		Time_of_start: c.TIME_OF_START,
		Time_of_end:   time.Now(),
	}
	return
}
