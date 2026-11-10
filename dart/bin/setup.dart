import 'dart:convert';
import 'dart:io';

import 'package:sqlite3/sqlite3.dart';

Future<void> main() async {
  final directory = Directory.systemTemp.createTempSync('interview-dart-');
  try {
    final path = '${directory.path}/setup.db';
    var db = sqlite3.open(path);
    try {
      db.execute('CREATE TABLE things (name TEXT)');
      db.execute('INSERT INTO things VALUES (?)', ['Hello, world']);
    } finally {
      db.close();
    }
    db = sqlite3.open(path);
    try {
      if (db.select('SELECT name FROM things').single['name'] !=
          'Hello, world') {
        throw StateError('SQLite persistence failed');
      }
    } finally {
      db.close();
    }
    final listener = await ServerSocket.bind(InternetAddress.loopbackIPv4, 0);
    Socket? client;
    Socket? server;
    try {
      final accepted = listener.first;
      client = await Socket.connect(
        InternetAddress.loopbackIPv4,
        listener.port,
        timeout: const Duration(seconds: 5),
      );
      server = await accepted.timeout(const Duration(seconds: 5));
      final serverLines = server
          .cast<List<int>>()
          .transform(utf8.decoder)
          .transform(const LineSplitter());
      final clientLines = client
          .cast<List<int>>()
          .transform(utf8.decoder)
          .transform(const LineSplitter());
      client.writeln(jsonEncode({'message': 'Hello, world'}));
      await client.flush();
      final request = jsonDecode(
        await serverLines.first.timeout(const Duration(seconds: 5)),
      );
      if (request['message'] != 'Hello, world')
        throw StateError('TCP request failed');
      server.writeln(jsonEncode(request));
      await server.flush();
      final response = jsonDecode(
        await clientLines.first.timeout(const Duration(seconds: 5)),
      );
      if (response['message'] != 'Hello, world')
        throw StateError('TCP response failed');
    } finally {
      client?.destroy();
      server?.destroy();
      await listener.close();
    }
    print(
      'Dart environment ready: SQLite persistence and JSON over TCP verified.',
    );
  } finally {
    directory.deleteSync(recursive: true);
  }
}
