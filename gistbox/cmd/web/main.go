package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"database/sql"

	"snippetbox-webapp/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger 	*slog.Logger
	gists		*models.GistModel
}

func main() {
	// Define a command-line flag with the name 'addr'
	// with a default value of :4000
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:suarex@/gistbox?parseTime=true", "MYSQL data source name")
	
	// Important to use the flag.Parse() to parse the command-line flag
	flag.Parse()

	// Use the slog.New to init a new structured logger, which
	// writes to stdout stream and uses the default settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application {
		logger: logger,
		gists: 	&models.GistModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// Return a pool of connections to the database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
