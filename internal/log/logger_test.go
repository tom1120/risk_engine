package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	//InitLogger("console", "")
	InitLogger("file", "./out")
	Errorf("this is error %s", "aa")
	Error("this is error!")
	Debug("this is debug!")
	Warn("this is warn!")
	Info("this is info!")
}
