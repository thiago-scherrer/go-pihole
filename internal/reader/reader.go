package reader

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
)

var (
	errNeedEnv = errors.New("[error]: need LIST env to check new blocklist")
)

// Getlist return blocklist to use on sync
func Getlist() ([]string, error) {
	var url []string

	list := os.Getenv("LIST")
	if len(list) == 0 {
		log.Println(errNeedEnv)
		return nil, errNeedEnv
	}

	file, err := os.Open(filepath.Clean(list))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url = append(url, string(scanner.Text()))

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return url, nil
}
