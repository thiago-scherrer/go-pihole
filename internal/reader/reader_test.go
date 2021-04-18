package reader

import (
	"os"
	"testing"
)

func TestGetlist(t *testing.T) {
	os.Setenv("LIST", "mock.txt")

	var want []string
	want = append(want, "example.com")

	got, err := Getlist()

	if err != nil {
		t.Error("Got error on Getlist(): ", err)
	}

	if got[0] != want[0] {
		t.Errorf("Getlist(): Want %v but got %v", want, got)
	}
}

func TestGetlistEnv(t *testing.T) {
	os.Setenv("LIST", "")
	_, err := Getlist()

	if err == nil {
		t.Error("Getlist(): Need error but working?")
	}
}

func TestGetlistFile(t *testing.T) {
	os.Setenv("LIST", "mockfake.txt")
	_, err := Getlist()

	if err == nil {
		t.Error("Getlist(): Need error but working?")
	}
}
