# C environment check

Requires a C11 compiler, Make, pkg-config, SQLite 3 development files, and
json-c development files.

On macOS: `brew install pkg-config sqlite json-c`.
On Ubuntu: `sudo apt-get install build-essential pkg-config libsqlite3-dev libjson-c-dev`.

From this directory, run:

```sh
make setup
```

The check verifies SQLite writes survive reopening the database and exchanges
JSON in both directions over a local TCP connection. It uses a temporary
database and an automatically assigned loopback port.

If it fails, send the output and your OS/compiler versions to your recruiting
coordinator. Windows users should run the C exercise and this check in WSL.
