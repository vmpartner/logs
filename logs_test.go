package logs

import (
	"os"
	"testing"
	"time"
)

func TestInitLogsToFile(t *testing.T) {
	err := InitLogsToFile("main.log")
	if err != nil {
		t.Fatal(err)
	}
	Info("Test")
	time.Sleep(100 * time.Microsecond)
	Close()
	f, err := os.Stat("main.log")
	if err != nil {
		t.Fatal(err)
	}
	if f.Size() == 0 {
		t.Fatal(err)
	}
}
