# Gistbox (Go Webapp)
Gistbox is a minimal web application built as a first backend-engineering project to learn how backend development works in Go and to prepare for backend job opportunities. The primary goal of this project is to demonstrate practical backend skills: routing, templates, request context, embedding static assets, simple persistence, and basic auth patterns.

Signed up and logged in Users can create short code snippets (gists) that other users can view. There are no interactions on gists (likes, comments) in this iteration— the focus is on backend fundamentals and clear, maintainable server-side code.

Let's go by Alex Edwards was the resource used in learning and building of this app.

## Requirements

- Go 1.16+ (required for `embed`; newer Go versions recommended)
- Git

## Project layout

- `cmd/web/` - application entrypoint and HTTP handlers
- `internal/models/` - application models
- `ui/html/` - HTML templates (embedded)
- `ui/static/` - static assets (embedded)
- `tls/` - TLS cert/key for HTTPS (local certs)

## Build & Run

Important: the current code starts an HTTPS server unconditionally using `server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")`. That means the process expects a certificate and key at `./tls/cert.pem` and `./tls/key.pem` and will fail to start if they are missing.

Run with HTTPS (recommended / current default):

```bash
mkdir -p tls
openssl req -x509 -newkey rsa:4096 -nodes -keyout tls/key.pem -out tls/cert.pem -days 365 -subj "/CN=localhost"
go run ./cmd/web
```

Build and run the binary (HTTPS):

```bash
go build -o bin/web ./cmd/web
./bin/web
```

Core features

- Create and view short code snippets (gists). No likes/comments/profiles in this version — focus is backend behavior.
- Server-side templates, embedded static assets, simple persistence, authentication and authorization, and session-based auth.

Requirements

- Go 1.16+ (for `embed` support)
- A MySQL instance for development (see `cmd/web/main.go` DSN)

Quick start (HTTPS — current behavior)

The code currently starts an HTTPS server using `server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")`. Provide a self-signed certificate pair in a `tls/` directory before running:

```bash
mkdir -p tls
openssl req -x509 -newkey rsa:4096 -nodes -keyout tls/key.pem -out tls/cert.pem -days 365 -subj "/CN=localhost"
go run ./cmd/web
```

Build and run the binary:

```bash
go build -o bin/web ./cmd/web
./bin/web
```

If you prefer to run without TLS for local testing, edit `cmd/web/main.go` to use `server.ListenAndServe()` and set `sessionManager.Cookie.Secure = false` (cookies require `Secure=true` for HTTPS).

Embedding & templates

Templates and static files are embedded via the `ui` package (`ui/efs.go`). The template cache uses `fs.Glob` and `template.ParseFS` to load templates from the embedded filesystem at startup. Edit files under `ui/html/` and rebuild.

Context & auth

The project uses a typed context key in `cmd/web/context.go` to store authentication state injected by middleware and consumed by handlers.

Database setup

Reference the docs/setup.md file, to see how to set up the database schema. 

Testing

- Add package tests with `go test ./...`.
- Prefer table-driven tests and subtests (`t.Run`) to cover edge cases.

Future UI updates

- Incrementally add UI features (likes, comments, profiles) once backend APIs and tests are stable.
- Improve styling under `ui/static/css` and template partials.
