package reader

import (
	"os"
	"testing"
)

func TestCheck(t *testing.T) {
	os.Setenv("LIST", "mock.txt")
	_, err := Getlist()

	check(err)

}

func TestGetlist(t *testing.T) {
	os.Setenv("LIST", "mock.txt")

	want := "example.com\n"
	got, err := Getlist()

	if err != nil {
		t.Error("Got error on Getlist(): ", err)
	}

	if got != "example.com\n" {
		t.Errorf("Erro on Getlist(). Want %v but got %v", want, got)
	}
}
