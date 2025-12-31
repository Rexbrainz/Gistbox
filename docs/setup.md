# Gistbox setup details

## Requirements
- Go 1.25+
- MySQL with a user that can create tables

## Database schema
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

-- Session store for github.com/alexedwards/scs/mysqlstore
CREATE TABLE sessions (
  token CHAR(43) PRIMARY KEY,
  data BLOB NOT NULL,
  expiry TIMESTAMP(6) NOT NULL,
  KEY sessions_expiry_idx (expiry)
);
```

## Running locally
```bash
go run ./cmd/web \
  -addr=":4000" \
  -dsn="user:pass@tcp(localhost:3306)/gistbox?parseTime=true"
```

## Flags and defaults
- `-addr`: defaults to `:4000`.
- `-dsn`: defaults to `web:suarex@/gistbox?parseTime=true` (local MySQL user `web`, database `gistbox`).
- Templates live in `ui/html`; static assets in `ui/static`.
- Module path is `snippetbox-webapp` (inherited from the tutorial source); rename if you fork under a different module path.

## TLS and ignores
- `.gitignore` excludes a local `tls/` directory, in case you generate self-signed certs for experiments.
