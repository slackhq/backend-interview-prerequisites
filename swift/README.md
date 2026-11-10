# Swift environment check

Requires Swift 5.9+ and SQLite 3 development headers.
On macOS, install a compatible Xcode or Swift toolchain and
`brew install sqlite`. On Ubuntu, install Swift using
the [Swift installation guide](https://www.swift.org/install/)
and `sudo apt-get install libsqlite3-dev pkg-config`.

From this directory, run:

```sh
swift run
```

The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output, `swift --version`, and your OS version to your
recruiting coordinator. Windows users should run the Swift exercise and this
check in WSL.
