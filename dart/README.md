# Dart environment check

Install [Dart SDK 3.13+](https://dart.dev/get-dart).
The sqlite3 package downloads its native SQLite library during the build;
the first run needs network access.

From this directory, run:

```sh
dart pub get
dart run bin/setup.dart
```

The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output, `dart --version`, and your OS version to your
recruiting coordinator.
