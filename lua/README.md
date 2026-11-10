# Lua environment check

Requires Lua 5.4 or 5.5, LuaRocks, and a C compiler/Make for native modules.
On macOS: `brew install lua luarocks`. On Ubuntu:
`sudo apt-get install lua5.4 liblua5.4-dev luarocks build-essential`.
Ensure LuaRocks targets the same Lua version as the `lua` command.

From this directory, run:

```sh
sh setup.sh
```

The script installs LuaSocket, lua-cjson, and lsqlite3complete into `.luarocks/`,
then verifies SQLite persistence and bidirectional JSON over TCP using a
temporary database and an automatically assigned loopback port.
If it fails, send the output, `lua -v`, `luarocks --version`, and your OS
version to your recruiting coordinator. Windows users can run this check in WSL.
