package reader

import (
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Getlist return blocklist to use on sync
func Getlist() (string, error) {
	list := os.Getenv("LIST")
	if len(list) <= 0 {
		log.Fatal("Need LIST env to check new blocklist")
	}

	dat, err := ioutil.ReadFile(list)
	check(err)

	return string(dat), nil
}
