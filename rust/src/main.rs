use rusqlite::Connection;
use serde_json::{json, Value};
use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};
use std::time::Duration;

fn exchange(
    sender: &mut TcpStream,
    receiver: &mut TcpStream,
) -> Result<(), Box<dyn std::error::Error>> {
    let message = serde_json::to_vec(&json!({"message": "Hello, world"}))?;
    sender.write_all(&message)?;
    let mut buffer = vec![0; message.len()];
    receiver.read_exact(&mut buffer)?;
    let json: Value = serde_json::from_slice(&buffer)?;
    assert_eq!(json["message"], "Hello, world", "JSON over TCP");
    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let path = std::env::temp_dir().join(format!("interview-rust-{}.db", std::process::id()));
    let result = (|| -> Result<(), Box<dyn std::error::Error>> {
        {
            let db = Connection::open(&path)?;
            db.execute_batch(
                "CREATE TABLE things (name TEXT); INSERT INTO things VALUES ('Hello, world');",
            )?;
        }
        {
            let db = Connection::open(&path)?;
            let name: String = db.query_row("SELECT name FROM things", [], |row| row.get(0))?;
            assert_eq!(name, "Hello, world", "SQLite persistence");
        }
        let listener = TcpListener::bind("127.0.0.1:0")?;
        let mut client =
            TcpStream::connect_timeout(&listener.local_addr()?, Duration::from_secs(5))?;
        let (mut server, _) = listener.accept()?;
        client.set_read_timeout(Some(Duration::from_secs(5)))?;
        server.set_read_timeout(Some(Duration::from_secs(5)))?;
        exchange(&mut client, &mut server)?;
        exchange(&mut server, &mut client)?;
        println!("Rust environment ready: SQLite persistence and JSON over TCP verified.");
        Ok(())
    })();
    std::fs::remove_file(path)?;
    result
}
