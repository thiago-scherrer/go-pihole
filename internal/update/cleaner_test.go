package update

import "testing"

func TestClean(t *testing.T) {
	mock := []byte("exemple.com")
	_, err := Clean(mock)

	if err != nil {
		t.Error("Clean() error, got: ", err)
	}
}

func TestRmComent(t *testing.T) {
	mock := []byte("exemple.com")
	_, err := rmComent(mock)

	if err != nil {
		t.Error("Clean() error, got: ", err)
	}
}

func TestRmNonDomains(t *testing.T) {
	mock := "exemple.com"
	_, err := rmNonDomains(mock)

	if err != nil {
		t.Error("Clean() error, got: ", err)
	}
}
