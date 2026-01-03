# Setup

This document explains how to set up a local MySQL database and the network port used by the application.

Default values in the code

- Default DSN (in `cmd/web/main.go`): `web:suarex@/gistbox?parseTime=true`
- Default HTTP(S) address: `:4000` (listen address)

1) Install and start MySQL (example for Debian/Ubuntu)

```bash
sudo apt update
sudo apt install -y mysql-server
sudo systemctl enable --now mysql
```

2) Create the database and a local user

Run the MySQL client and execute:

```sql
CREATE DATABASE gistbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Example user (change password for your environment)
CREATE USER 'web'@'localhost' IDENTIFIED BY 'suarex';
GRANT ALL PRIVILEGES ON gistbox.* TO 'web'@'localhost';
FLUSH PRIVILEGES;
```

3) Update the DSN if needed

The application reads the DSN from the `-dsn` flag. Example:

```bash
go run ./cmd/web -dsn "user:password@/gistbox?parseTime=true"
```

4) Port & firewall

The server listens on the address set by the `-addr` flag (default `:4000`). If running on a system with a firewall, allow access to that port for testing:

```bash
# example using ufw
sudo ufw allow 4000/tcp
```

5) TLS certs

The repository expects certificate files at `./tls/cert.pem` and `./tls/key.pem`. For local testing you can create self-signed certs and save them in the `tls/` directory:

```bash
mkdir -p tls
openssl req -x509 -newkey rsa:4096 -nodes \
  -keyout tls/key.pem -out tls/cert.pem -days 365 -subj "/CN=localhost"
```

Notes

- Change the MySQL username/password in the DSN for real deployments; do not use sample passwords in production.
- If you want to run without TLS for local development, you must modify `cmd/web/main.go` to use `server.ListenAndServe()` and set `sessionManager.Cookie.Secure = false`.
