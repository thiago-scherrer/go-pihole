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
				t.Errorf("rmComent() = %v, want %v", got, tt.want)
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
		{"Comment start", "repo example.com", ""},
		{"Comment middle", "0.0.0.0 example.com", "example.com"},
		{"Comment middle", "127.0.0.1 example.com", "example.com"},
		{"Comment middle", "localhost example.com", "example.com"},
		{"Comment middle", "\n\n example.com", "example.com"},
		{"Comment middle", "\t example.com", "example.com"},
		{"Comment middle", "@ example.com", "example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := rmNonDomains(tt.mock); got != tt.want {
				t.Errorf("rmComent() = %v, want %v", got, tt.want)
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
