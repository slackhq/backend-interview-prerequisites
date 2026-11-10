# Rust environment check

Install Rust 1.89+ using [rustup](https://rustup.rs), plus a C compiler
for the bundled SQLite library (Xcode Command Line Tools on macOS,
`build-essential` on Ubuntu, or Visual Studio C++ Build Tools on Windows).

From this directory, run:

```sh
cargo run --locked
```

Cargo restores rusqlite with bundled SQLite and serde_json.
The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output, `rustc --version`, and your OS version to your
recruiting coordinator.
