# Gistbox

Gistbox is a small, server-rendered web app for creating and viewing short text snippets that auto-expire. It is a learning project focused on backend wiring in Go rather than product polish.

## What it does now
- Sign up or log in, then create a gist with title, body, and expiry (1/7/365 days) and get a flash confirmation.
- List the 10 most recent, non-expired gists and view a single gist.
- Server-rendered HTML with cached templates, human-friendly dates, and a minimal CSS/JS theme.
- CSRF tokens plus secure session cookies (HTTPS-only) on all dynamic routes.
- Middleware covers panic recovery, structured request logging, security headers, and authentication gates.

## Why it exists
- Practice the full request/response stack without heavy frameworks.
- Explore middleware chaining, validation, sessions/flash messages, and DB-backed persistence.
- Keep scope small (gists + expiry) to emphasize backend fundamentals.

## Stack
- Go (standard library `net/http`, `html/template`, `log/slog`, `database/sql`)
- MySQL for persistence
- Sessions & flash messages via `github.com/alexedwards/scs` + `mysqlstore`
- Form decoding via `github.com/go-playground/form/v4`
- Middleware chaining via `github.com/justinas/alice`
- CSRF protection via `github.com/justinas/nosurf`

## How to run locally (quick path)
Prereqs: Go (module currently targets Go 1.25), a MySQL instance, and a way to generate local TLS certs (e.g. OpenSSL or mkcert).

1) Clone and install deps:
```bash
git clone <repo-url>
cd Go-Webapp/gistbox
go mod download
```

2) Create the MySQL schema (gists, users, sessions) using the statements in `docs/setup.md`:
```bash
mysql -u <user> -p   # then paste the SQL from docs/setup.md
```

3) Generate local TLS certs at `./tls/cert.pem` and `./tls/key.pem` (self-signed example):
```bash
mkdir -p tls
openssl req -x509 -newkey rsa:2048 -nodes -keyout tls/key.pem -out tls/cert.pem -days 365 -subj "/CN=localhost"
```

4) Run the server (point `-dsn` to your MySQL):
```bash
go run ./cmd/web \
  -addr=":4000" \
  -dsn="user:pass@tcp(localhost:3306)/gistbox?parseTime=true"
```
Then visit https://localhost:4000 and accept the self-signed cert. See `docs/setup.md` for schema and session-store tables.

## Review highlights
- Expiry is enforced in queries (expired gists are filtered).
- Template caching keeps rendering fast; helpers format dates.
- Sessions carry one-time flash messages after writes; auth state is session-backed.
- Middleware covers recovery, logging, security headers, CSRF tokens, and HTTPS-only cookies.

## Roadmap
- Short term: finish user flows (logout UX, account checks), add edit/delete, better validation, and pagination/search.
- Later: syntax highlighting, Docker/CI, and automated tests for handlers/middleware.

## Repository layout
- `cmd/web` — HTTP handlers, middleware, routing, template rendering, and program entrypoint.
- `internal/models` — MySQL-backed data access for gists and shared errors.
- `internal/validator` — Small validation helper used by form bindings.
- `ui` — HTML templates, CSS/JS, and static assets.

## Further details
See `docs/setup.md` for database schema, session store table, flags/defaults, and TLS notes.
