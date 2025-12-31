# Gistbox

Gistbox is a small, server-rendered web app for creating and viewing short text snippets that auto-expire. It is a learning project focused on backend wiring in Go rather than product polish.

## What it does now
- Create a gist with title, body, and expiry (1/7/365 days) and get a flash confirmation.
- List the 10 most recent, non-expired gists and view a single gist.
- Server-rendered HTML with cached templates, human-friendly dates, and a minimal CSS/JS theme.
- Middleware covers panic recovery, structured request logging, and security headers.

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

## How to run locally (quick path)
Prereqs: Go (module currently targets Go 1.25) and a MySQL instance.

1) Clone and install deps:
```bash
git clone <repo-url>
cd Go-Webapp/gistbox
go mod download
```

2) Run the server (point `-dsn` to your MySQL):
```bash
go run ./cmd/web \
  -addr=":4000" \
  -dsn="user:pass@tcp(localhost:3306)/gistbox?parseTime=true"
```
Then visit http://localhost:4000. See `docs/setup.md` for schema and session-store tables.

## Review highlights
- Expiry is enforced in queries (expired gists are filtered).
- Template caching keeps rendering fast; helpers format dates.
- Sessions carry one-time flash messages after writes.
- Middleware covers recovery, logging, and common security headers.

## Roadmap
- Short term: add auth, edit/delete, better validation, and pagination/search.
- Later: syntax highlighting, Docker/CI, and automated tests.

## Repository layout
- `cmd/web` — HTTP handlers, middleware, routing, template rendering, and program entrypoint.
- `internal/models` — MySQL-backed data access for gists and shared errors.
- `internal/validator` — Small validation helper used by form bindings.
- `ui` — HTML templates, CSS/JS, and static assets.

## Further details
See `docs/setup.md` for database schema, session store table, flags/defaults, and TLS notes.
