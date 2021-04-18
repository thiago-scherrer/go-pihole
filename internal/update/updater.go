package update

import (
	"fmt"
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

func getDomains(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	content := make([]byte, 1014)
	c, err := resp.Body.Read(content)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	a := string(content[:c])

	fmt.Println(a)

	return nil, nil
}
