# Gistbox setup details

## Requirements
- Go 1.25+
- MySQL with a user that can create tables
- OpenSSL or mkcert for generating local TLS certificates

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

CREATE TABLE users (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created DATETIME NOT NULL,
  CONSTRAINT users_uc_email UNIQUE (email)
);

-- Session store for github.com/alexedwards/scs/mysqlstore
CREATE TABLE sessions (
  token CHAR(43) PRIMARY KEY,
  data BLOB NOT NULL,
  expiry TIMESTAMP(6) NOT NULL,
  KEY sessions_expiry_idx (expiry)
);
```

## TLS for local dev
The server calls `ListenAndServeTLS` and session/CSRF cookies are marked `Secure`, so HTTPS is required even locally. Generate self-signed certs at `./tls/cert.pem` and `./tls/key.pem` (relative to the `gistbox` directory):
```bash
mkdir -p tls
openssl req -x509 -newkey rsa:2048 -nodes -keyout tls/key.pem -out tls/cert.pem -days 365 -subj "/CN=localhost"
```
The `tls/` directory is gitignored and safe for local-only keys.

## Running locally
```bash
go run ./cmd/web \
  -addr=":4000" \
  -dsn="user:pass@tcp(localhost:3306)/gistbox?parseTime=true"
```
Then visit https://localhost:4000 and accept the self-signed cert.

## Flags and defaults
- `-addr`: defaults to `:4000`.
- `-dsn`: defaults to `web:suarex@/gistbox?parseTime=true` (local MySQL user `web`, database `gistbox`).
- TLS files: `./tls/cert.pem` and `./tls/key.pem` relative to the `gistbox` directory are used by default.
- Templates live in `ui/html`; static assets in `ui/static`.
- Module path is `snippetbox-webapp` (inherited from the tutorial source); rename if you fork under a different module path.

## TLS and ignores
- `.gitignore` excludes a local `tls/` directory, in case you generate self-signed certs for experiments (see above).
