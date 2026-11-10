# Kotlin environment check

Requires JDK 11. The checked-in Gradle wrapper downloads Gradle 8.11.1,
Kotlin 2.1.0, org.json, and SQLite JDBC on first use.

From this directory, run:

```sh
./gradlew run
```

On Windows, use `gradlew.bat run`.

The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output, `java -version`, and your OS version to your
recruiting coordinator. Ensure `JAVA_HOME` points to your JDK.
