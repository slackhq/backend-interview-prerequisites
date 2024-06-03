import { AsyncDatabase } from "promised-sqlite3";
import { readFile } from "node:fs/promises";
import * as path from "node:path";

let instance: AsyncDatabase | undefined;
let schema: string;

const dbFile = process.env.NODE_ENV === "test" ? ":memory:" : path.resolve(__dirname, "datastore.db");

export async function sqlConnection(): Promise<AsyncDatabase> {
  if (!instance) {
    try {
      instance = await AsyncDatabase.open(dbFile);
      const schemaPath = path.resolve(__dirname, "schema.sql");
      schema =
        schema || (await readFile(schemaPath, { encoding: "utf-8" }));
      const statements = schema
        .split(";")
        .map((s) => s.trim())
        .filter((s) => !!s)
        .map((s) => `${s};`);
      for (const stmt of statements) {
        await instance?.run(stmt);
      }
    } catch (err) {
      console.error("Unable to initialize database", err);
      throw err;
    }
  }
  return instance;
}

export function resetDb() {
  instance = undefined;
}
