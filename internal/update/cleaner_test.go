package update

import (
	"os"
	"testing"
)

func TestClean(t *testing.T) {
	mock := []byte("exemple.com")
	_, err := Clean(mock)

	if err != nil {
		t.Error("Clean() error, got: ", err)
	}
}

func TestRmComent(t *testing.T) {

	tests := []struct {
		name string
		mock string
		want string
	}{
		{"Comment start", "#example.com", ""},
		{"Comment middle", "example.com#This is commet", "example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := rmComent([]byte(tt.mock)); got != tt.want {
				t.Errorf("%v: rmComent() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestRmNonDomains(t *testing.T) {
	tests := []struct {
		name string
		mock string
		want string
	}{
		{"Test 1", "repo example.com", ""},
		{"Test 2", "0.0.0.0 example.com", "example.com"},
		{"Test 3", "127.0.0.1 example.com", "example.com"},
		{"Test 4", "localhost example.com", "example.com"},
		{"Test 5", "example.org<br>exemple.com", "example.org\nexemple.com"},
		{"Test 6", "example.org<BR>exemple.com", "example.org\nexemple.com"},
		{"Test 7", "@ example.com", "example.com"},
		{"Test 8", "::1example.com", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := rmNonDomains(tt.mock); got != tt.want {
				t.Errorf("%v: rmNonDomains() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestWriteDom(t *testing.T) {
	mock := "exemple.com"
	os.Setenv("OUTPUT", "/tmp/.go-pihole.txt")

	_, err := writeDom(mock)

	if err != nil {
		t.Error("writeDom() error, got: ", err)
	}
}
