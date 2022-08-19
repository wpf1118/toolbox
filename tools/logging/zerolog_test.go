package logging

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	l := GetLogger()
	l.Error().Msgf("%10d xxxxx%sxxxx", 10110000001, "sbsb")

	DebugF("aaa")
	DebugF("bbb")
}
