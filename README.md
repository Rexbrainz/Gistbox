# Gistbox

Gistbox is a small, server-rendered web app for creating and viewing short text snippets that auto-expire. The project exists as a learning playground for backend engineering: wiring HTTP handlers, middleware, templates, validation, and a relational database together in Go.

## Why it exists
- Practice building a full request/response stack without heavy frameworks.
- Learn how to layer middleware for logging, recovery, and security headers.
- Experiment with session-backed flash messages and form validation.
- Keep the surface area small (gists + expiry) while focusing on backend fundamentals.

## What works today
- Create a gist with a title, body, and expiry (1/7/365 days) and get a confirmation flash.
- List the 10 most recent, non-expired gists and drill into a single gist view.
- Server-rendered HTML with cached templates, human-friendly dates, and a minimal CSS/JS theme.
- Middleware for panic recovery, structured request logging, and sensible security headers.

## Stack
- Go (standard library `net/http`, `html/template`, `log/slog`, `database/sql`)
- MySQL for persistence
- Sessions & flash messages via `github.com/alexedwards/scs` + `mysqlstore`
- Form decoding via `github.com/go-playground/form/v4`
- Middleware chaining via `github.com/justinas/alice`

## Project status and next steps
- Current: anonymous gist creation + read-only views; no auth; minimal validation; no deployment scripts.
- Planned ideas: user accounts, edit/delete, syntax highlighting, pagination/search, Dockerfile + CI, and automated tests.

## Getting started
Prereqs: Go (module currently targets Go 1.25), MySQL running locally, and a database user with rights to create tables.

1) Clone and install deps
```bash
git clone <repo-url>
cd Go-Webapp/gistbox
go mod download
```

2) Prepare the database (example schema)
```sql
CREATE DATABASE gistbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE gistbox;

CREATE TABLE gists (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created DATETIME NOT NULL,
  expires DATETIME NOT NULL,
  KEY expires_idx (expires)
);

-- Session store for SCS (see github.com/alexedwards/scs/mysqlstore for details)
CREATE TABLE sessions (
  token CHAR(43) PRIMARY KEY,
  data BLOB NOT NULL,
  expiry TIMESTAMP(6) NOT NULL,
  KEY sessions_expiry_idx (expiry)
);
```

3) Run the server
```bash
# Override the defaults as needed
go run ./cmd/web \
  -addr=":4000" \
  -dsn="user:pass@tcp(localhost:3306)/gistbox?parseTime=true"
```
Then visit http://localhost:4000.

## Repository layout
- `cmd/web` — HTTP handlers, middleware, routing, template rendering, and program entrypoint.
- `internal/models` — MySQL-backed data access for gists and shared errors.
- `internal/validator` — Small validation helper used by form bindings.
- `ui` — HTML templates, CSS/JS, and static assets.

## Notes for reviewers
- The app is intentionally simple to spotlight backend wiring rather than product polish.
- Gists expire in the database query layer; expired entries are filtered from reads.
- No authentication yet—everyone can create and view public gists.
