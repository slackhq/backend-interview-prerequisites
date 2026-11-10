# C++ environment check

Requires CMake 3.20+, a C++20 compiler, SQLite 3 development files, and Git.
CMake downloads nlohmann/json 3.11.3 on first configuration.

On macOS: `brew install cmake sqlite`.
On Ubuntu: `sudo apt-get install build-essential cmake libsqlite3-dev git`.

From this directory, run:

```sh
cmake -S . -B build
cmake --build build
ctest --test-dir build --output-on-failure
```

The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output and OS/compiler versions to your recruiting
coordinator. Windows users should run the C++ exercise and this check in WSL.
