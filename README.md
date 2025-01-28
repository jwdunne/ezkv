# ezkv

A distributed key-value store implemented as a learning project. This means there is a focus on correctness over performance.

v0.1:

- Single-node KV store
- Write-ahead log
- Transactions
- String-only values

v0.2:

- Server
- Small command language (SET, GET, DEL, KEYS, etc)

v0.3:

- Two-node KV store
- Network communication
- Primary-backup replication
- Two-phase commit protocol