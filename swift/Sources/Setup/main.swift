import CSQLite
import Foundation
#if os(Linux)
import Glibc
let socketType = Int32(SOCK_STREAM.rawValue)
#else
import Darwin
let socketType = SOCK_STREAM
#endif

func exchange(_ sender: Int32, _ receiver: Int32) throws {
    let message = try JSONSerialization.data(withJSONObject: ["message": "Hello, world"])
    var sent = 0
    while sent < message.count {
        let count = message.withUnsafeBytes {
            send(sender, $0.baseAddress!.advanced(by: sent), message.count - sent, 0)
        }
        precondition(count > 0, "TCP send failed")
        sent += count
    }
    var buffer = Data(count: message.count)
    var received = 0
    while received < message.count {
        let remaining = message.count - received
        let count = buffer.withUnsafeMutableBytes {
            recv(receiver, $0.baseAddress!.advanced(by: received), remaining, 0)
        }
        precondition(count > 0, "TCP receive failed")
        received += count
    }
    let json = try JSONSerialization.jsonObject(with: buffer) as? [String: String]
    precondition(json?["message"] == "Hello, world", "JSON over TCP failed")
}

let path = FileManager.default.temporaryDirectory.appendingPathComponent("interview-\(UUID()).db").path
defer { try? FileManager.default.removeItem(atPath: path) }
var db: OpaquePointer?
precondition(sqlite3_open(path, &db) == SQLITE_OK, "SQLite open failed")
precondition(sqlite3_exec(db, "CREATE TABLE things (name TEXT); INSERT INTO things VALUES ('Hello, world');",
                         nil, nil, nil) == SQLITE_OK, "SQLite write failed")
precondition(sqlite3_close(db) == SQLITE_OK, "SQLite close failed")
precondition(sqlite3_open(path, &db) == SQLITE_OK, "SQLite reopen failed")
var statement: OpaquePointer?
precondition(sqlite3_prepare_v2(db, "SELECT name FROM things", -1, &statement, nil) == SQLITE_OK,
             "SQLite prepare failed")
precondition(sqlite3_step(statement) == SQLITE_ROW, "SQLite read failed")
precondition(String(cString: sqlite3_column_text(statement, 0)) == "Hello, world",
             "SQLite persistence failed")
sqlite3_finalize(statement)
sqlite3_close(db)

let listener = socket(AF_INET, socketType, 0)
precondition(listener >= 0, "TCP listener failed")
defer { close(listener) }
var address = sockaddr_in()
address.sin_family = sa_family_t(AF_INET)
precondition(inet_pton(AF_INET, "127.0.0.1", &address.sin_addr) == 1)
let addressSize = socklen_t(MemoryLayout<sockaddr_in>.size)
let bound = withUnsafePointer(to: &address) {
    $0.withMemoryRebound(to: sockaddr.self, capacity: 1) {
        bind(listener, $0, addressSize)
    }
}
precondition(bound == 0, "TCP bind failed")
precondition(listen(listener, 1) == 0, "TCP listen failed")
var size = addressSize
let located = withUnsafeMutablePointer(to: &address) {
    $0.withMemoryRebound(to: sockaddr.self, capacity: 1) {
        getsockname(listener, $0, &size)
    }
}
precondition(located == 0, "TCP port failed")
let client = socket(AF_INET, socketType, 0)
precondition(client >= 0, "TCP client failed")
defer { close(client) }
let connected = withUnsafePointer(to: &address) {
    $0.withMemoryRebound(to: sockaddr.self, capacity: 1) {
        connect(client, $0, addressSize)
    }
}
precondition(connected == 0, "TCP connect failed")
let server = accept(listener, nil, nil)
precondition(server >= 0, "TCP accept failed")
defer { close(server) }
var timeout = timeval(tv_sec: 5, tv_usec: 0)
for socket in [client, server] {
    precondition(setsockopt(socket, SOL_SOCKET, SO_RCVTIMEO, &timeout,
                           socklen_t(MemoryLayout<timeval>.size)) == 0, "TCP timeout failed")
}
try exchange(client, server)
try exchange(server, client)
print("Swift environment ready: SQLite persistence and JSON over TCP verified.")
