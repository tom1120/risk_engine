package util

import (
	"time"
)

func TimeSince(from time.Time) int64 {
	return int64(time.Since(from)) / 1e6
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
