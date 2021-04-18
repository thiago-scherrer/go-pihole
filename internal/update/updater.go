package update

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thiago-scherrer/go-pihole/internal/reader"
)

func Run() error {
	list, err := reader.Getlist()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, k := range list {
		getDomains(k)
	}
	return nil
}

func getDomains(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[erro] http.get failed: ", err)
		return nil, err
	}

	Clean(body)

	return nil, nil
}
