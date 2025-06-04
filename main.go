package main

import (
	"github.com/katallaxie/service-lens/cmd"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		panic(err)
	}
}
