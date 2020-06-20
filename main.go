package main

import (
	"log"
	"os"

	"github.com/create-go-app/cli/pkg/cgapp"
)

func main() {
	// Start new CLI app
	cli, err := cgapp.New()
	if err != nil {
		log.Fatal(err)
	}

	// Run new CLI
	if err = cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
