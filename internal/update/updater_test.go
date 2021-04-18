package update

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	os.Setenv("LIST", "mock.txt")
	got := Run()

	if got != nil {
		t.Error("Run() return error: ", got)
	}
}
