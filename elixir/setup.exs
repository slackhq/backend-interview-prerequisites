path = Path.join(System.tmp_dir!(), "interview-elixir-#{System.unique_integer([:positive])}.db")

try do
  {:ok, db} = Exqlite.Sqlite3.open(path)

  try do
    :ok = Exqlite.Sqlite3.execute(db, "CREATE TABLE things (name TEXT)")
    :ok = Exqlite.Sqlite3.execute(db, "INSERT INTO things VALUES ('Hello, world')")
  after
    :ok = Exqlite.Sqlite3.close(db)
  end

  {:ok, db} = Exqlite.Sqlite3.open(path)

  try do
    {:ok, statement} = Exqlite.Sqlite3.prepare(db, "SELECT name FROM things")
    {:row, ["Hello, world"]} = Exqlite.Sqlite3.step(db, statement)
    :ok = Exqlite.Sqlite3.release(db, statement)
  after
    :ok = Exqlite.Sqlite3.close(db)
  end

  {:ok, listener} =
    :gen_tcp.listen(0, [:binary, active: false, packet: :line, ip: {127, 0, 0, 1}])

  try do
    {:ok, port} = :inet.port(listener)

    {:ok, client} =
      :gen_tcp.connect({127, 0, 0, 1}, port, [:binary, active: false, packet: :line], 5000)

    try do
      {:ok, server} = :gen_tcp.accept(listener, 5000)

      try do
        for {sender, receiver} <- [{client, server}, {server, client}] do
          :ok = :gen_tcp.send(sender, Jason.encode!(%{message: "Hello, world"}) <> "\n")
          {:ok, message} = :gen_tcp.recv(receiver, 0, 5000)
          %{"message" => "Hello, world"} = Jason.decode!(message)
        end
      after
        :gen_tcp.close(server)
      end
    after
      :gen_tcp.close(client)
    end
  after
    :gen_tcp.close(listener)
  end

  IO.puts("Elixir environment ready: SQLite persistence and JSON over TCP verified.")
after
  File.rm!(path)
end
