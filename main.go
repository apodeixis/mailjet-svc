package main

import (
	"os"

	"github.com/apodeixis/mailjet-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
