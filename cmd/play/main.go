package main

import (
	"github.com/Qs-F/go-play"
)

func main() {
	s, err := play.Compose()
	if err != nil {
		play.OutputErr(err.Error())
		return
	}
	play.Output(s)
}
