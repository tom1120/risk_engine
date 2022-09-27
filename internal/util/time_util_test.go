package util

import (
	"testing"
	"time"
)

func TestTimeSince(t *testing.T) {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	t.Log(TimeSince(start))
}
