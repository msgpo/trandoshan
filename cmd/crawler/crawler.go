package main

import (
	"github.com/creekorful/trandoshan/internal/crawler"
	"os"
)

func main() {
	app := crawler.GetApp()
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
