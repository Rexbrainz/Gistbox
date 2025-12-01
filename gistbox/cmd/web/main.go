package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Define a command-line flag with the name 'addr'
	// with a default value of :4000
	addr := flag.String("addr", ":4000", "HTTP network address")
	
	// Important to use the flag.Parse() to parse the command-line flag
	flag.Parse()

	// Use the slog.New to init a new structured logger, which
	// writes to stdout stream and uses the default settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application {
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
