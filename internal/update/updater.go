package update

import (
	"fmt"
	"log"

	"github.com/thiago-scherrer/go-pihole/internal/reader"
)

func Run() {
	list, err := reader.Getlist()
	if err != nil {
		log.Fatal("Getlist() problem, got: ", err)
	}

	for _, k := range list {
		fmt.Println(k)
	}
}
