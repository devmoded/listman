package main

import (
	"log"

	"github.com/devmoded/listman/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
