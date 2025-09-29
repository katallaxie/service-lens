package main

import (
	"log"
	"os"

	"github.com/katallaxie/service-lens/cmd"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	if err := cmd.Init(); err != nil {
		log.Fatal(err)
	}
}
