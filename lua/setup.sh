#!/bin/sh
set -eu
cd "$(dirname "$0")"
luarocks --tree .luarocks install lua-cjson 2.1.0.10-1
luarocks --tree .luarocks install luasocket 3.1.0-1
luarocks --tree .luarocks install lsqlite3complete 0.9.7-1
export LUA_PATH="$(luarocks --tree .luarocks path --lr-path)"
export LUA_CPATH="$(luarocks --tree .luarocks path --lr-cpath)"
exec lua setup.lua
