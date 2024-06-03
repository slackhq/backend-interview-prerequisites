# Backend interview prep (TypeScript / JavaScript)

This is a sample project intended to verify environment compatibility with the Slack backend coding interview. While 
this is written in TypeScript, our JavaScript repository uses the same environment. If you can successfully run `npm test` and 
`npm run dev-ts` in this project, you should have no problem with running either the JavaScript or TypeScript projects.

## Requirements

- Locally installed nodejs >= 20.10.0 with npm >= 10.2.3

## Installation

The easiest way to install is via [nvm](https://github.com/nvm-sh/nvm).

1. Install nvm via the instructions in the [nvm github README](https://github.com/nvm-sh/nvm)
2. Navigate to this directory in a terminal and run `nvm install`. The correct version should be detected from the `.nvmrc` file and installed

## Running the server & the test suite

To install dependencies, run `npm i` from this directory

To start up the server in a watch mode: `npm run dev-ts`

Whenver a file is changed, the server will automatically reload with the latest source

To run unit tests:

- `npm test` or `npm test -- --watch` (for automatic execution on file change)

## Data Storage

Data is stored in a local sqlite3 database, `datastore.db`. The schema is in `src/database/schema.sql`. During test 
execution, an in-memory database is used.  Outside of tests, if anything is written to the database, it will be
stored in `src/database/datastore.db`.

## Client API

The following API requests can be sent by clients to interact with the server.

- `hello.get` 
- `hello.post`

## Invoking the API

The API is exposed over a RESTful(ish) HTTP server and a raw TCP socket interface that can be locally 
invoked via netcat.  Read on for information about how to invoke this:

### Over TCP

Each API method is exposed over TCP at `localhost:8000` (port can be changed via `TCP_PORT` env var). To invoke an API method,
send a line containing a JSON encoded object with a `type` field that indicates the API method you want to invoke.

```
❯ nc localhost 8000
{"type":"hello.get","testing":"123"}
{"ok":true,"msg":"hello","req":{"type":"hello.get","testing":"123"}}

{"type":"hello.post","testing":"555"}
{"ok":true,"msg":"hello","req":{"type":"hello.post","testing":"555"}}
```

### Over HTTP

Each API method is exposed over HTTP at `localhost:3033/api/${METHOD_NAME}`. Differerent APIs use different HTTP methods as noted in their handler
definitions (see `src/api/$methodName`). User sessions are tracked using the `user-id` cookie. Some example invocations are noted below.

You can optionally pipe to `jq` (if locally installed) for JSON formatting.

#### `hello.post`

```
❯ curl -d '{"name": "ada"}' -H 'content-type: application/json' -X POST localhost:3033/api/hello.post | jq
{
  "ok": true,
  "msg": "hello",
  "params": {},
  "body": {
    "name": "ada"
  }
}
```

#### `hello.get`

```
❯ curl 'localhost:3033/api/hello.get?param1=val1' | jq
{
  "ok": true,
  "msg": "hello",
  "params": {
    "param1": "val1"
  }
}
```

## Schema Changes

To make a schema change, you can modify the `schema.sql` file in the root `nano-slack` directory. The schema is
automatically recreated from this file on each run of the tests or when the `datastore.db` file is cleared on
the server. You can manually do this by `rm src/database/datastore.db` if you'd like to wipe to local database.

## Querying the database

After you have run the server at least once, a `datastore.db` file will exist in `src/database/datastore.db`. You
can use `sqlite3 -column -header datastore.db` to open and explore the database. This will show the database as
it looks after the most server run.
