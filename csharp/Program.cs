using System.Net;
using System.Net.Sockets;
using System.Text.Json;
using Microsoft.Data.Sqlite;

var path = Path.Combine(Path.GetTempPath(), $"interview-{Guid.NewGuid()}.db");
try
{
    var connectionString = new SqliteConnectionStringBuilder
    {
        DataSource = path,
        Pooling = false
    }.ToString();
    using (var db = new SqliteConnection(connectionString))
    {
        db.Open();
        using var command = db.CreateCommand();
        command.CommandText = "CREATE TABLE things (name TEXT); INSERT INTO things VALUES ('Hello, world');";
        command.ExecuteNonQuery();
    }
    using (var db = new SqliteConnection(connectionString))
    {
        db.Open();
        using var command = db.CreateCommand();
        command.CommandText = "SELECT name FROM things";
        if ((string?)command.ExecuteScalar() != "Hello, world")
            throw new InvalidOperationException("SQLite persistence failed");
    }
    using var listener = new TcpListener(IPAddress.Loopback, 0);
    listener.Start();
    using var client = new TcpClient();
    await client.ConnectAsync((IPEndPoint)listener.LocalEndpoint);
    using var server = await listener.AcceptTcpClientAsync();
    using var timeout = new CancellationTokenSource(TimeSpan.FromSeconds(5));
    var message = JsonSerializer.SerializeToUtf8Bytes(new { message = "Hello, world" });
    foreach (var (sender, receiver) in new[] { (client, server), (server, client) })
    {
        await sender.GetStream().WriteAsync(message, timeout.Token);
        var buffer = new byte[message.Length];
        await receiver.GetStream().ReadExactlyAsync(buffer, timeout.Token);
        using var json = JsonDocument.Parse(buffer);
        if (json.RootElement.GetProperty("message").GetString() != "Hello, world")
            throw new InvalidOperationException("JSON over TCP failed");
    }
    Console.WriteLine("C# environment ready: SQLite persistence and JSON over TCP verified.");
}
finally
{
    File.Delete(path);
}
