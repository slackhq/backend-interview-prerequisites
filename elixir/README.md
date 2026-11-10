# Elixir environment check

Requires Elixir 1.17+, a compatible Erlang/OTP release, and a C compiler/Make
for SQLite's native binding if a precompiled binary is unavailable.
On macOS: `brew install elixir`. On Ubuntu, install Elixir/Erlang using
the [Elixir installation guide](https://elixir-lang.org/install.html)
and install `build-essential`.

From this directory, run:

```sh
mix local.hex --force
mix local.rebar --force
mix deps.get
mix run setup.exs
```

The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output, `elixir --version`, and your OS version to your
recruiting coordinator.
