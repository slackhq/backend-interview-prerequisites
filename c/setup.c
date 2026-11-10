#define _POSIX_C_SOURCE 200809L
#include <arpa/inet.h>
#include <json-c/json.h>
#include <sqlite3.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <unistd.h>

static void check(int ok, const char *operation) {
    if (!ok) {
        fprintf(stderr, "Failed: %s\n", operation);
        exit(EXIT_FAILURE);
    }
}

static void exchange(int sender, int receiver, const char *message) {
    size_t length = strlen(message), sent = 0, received = 0;
    char buffer[128] = {0};
    check(length < sizeof(buffer), "message size");
    while (sent < length) {
        ssize_t count = send(sender, message + sent, length - sent, 0);
        check(count > 0, "TCP send");
        sent += (size_t)count;
    }
    while (received < length) {
        ssize_t count = recv(receiver, buffer + received, length - received, 0);
        check(count > 0, "TCP receive");
        received += (size_t)count;
    }
    struct json_object *json = json_tokener_parse(buffer), *value = NULL;
    check(json != NULL, "JSON parse");
    check(json_object_object_get_ex(json, "message", &value), "JSON field");
    check(strcmp(json_object_get_string(value), "Hello, world") == 0, "JSON value");
    json_object_put(json);
}

int main(void) {
    char path[] = "/tmp/interview-c-XXXXXX";
    int file = mkstemp(path);
    check(file >= 0, "temporary database");
    close(file);
    sqlite3 *db = NULL;
    check(sqlite3_open(path, &db) == SQLITE_OK, "SQLite open");
    check(sqlite3_exec(db, "CREATE TABLE things (name TEXT);"
                          "INSERT INTO things VALUES ('Hello, world');",
                      NULL, NULL, NULL) == SQLITE_OK, "SQLite write");
    check(sqlite3_close(db) == SQLITE_OK, "SQLite close");
    check(sqlite3_open(path, &db) == SQLITE_OK, "SQLite reopen");
    sqlite3_stmt *statement = NULL;
    check(sqlite3_prepare_v2(db, "SELECT name FROM things", -1, &statement, NULL)
              == SQLITE_OK, "SQLite prepare");
    check(sqlite3_step(statement) == SQLITE_ROW, "SQLite read");
    check(strcmp((const char *)sqlite3_column_text(statement, 0), "Hello, world") == 0,
          "SQLite persistence");
    sqlite3_finalize(statement);
    sqlite3_close(db);
    check(unlink(path) == 0, "database cleanup");

    int listener = socket(AF_INET, SOCK_STREAM, 0);
    check(listener >= 0, "TCP listener");
    struct sockaddr_in address = {0};
    address.sin_family = AF_INET;
    check(inet_pton(AF_INET, "127.0.0.1", &address.sin_addr) == 1, "loopback address");
    check(bind(listener, (struct sockaddr *)&address, sizeof(address)) == 0, "TCP bind");
    check(listen(listener, 1) == 0, "TCP listen");
    socklen_t size = sizeof(address);
    check(getsockname(listener, (struct sockaddr *)&address, &size) == 0, "TCP port");
    int client = socket(AF_INET, SOCK_STREAM, 0);
    check(client >= 0, "TCP client");
    check(connect(client, (struct sockaddr *)&address, sizeof(address)) == 0, "TCP connect");
    int server = accept(listener, NULL, NULL);
    check(server >= 0, "TCP accept");
    struct timeval timeout = {.tv_sec = 5};
    check(setsockopt(client, SOL_SOCKET, SO_RCVTIMEO, &timeout, sizeof(timeout)) == 0,
          "client timeout");
    check(setsockopt(server, SOL_SOCKET, SO_RCVTIMEO, &timeout, sizeof(timeout)) == 0,
          "server timeout");
    struct json_object *json = json_object_new_object();
    json_object_object_add(json, "message", json_object_new_string("Hello, world"));
    const char *message = json_object_to_json_string_ext(json, JSON_C_TO_STRING_PLAIN);
    exchange(client, server, message);
    exchange(server, client, message);
    json_object_put(json);
    close(server);
    close(client);
    close(listener);
    puts("C environment ready: SQLite persistence and JSON over TCP verified.");
    return EXIT_SUCCESS;
}
