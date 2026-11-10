# C# environment check

Install the [.NET 8 SDK](https://dotnet.microsoft.com/download/dotnet/8.0)
for your OS. The SDK restores Microsoft.Data.Sqlite 8.0.11 and its native
SQLite dependency.

From this directory, run:

```sh
dotnet restore --locked-mode
dotnet run --no-restore
```

The check verifies SQLite persistence and bidirectional JSON over TCP using
a temporary database and an automatically assigned loopback port.
If it fails, send the output, `dotnet --info`, and your OS version to your
recruiting coordinator.
