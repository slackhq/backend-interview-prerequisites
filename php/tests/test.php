<?php

require_once(__DIR__.'/Client.php');
require_once(__DIR__.'/Expect.php');
require_once(__DIR__.'/../src/SQLConnectionManager.php');

use BackendInterviewPrerequisites\SQLConnectionManager;

function test_database() {
    $sqlite = SQLConnectionManager::getSQLConnection();
    try {
        $stmt = $sqlite->prepare("INSERT INTO messages (msg) VALUES (\"testing 123\")");
        expect($stmt->execute())->toBeTruthy('Inserting a message works');

        $ret = $sqlite->query("SELECT * FROM messages ORDER BY id DESC LIMIT 1");
        expect($ret)->toBeTruthy('Querying the messages table works');

        $row = $ret->fetchArray(SQLITE3_ASSOC);
        expect($row['msg'])->toBeSame('testing 123', 'The expected data was persisted');
    } finally {
        SQLConnectionManager::cleanup();
    }
}

function test_server() {
    $client = new Client();
    $client->connect();

    $ret = $response = $client->sendRequest('ping 1');
    expect($ret)->toBeSame('ping 1', 'Sends and receives a message');

    $ret = $response = $client->sendRequest('ping 2');
    expect($ret)->toBeSame('ping 2', 'Sends and receives a second message');

    $client->disconnect();
}

function test() {
    test_database();
    test_server();

    if (ExpectObj::$failures) {
        echo "Tests failed: ".ExpectObj::$failures. " assertions failed\n";
        exit(1);
    } else {
        echo "Tests passed! 🎉\n";
    }
}

test();