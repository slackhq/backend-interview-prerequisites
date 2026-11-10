# Scala environment check

Requires JDK 11+ and [sbt](https://www.scala-sbt.org/download/).
The build pins sbt 1.10.11 and Scala 3.3.6, and restores ujson and SQLite JDBC.

From this directory, run:

```sh
sbt run
```

The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output, `java -version`, and your OS version to your
recruiting coordinator. Ensure `JAVA_HOME` points to your JDK.
