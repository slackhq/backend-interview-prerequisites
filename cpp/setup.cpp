#include <arpa/inet.h>
#include <nlohmann/json.hpp>
#include <sqlite3.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <unistd.h>

#include <cstdlib>
#include <iostream>
#include <stdexcept>
#include <string>

void check(bool ok, const char* operation) {
    if (!ok) throw std::runtime_error(operation);
}

void exchange(int sender, int receiver, const std::string& message) {
    size_t sent = 0, received = 0;
    std::string buffer(message.size(), '\0');
    while (sent < message.size()) {
        auto count = send(sender, message.data() + sent, message.size() - sent, 0);
        check(count > 0, "TCP send");
        sent += static_cast<size_t>(count);
    }
    while (received < buffer.size()) {
        auto count = recv(receiver, buffer.data() + received, buffer.size() - received, 0);
        check(count > 0, "TCP receive");
        received += static_cast<size_t>(count);
    }
    check(nlohmann::json::parse(buffer).at("message") == "Hello, world", "JSON value");
}

int main() {
    try {
        char path[] = "/tmp/interview-cpp-XXXXXX";
        int file = mkstemp(path);
        check(file >= 0, "temporary database");
        close(file);
        sqlite3* db = nullptr;
        check(sqlite3_open(path, &db) == SQLITE_OK, "SQLite open");
        check(sqlite3_exec(db, "CREATE TABLE things (name TEXT);"
                              "INSERT INTO things VALUES ('Hello, world');",
                          nullptr, nullptr, nullptr) == SQLITE_OK, "SQLite write");
        check(sqlite3_close(db) == SQLITE_OK, "SQLite close");
        check(sqlite3_open(path, &db) == SQLITE_OK, "SQLite reopen");
        sqlite3_stmt* statement = nullptr;
        check(sqlite3_prepare_v2(db, "SELECT name FROM things", -1, &statement, nullptr)
                  == SQLITE_OK, "SQLite prepare");
        check(sqlite3_step(statement) == SQLITE_ROW, "SQLite read");
        check(std::string(reinterpret_cast<const char*>(sqlite3_column_text(statement, 0)))
                  == "Hello, world", "SQLite persistence");
        sqlite3_finalize(statement);
        sqlite3_close(db);
        check(unlink(path) == 0, "database cleanup");

        int listener = socket(AF_INET, SOCK_STREAM, 0);
        check(listener >= 0, "TCP listener");
        sockaddr_in address{};
        address.sin_family = AF_INET;
        address.sin_addr.s_addr = htonl(INADDR_LOOPBACK);
        check(bind(listener, reinterpret_cast<sockaddr*>(&address), sizeof(address)) == 0,
              "TCP bind");
        check(listen(listener, 1) == 0, "TCP listen");
        socklen_t size = sizeof(address);
        check(getsockname(listener, reinterpret_cast<sockaddr*>(&address), &size) == 0,
              "TCP port");
        int client = socket(AF_INET, SOCK_STREAM, 0);
        check(client >= 0, "TCP client");
        check(connect(client, reinterpret_cast<sockaddr*>(&address), sizeof(address)) == 0,
              "TCP connect");
        int server = accept(listener, nullptr, nullptr);
        check(server >= 0, "TCP accept");
        timeval timeout{5, 0};
        check(setsockopt(client, SOL_SOCKET, SO_RCVTIMEO, &timeout, sizeof(timeout)) == 0,
              "client timeout");
        check(setsockopt(server, SOL_SOCKET, SO_RCVTIMEO, &timeout, sizeof(timeout)) == 0,
              "server timeout");
        auto message = nlohmann::json{{"message", "Hello, world"}}.dump();
        exchange(client, server, message);
        exchange(server, client, message);
        close(server);
        close(client);
        close(listener);
        std::cout << "C++ environment ready: SQLite persistence and JSON over TCP verified.\n";
        return EXIT_SUCCESS;
    } catch (const std::exception& error) {
        std::cerr << "Failed: " << error.what() << '\n';
        return EXIT_FAILURE;
    }
}
