# notRedis — Design Document

## Goals

- Build a Redis-compatible in-memory datastore
- Focus on correctness, concurrency, and performance
- Support a subset of Redis commands initially

## Non-Goals

- Full Redis feature parity
- Clustering (for now)
- Production-grade security

## Architecture Overview

- TCP server listening on port 6379
- One connection per client
- Command parsing using RESP
- In-memory key-value store
- Command dispatcher

## Request Flow

1. Client connects via TCP
2. Server reads raw bytes
3. RESP parser converts bytes → command
4. Dispatcher routes command
5. Store executes command
6. Response encoded in RESP
7. Sent back to client

## Concurrency Model (Initial)

- One goroutine per connection, Not implementing Event Loops like redis does
- Shared in-memory store
- Mutex-protected writes

## Future Work

- Persistence (AOF)
- Replication
- Eviction policies
