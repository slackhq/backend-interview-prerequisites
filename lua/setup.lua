local socket = require("socket")
local json = require("cjson")
local sqlite = require("lsqlite3complete")

local path = os.tmpname()
local db = assert(sqlite.open(path))
assert(db:exec("CREATE TABLE things (name TEXT); INSERT INTO things VALUES ('Hello, world');") == sqlite.OK)
assert(db:close() == sqlite.OK)
db = assert(sqlite.open(path))
local statement = assert(db:prepare("SELECT name FROM things"))
assert(statement:step() == sqlite.ROW)
assert(statement:get_value(0) == "Hello, world", "SQLite persistence failed")
assert(statement:finalize() == sqlite.OK)
assert(db:close() == sqlite.OK)
assert(os.remove(path))

local listener = assert(socket.bind("127.0.0.1", 0))
assert(listener:settimeout(5))
local _, port = listener:getsockname()
local client = assert(socket.tcp())
assert(client:settimeout(5))
assert(client:connect("127.0.0.1", port))
local server = assert(listener:accept())
assert(server:settimeout(5))
for _, pair in ipairs({{client, server}, {server, client}}) do
    local message = json.encode({message = "Hello, world"}) .. "\n"
    local sent = 0
    while sent < #message do
        sent = assert(pair[1]:send(message, sent + 1))
    end
    local received = assert(pair[2]:receive("*l"))
    assert(json.decode(received).message == "Hello, world", "JSON over TCP failed")
end
assert(server:close())
assert(client:close())
assert(listener:close())
print("Lua environment ready: SQLite persistence and JSON over TCP verified.")
