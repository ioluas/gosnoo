package main

import (
	"fmt"
	"os"

	"github.com/ioluas/gosnoo/internal/app"
	"github.com/ioluas/gosnoo/internal/reddit"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("reddit-tui %s (commit: %s, built: %s)\n", version, commit, date)
		return
	}

	client, err := reddit.NewReadonlyClient()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error creating reddit client: %v\n", err)
		os.Exit(1)
	}

	svc := reddit.NewService(client)

	if err := app.Run(svc); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
