package logctx

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_LOG_LEVEL LogLevel = LogLevel(logrus.InfoLevel)
)

type LogLevel uint32

func (l *LogLevel) MarshalJSON() (b []byte, err error) {
	level := logrus.Level(*l)
	return level.MarshalText()
}

func (l *LogLevel) UnmarshalJSON(b []byte) (err error) {
	var lstr string
	if err = json.Unmarshal(b, &lstr); err != nil {
		return
	}
	var level logrus.Level
	if level, err = logrus.ParseLevel(lstr); err == nil {
		*l = LogLevel(level)
	}
	return
}
