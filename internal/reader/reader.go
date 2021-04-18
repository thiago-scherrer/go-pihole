package reader

import (
	"bufio"
	"errors"
	"os"
)

var (
	ErrNeedEnv = errors.New("need LIST env to check new blocklist")
)

// Getlist return blocklist to use on sync
func Getlist() ([]string, error) {
	var url []string

	list := os.Getenv("LIST")
	if len(list) == 0 {
		return nil, ErrNeedEnv
	}

	file, err := os.Open(list)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url = append(url, string(scanner.Text()))

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return url, nil
}
