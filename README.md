# 🚀 notRedis – Our Attempt to Build a Simple Key-Value Store in Go

**notRedis** is a small learning project where we (Gokul & Shreeom) are trying to build our own in-memory key-value store, inspired by Redis, but made from scratch in **Go**.

We are beginners in systems programming, so this project starts very simple:
just a map, a server, and a few basic commands.
As we learn more, we’ll keep adding features like TTL, concurrency handling, caching, and maybe persistence.

The goal isn’t to replace Redis — it’s to understand **how Redis-like systems actually work under the hood** and become better backend engineers.

---

# ⚙️ What it does right now (or soon will)

* Basic **GET / SET / DELETE** operations
* Simple in-memory storage
* HTTP-based API
* Very straightforward code that we can understand and improve

---

# 🌱 What we plan to add as we learn

* **TTL (expire a key after some time)**
* **Concurrent access support** using goroutines + locks
* **Background cleanup of expired keys**
* **LRU/LFU caching** (optional)
* **Persistence using a simple log file**
* **Custom TCP protocol** (advanced, later)
* **Sharding/consistent hashing** (much later)

---

# 🎯 Why we’re building this

We want to learn:

* How a simple database works
* How to build server software in Go
* How concurrency, locking, and goroutines work
* How Redis-like systems store and manage data
* How real backend systems are designed

This project is our way of getting hands-on practice by actually building, breaking, and fixing things ourselves.

---

# 🚧 Status

We’re just starting out 😄
The code will grow slowly as we understand more concepts and try new ideas.
