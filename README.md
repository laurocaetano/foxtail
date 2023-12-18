[![Go](https://github.com/laurocaetano/foxtail/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/laurocaetano/foxtail/actions/workflows/go.yml)

# foxtail
Foxtail is a distributed KV database

## Disclaimer
This project is not meant to be used in production. Please be aware of that!
This project aims to sharpen my Go and distributed systems skills.

## Plan
The original idea is to take an incremental approach while developing this database,
and the first high-level plan is as follows:

1. Implement a basic set of functions for an in-memory KV database
1. Expose the KV functionality over a high-level API (gRPC, HTTP, etc)
1. Implement a client for the public API
1. Make things more interesting, by throwing Raft on top of the existing codebase
1. Write down an extensive set of acceptance tests for the new distributed KV setup
1. Find out a way to implement durable storage, but still offer the in-memory option

### Mode details

This section will be updated as I make progress in the areas listed above.
The goal is to build a complex system but in an incremental fashion. I expect to
discover more requirements on the way, so this means the design of the DB can change
quite frequently.

#### In-memory KV database

To keep things simple at the beginning, Foxtail will only provide
a set of functions:

`PUT(key, value)` where `key` is a `string` and `value` is a `binary`. It saves
the new value into the given key. If the value does not previously exist, it
will create a new entry on the database, otherwise, the existent one will be
overwritten.

`GET(key)` where `key` is a `string`. It will return the `binary` associated with
that given `key`, and in case the key does not exist, returns `nil`.

`DELETE(key)` where `key` is a `string`. It will delete an existing key and its
associated value. If the `key` does not exist, it returns `nil`.

The actual
constraints for the `key` and `value`, such as max length, allowed characters and
etc, are yet to be defined.
