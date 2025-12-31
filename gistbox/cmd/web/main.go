package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"snippetbox-webapp/internal/models"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger         *slog.Logger
	gists          *models.GistModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
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

	// Initialize a new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	// Initialize a New session Managager. Then configure it to use
	// the Mysql database as session store.
	// set a lifetime of 12 hours (sessions expire in 12 hours after being created)
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	// Make sure that the Secure attribute is set on our session cookies.
	// Which says that the cookie will only be sent by a user's web browser when
	// a HTTPS connection is being used (and will not be sent over an unsecure HTTP connection)
	sessionManager.Cookie.Secure = true
	// And add it to the application dependencies.
	app := &application{
		logger:         logger,
		gists:          &models.GistModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	// Initialize a tls.Config struct to build the non default TLS settings
	// we want the server to use. In this we're changing the curve preferences
	// value, so that only elliptic curves with assembly implementations are used.
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// Initialize a new http.Server struct. And set the Addr and Handler fields so
	// that the server uses the same network address and routes as before.
	// Also set the server's TLSConfig field to use the tlsConfig variable we just
	// created.
	server := &http.Server{
		Addr:    		*addr,
		Handler: 		app.routes(),

		// Create a *log.Logger from our structured logger handler, which write
		// log entries at Error level, and assign it to the ErrorLog field.
		ErrorLog:		slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:		tlsConfig,
		// Add idle, Read and Write timeouts to the server.
		IdleTimeout:	time.Minute,
		ReadTimeout:	5 * time.Second,
		WriteTimeout:	10 * time.Second,
	}

	logger.Info("starting server", "addr", *addr)

	// Call the ListenAndServe() method of our new http.Server struct to start the server
	// Use the ListenAndServeTLS() method to start the HTTPS server. We pass in the
	// paths to the TLS certificate and corresponding private key as the two params
	err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
