package main

import (
	"github.com/Qs-F/go-play"
	"log"
)

func main() {
	s, err := play.Compose()
	if err != nil {
		log.Fatal(err.Error)
		return
	}
	log.Println(s)
}
