# DDIA Chapter 3 — Storage and Retrieval

## User Summary (Jul 21)
- Log-structured (LSM): fast writes, background compaction, memtable flushed to disk as SSTables
- Page-oriented (B-tree): in-place updates, aligned to disk/OS blocks, better for reads
- Key difference: append-only (LSM) vs in-place update (B-tree)

## Quiz

### Q1: LSM write path — two structures and what happens to each?

**Answer:** Inserts written to WAL, then to memtable (red-black/skip list), appended to most recent SSTable with sparse index. SSTable overflow triggers compaction.

**Feedback:** Correct. WAL for durability, memtable is in-memory sorted structure, SSTable = Sorted String Table (not "String Search Table"). Compaction merges SSTables, deduplicates (newer wins), removes tombstones — not just "moving older records."

---

### Q2: Why does LSM need compaction?

**Answer:** Inserts and updates are append-only, SSTables on disk grow and read queries degrade. Compaction merges and dedupes records in SSTables.

**Feedback:** Correct. Two problems compaction solves: (1) storage bloat from duplicate keys across SSTables, (2) read performance — more SSTables = more files to check on read. Compaction merges, dedupes (newer version wins), drops tombstones.

---

### Q3: B-tree write amplification — why is updating a single key expensive?

**Answer:** Inserting a single key in a full page requires a page split and updates to any index reference by it or replacing it with a forwarding index — all of which require additional writes beyond the insert of the key itself.

**Feedback:** Correct. Core issue: one logical write becomes multiple physical writes. Page split = write 2 new pages + update parent index entry. Also: WAL write for durability, entire page rewritten even for partial update. Write amplification = ratio of physical bytes written to logical bytes.

---

### Q4: LSM vs B-trees — read performance tradeoff?

**Answer:** LSM reads could require scanning back to oldest SSTable for recently-updated records. Recent writes read fast, especially consulting memtable before final lookup. B-trees require walking a fixed height — each descent accesses a disk page. Clustered indexes store entire row on leaves; covering indexes read only pertinent columns.

**Feedback:** Correct. LSM read cost depends on how many SSTables to check (Bloom filters help skip SSTables that don't contain the key). Memtable check is fast (in-memory). B-tree reads are predictable O(log n) with ~3-4 disk reads for millions of keys. Tradeoff: LSM reads get slower as data grows (more SSTables), B-tree reads stay constant height.

---

### Q5: Sparse index vs dense index?

**Answer:** Sparse index for LSM — requires an index lookup followed by range scan to queried record. Sparse indexes don't reference each and every record in an SSTable. B-trees use dense indexes — every row has a primary key or index.

**Feedback:** Correct. Sparse index = one entry per chunk/page of records (not per record). SSTables use this because records are sorted — find the chunk, then scan within it. Dense index = one entry per record. B-trees use this — every key exists at a leaf node. Tradeoff: sparse = smaller index (fits in memory better), dense = direct lookup (no range scan needed).

---

# DDIA Chapter 7 — Transactions

## Quiz

### Q1: What are the four properties of ACID?

**Answer:** Atomicity — transaction either succeeds or fails, no partial updates. Consistency — transactions should not alter referential integrity or violate constraints. Isolation — concurrent transactions execute as though queued and run sequentially. Durability — data should be stored indefinitely until manual deletion.

**Feedback:** Atomicity, Consistency, Isolation — correct. Durability off: it means committed data survives crashes (written to non-volatile storage like WAL/disk), NOT "stored indefinitely until manual deletion" — that's a retention policy, not the D in ACID.

### Q2: What isolation levels exist and what anomalies does each prevent?

**Answer:** Read uncommitted — allows dirty reads (identical reads in single transaction return differing values). Read committed — every read gets a snapshot, prevents dirty reads. Snapshot isolation — reads get a snapshot, prevents dirty and non-repeatable reads (inserts/deletes in common range committed by concurrent transactions). SSI/serializable — prevents dirty and non-repeatable reads, executes rw-rw cycle detection to prevent lost updates and r/w skew.

**Feedback:** Correct progression. Each level prevents more anomalies. Key distinction: SSI adds dependency tracking (read-write cycles) on top of snapshot isolation. One nuance: SSI detects dependencies but the book calls them "pending" vs "committed" dependencies — cycle detection happens at commit time, not read time.

### Q3: What is a write skew anomaly?

**Answer:** Two concurrent transactions write data to separate rows that violate a conceptual/app layer invariant.

**Feedback:** Correct. Classic example: two doctors on call, each sees another doctor is on call, each removes themselves from on-call. Both commit — no doctors on call. Snapshot isolation doesn't catch it because neither transaction wrote to the other's row. Only SSI (serializable) detects the rw dependency cycle.

### Q4: What problem do serializable snapshots solve?

**Answer:** R/W skew.

**Feedback:** Correct. SSI adds rw dependency tracking on top of snapshot isolation, detecting cycles at commit time. This prevents write skew anomalies that snapshot isolation alone cannot catch.

### Q5: What is the difference between 2PL and SSI?

**Answer:** 2PL is pessimistic — reads acquire shared locks and upgrade to XL when writing. SSI is optimistic — transactions run to completion and commit as long as cycles are undetected.

**Feedback:** Correct. Key tradeoffs: 2PL blocks other transactions (deadlock risk, lower throughput). SSI doesn't block (higher throughput) but may abort more transactions under contention. SSI wins for read-heavy workloads; 2PL may be needed for write-heavy with strict consistency.

---

# DDIA Chapter 4 — Encoding and Evolution

## User Summary (Jul 22)
- Evolving data — old/new code reading old/new data formats
- Plain-text vs binary: space savings, versioning, data outlays
- Backward/forward compatibility within a single process reading multi-versioned data via a database, multiple processes sending old vs new formats via HTTP, RPC, SOAP, async message queues
- Dynamic vs static schema definitions for dynamically/statically-typed languages

## Quiz

### Q1: Backward compatibility (new code reads old data) vs forward compatibility (old code reads new data) — which is harder and why?

**Answer:** Old schema reading new data shouldn't be an issue if new fields have nullable/sensible defaults — new fields are ignored anyway. Type changes are breaking in both directions. Backward compatibility is harder.

**Feedback:** Forward compatibility is harder. New code encountering missing fields can use defaults — it's aware of what it might be missing. Old code encountering unknown fields/values has no concept of what to ignore — a new enum value, new field type, or restructured message causes silent failures (data dropped, misinterpreted bytes, wrong defaults). Backward failures are visible (new code crashes), forward failures are silent.

### Q2: RPC system with clients upgraded in waves — which compatibility direction matters more, and what schema design choice helps?

**Answer:** Forward compatibility — old schema reads new data, old code breaks. If RPC schema generation is dynamic, old code ignores new fields.

**Feedback:** Direction correct. Schema design choice is right idea but underspecified — the concrete mechanism is field numbers/tags (like Protobuf). New fields get new field numbers, old clients skip unknown field numbers silently. This is a wire format decision: encoding field identity by number/position rather than raw struct layout. It's not about the language being dynamic — it's about the schema definition language (Protobuf, Avro, Thrift) handling unknown fields by design.

### Q3: JSON payloads in DB — V1 has name/email, V2 adds phone (optional). Old code reads V2. Why does JSON handle this better than Protobuf — or does it?

**Answer:** JSON unmarshalling can be configured to throw error on extra/unknown properties — or not. Protobuf schema fields are versioned, thus are equally flexible.

**Feedback:** Correct pushback on the premise. JSON's advantage is no schema to coordinate — data is self-describing, no code generation or schema distribution needed. Protobuf handles this equally well via field numbers (skip unknowns, defaults for missing). The difference is operational: Protobuf requires schema distribution + build coordination between producer/consumer; JSON doesn't. Both handle version evolution — choice is simplicity (JSON) vs space + type safety (Protobuf), not compatibility.

### Q4: Plain-text vs binary log format — binary is 3x smaller. What's the tradeoff when debugging at 2 AM?

**Answer:** Human vs machine readable. Plain-text is a good starting point but binary might be useful if tools exist to read and pinpoint issues (say from a serialized stack trace).

**Feedback:** Correct. Core tradeoff: human vs machine readability. Plain-text degrades gracefully — `cat`, `grep`, `less` always work, zero tool dependency. Binary adds a deserializer dependency that may not be on the production box or may version-mismatch. At 2 AM under pressure, plain-text is safer. Binary wins for storage/transmission efficiency. Production systems often keep both — binary for hot path, plain-text exports for human debugging.

### Q5: Async message queue — producer sends V2 with new field, consumer is still V1. What happens to the new field, and is that safe?

**Answer:** Message queues are the communication method — but the underlying phenomenon of old code reading new data is unchanged. Old code skips new data.

**Feedback:** Correct — forward compatibility, same mechanism. Silent semantic bug in both RPC and async queue — old code skips new field, no visible failure. RPC vs async queue distinction is false; both fail silently. If broker waits for processed ack, async queue is functionally equivalent to RPC (request-response with different transport). The only real difference is delivery guarantee: fire-and-forget queues have worse failure mode (no retry path, message gone permanently) vs RPC where caller could retry or check response. But that's a delivery guarantee distinction, not RPC-vs-queue.

---

# DDIA Chapter 11 — Stream Processing

## Quiz

### Q1: What's the difference between a stream and a batch?

**Answer:** Streams are generalized batches in that they're unbounded.

**Feedback:** Correct. Batch = bounded input, processed and completed. Stream = unbounded input, continuous processing. Book describes streams as "a batch processing that is not bounded."

### Q2: What is a log-based message broker vs a message queue?

**Answer:** Message queues remain in RAM and have weaker delivery guarantees. Log-based message brokers persist to disk or keep WALs in event of crashes and guarantee at-least-once delivery.

**Feedback:** Correct. Key difference: log-based brokers retain messages (consumers can replay), queues discard after consumption. Log-based brokers also support consumer groups and partitioning natively.

### Q3: What is event time vs processing time?

**Answer:** Event time is a logical timestamp that is app-layer relevant. Processing time is when the stream queued the record.

**Feedback:** Correct in spirit. More precisely: event time = when the event actually occurred (set by producer). Processing time = when the stream processing system observes/records the record. The gap between them is "processing delay." Late-arriving events are events where event time is much older than processing time — windows must handle this.

### Q4: What is a window in stream processing?

**Answer:** A window chunks records by processing time in order to process downstream. Windows can be tumbling, fixed, or sliding and contain watermarks that allow time for straggler records to queue.

**Feedback:** Mostly correct. One key correction: windows chunk records by **event time**, not processing time — that's the whole point of windows. Watermarks signal "no events with timestamp < watermark are expected to arrive." Window types: tumbling (fixed, non-overlapping), sliding (fixed size, overlapping), session (variable size based on activity gaps).

### Q5: What problem do exactly-once semantics solve and how?

**Answer:** In the event of a failure in producer/broker/consumer, records should not be re-processed unless they can be guaranteed idempotent using a key or having broker keep track of consumer offsets.

**Feedback:** Correct. Exactly-once semantics solve duplicate processing during failures. Two mechanisms: (1) idempotent writes — dedup using unique ID/key, (2) atomic transactions — broker commits consumer offsets atomically with output writes. Kafka achieves this with transactional producers and consumer group coordination.

---

# Cumulative Quiz — Ch 3, 7, 11 (Jul 22)

### Q1: Write-heavy, read-rare key-value store — LSM-tree or B-tree?

**Answer:** LSM-tree maintained in RAM — but writes will require compaction and ordered inserts for records that may not be read at all. B-tree could potentially have write amplification and in-place updates would require versioning overhead.

**Feedback:** Right choice. LSM wins because sequential writes (memtable → SSTable flush) >> random I/O (B-tree page updates). Compaction is background cost. Read penalty (multiple SSTable lookups, mitigated by Bloom filters) is irrelevant when reads are rare. B-tree write amplification: one logical write = page rewrite + WAL + possible split + parent index update. "Versioning overhead" isn't the real B-tree cost — write amplification is. Key nuance: LSM writes are sequential **regardless of key distribution**. B-tree writes are sequential only for sorted key inserts (auto-increment IDs, timestamps). For random keys (UUIDs, hashes), B-tree writes are random I/O. B-tree advantage: sorted index and data are the same structure — one write maintains both. LSM has separate concerns (memtable sorted in memory, SSTables need compaction for global order on disk).

### Q2 (Ch 7): Snapshot isolation vs SSI — what anomaly does SSI catch that snapshot isolation doesn't?

**Answer:** (Repeat from Ch 7 quiz — already logged.) SI prevents dirty and non-repeatable reads via committed snapshots, but read/write skew can't be prevented without rw-rw cycle detection offered by SSI.

### Q3 (Ch 11): Real-time analytics pipeline, events arrive out of order. Why can't you window by processing time, and what problem do watermarks solve?

**Answer:** Events occur and are ordered by event time that the producer publishes — processing time is not relevant to actual app-layer event ordering. A watermark is a buffer that tells the broker how long to wait for stragglers.

**Feedback:** Correct. Watermarks give the system a principled cutoff: "no events older than T will arrive, safe to finalize window." Without watermarks, windows close on processing time and late events get silently dropped/misattributed. Watermark accuracy is a tuning tradeoff: too short → late events dropped (incomplete results); too long → windows hold open (increased latency).

### Q4 (Ch 3): What problem do Bloom filters solve for LSM-trees, and what's the tradeoff?

**Answer:** LSM-trees are self-balancing trees like red-black trees that allow looking up the block where a key might be stored in the latest corresponding sorted string table. Because SSTables are append-only, older writes might not be in the latest SSTable file. A Bloom filter produces false-negatives and tells if a key is likely stored in a given SSTable file.

**Feedback:** Core idea right — Bloom filters skip SSTables that definitely don't contain the key. Two corrections: (1) Bloom filters produce **false positives** (say "might be here" when it isn't → unnecessary disk read), never false negatives (never miss a real key). (2) LSM-tree isn't a self-balancing tree — the **memtable** is (red-black/skip list, in memory). The LSM-tree is the overall architecture: memtable + SSTables. Bloom filter sits on each SSTable. Tradeoff: bits per SSTable (space) vs unnecessary reads saved. ~10 bits/key → ~1% false positive rate.

### Q5 (Ch 7): Write skew under snapshot isolation — what happened and why can't SI catch it?

**Answer:** If T1's write of row B depended on its snapshot read (SI) of row A, and T2 updates row A while T1 is in transaction, then that is write skew. SI cannot detect cyclic RW dependencies.

**Feedback:** Correct. Key nuance: if T1 and T2 wrote to the **same** row, SI catches it (write-write conflict). Write skew requires overlapping reads + writes to **different** rows — SI tracks write-write conflicts but not read-write dependencies. SSI adds dependency tracking, detects cycle at commit time. Textbook example: two doctors on call, each sees the other is on call, each removes themselves — both commit, nobody on call.

### Q6 (Ch 11): Slow downstream consumer in a stream pipeline — what happens, how do you handle it?

**Answer:** Buffer overflow and spill to disk. Introduce backpressure.

**Feedback:** Correct. Tradeoffs in handling: (1) No backpressure → buffer fills → spill to disk (latency/I/O hit) or drop messages (data loss). (2) Pull-based backpressure (Kafka) → consumer pulls at own rate, broker holds messages, lag grows but no data loss. (3) Push-based with credit → consumer tells producer capacity, tighter coupling. Key insight: backpressure propagates cost to producer — slow consumer becomes pipeline throughput limit. Usually right tradeoff: slow down ingestion vs lose data or crash downstream.

### Q7 (Ch 3): B-tree page split — what triggers it, and why is it expensive?

**Answer:** An insert that takes the page over its fan-out limit — creates two left-aligned pages — old pages have to be moved on disk.

**Feedback:** Correct trigger. "Left-aligned" isn't standard — keys distribute roughly half-half across two new pages, middle key promoted to parent. "Old pages moved" is the expensive part but underspecified. Real cost chain: (1) two new pages written, (2) parent index updated with promoted key, (3) parent may also split (cascading), (4) WAL entry for durability, (5) free space allocation. One logical insert → 2+ page writes + WAL + possible parent splits = write amplification. This is why B-trees have worse write throughput than LSM for random inserts.

### Q8 (Ch 4): Adding a required field in Protobuf V2 — how do you handle backward compatibility?

**Answer:** Old clients reading V2 messages will ignore the field on unmarshalling even if it's required. On marshalling, old code will break unless required field has a default configured.

**Feedback:** Correct on both counts. Key insight: "required" is a schema-level constraint, not wire-level — wire format is just field numbers and bytes, old clients skip unknowns regardless. Breakage is on write side: old code generates V2 without new required field → schema validation fails at marshal. Best practice: never use `required` (Protobuf creators have called it a mistake). Use `optional` with sensible defaults. If must add required: two-phase rollout — deploy code populating default first, then mark required in schema.

---

# Cumulative Quiz Round 2 — Ch 3, 4, 7, 11 (Jul 24)

### Q1: Write-heavy event log with LSM-tree — compaction spikes causing latency. What's happening and mitigation?

**Answer:** As append-only SSTable files grow during bursts of record inserts, compaction is triggered in the background to merge SSTables and delete out-of-date records. Tune compaction params and compaction scheduling.

**Feedback:** Correct mechanism. The spikes are I/O contention — compaction reads/writes compete with foreground traffic. Underspecified on *what* spikes and *what* to tune. Concrete mitigations: (1) leveled compaction spreads work into smaller, more frequent merges vs size-tiered which batches into large expensive ones, (2) rate-limit compaction I/O (RocksDB `rate_limiter`), (3) dedicated compaction threads with CPU/IO budgets, (4) split hot/cold data — recent hot data stays in fewer SSTables, cold data compacts offline. "Tune compaction params" is the right idea but interviewers want to hear *which* params and *why*.

### Q2: Protobuf over gRPC — Service A adds new field, Service B hasn't deployed. What happens on the wire?

**Answer:** New fields should typically be marked optional or have defaults if required. Old code unmarshals using the previous schema w/o the new field, it simply drops the new field.

**Feedback:** Correct. Protobuf encodes by field number, not name — Service B's generated code has no entry for the new field number, so it skips those bytes. This is wire format design, not a convention. Optional + sensible defaults is the right call.

### Q3: Financial system — two transfers, concurrent reads must see consistent total. Which ACID properties and isolation level?

**Answer:** Account transfers involve mutating disjoint data (ie separate records - without a materializing conflicts table, Snapshot Serializable Isolation is required to detect potential cycles and rollback all but one transaction. All four acid properties are at play - atomicity: any transaction that could create a rw-rw cycle is rolled back as if it never ran, and any transaction that completes updates all necessary records as though it was a single operation; consistency: like a ledger, the transfer transaction maintains the validity and consistency of amounts that are added to one record and then deducted from another; Isolation: each transaction appears to run as though it is the only one executing at the moment - the activity of other transactions don't interfere due to provided snapshots and commit-time clean up; Durability: with the WAL, even during node failure the data for all accounts can be restored to their consistent state.

**Feedback:** Strong. Correctly identified SSI for disjoint records and all four properties. One nuance: "concurrent reads see a consistent total" targets **read skew** — two reads each see valid snapshots but the combination is inconsistent. Snapshot isolation alone doesn't prevent this. SSI catches it via dependency tracking. Calling it out by name ("read skew") would score points.

### Q4: Kafka topic, events out of order. What goes wrong windowing by processing time instead of event time?

**Answer:** Processing time is logged by the broker not the application - processing time has no semantic value to the application.

**Feedback:** Correct principle but undersells the concrete problem. Spell out the failure mode: a page view at 12:00:59 arriving at broker at 12:01:02 lands in the 12:01 window instead of 12:00. Results: (1) events in wrong windows, (2) aggregations incomplete (12:00 window missing data), (3) late-arriving events silently dropped or polluting next window. Fix is event-time windows + watermarks. Interviewers want failure mode, not just principle.

### Q5: B-tree vs LSM — read-heavy workload with occasional large batch writes (nightly ETL). Which wins and why?

**Answer:** Large batch writes, even at off-peak hours, can trigger heavy compaction for a LSM datastore - additionally, reads can be slower, especially for older data in older SSTables; a well indexed B-tree can be more performant for reads, even if they consist mostly of random I/O, large batch writes can trigger write amplification in the worst case - but should be negligible during off-peak hours

**Feedback:** Solid analysis. Correctly identified LSM compaction risk on batch writes, read degradation on older SSTables, and B-tree's read advantage. Write amplification during off-peak is correctly dismissed. But didn't explicitly state "B-tree wins." In an interview: say it clearly. B-tree wins — predictable O(log n) reads, ~3-4 disk reads for millions of keys. Write amplification from nightly ETL is acceptable when reads dominate.

### Q6: Exactly-once semantics in Kafka vs at-least-once with idempotent consumers — difference and when to pick which?

**Answer:** Kafka ensures during node failures that one consumer from a consumer-group can resume a stream, as the broker maintains offsets for each consumer/consumer group. Other vendors may retry un-acked events - in the event that a consumer had processed an event but failed before sending an ack, the consumer uses an idempotency key to determine if it has already processed the event - guaranteeing exactly-once semantics

**Feedback:** Correctly described Kafka's offset mechanism and idempotent consumers, but blended the two instead of separating them. Exactly-once is a system-level guarantee (broker commits offsets atomically with output writes). Idempotent consumers are a consumer-side workaround (dedup via key before processing). Pick exactly-once when broker supports it (Kafka). Pick at-least-once + idempotent when using queues without exactly-once guarantees (RabbitMQ, SQS) or when output isn't transactional with the broker.

### Q7: Migrating Protobuf V1 to V2 with a renamed field — compatibility concerns and safest rollout order?

**Answer:** Protobuf uses field numbers to match fields, a rename shouldn't be an issue

**Feedback:** Correct on wire format — field numbers handle it. But underspecified. Compatibility concerns beyond wire: (1) generated code changes variable names — compile-time breakage, (2) documentation/searches referencing old name stale, (3) downstream systems caching schema by name (not number) break. Safest rollout: (1) deploy dual-write (both old and new name), (2) update readers to tolerate either, (3) stop writing old name, (4) remove from schema. Never rename and deploy simultaneously.

### Q8: Stream processor handles 100K events/sec, downstream writes 20K/sec. What happens without intervention, and two strategies?

**Answer:** Stream processor buffers events - buffer overflows to disk. Downstream can apply back pressure

**Feedback:** Correct on buffer overflow and back pressure. But question asks for TWO strategies. Second options: (1) drop messages — accept data loss, (2) scale downstream — add consumers/partitions, (3) rate-limit ingestion — slow producer to match capacity. Back pressure is usually right tradeoff, but interviewers want to hear alternatives and when each is acceptable.

---

# DDIA Chapter 5 — Replication (Single Leader, up to pg 207)

## Quiz

### Q1: What's the difference between synchronous and asynchronous replication, and what's the practical compromise most systems use?

**Answer:** In synchronous rep, leader waits for followers to update. In async, leaders do not wait. Node failures can cause loss of data in async, and sync can affect availability. Semi-sync ensures that at least one follower is in sync with leader.

**Feedback:** Correct. Semi-sync is the sweet spot — one synchronous follower guarantees durability while keeping writes fast. If that follower dies, you promote another to synchronous (or tolerate the risk temporarily).

### Q2: A client writes to the leader, then immediately reads from a follower and doesn't see the write. What's this problem called, and what are two ways to solve it?

**Answer:** Read-your-writes. Either read from the leader for certain app-specific logical writes, or find a replica to read from that has the greatest logical timestamp among replicas for that user/write.

**Feedback:** First solution correct. Second is slightly off — standard solutions are: (1) read from leader after writes, (2) monotonic reads (stick to same replica), (3) causal timestamps (attach logical timestamp to writes, replicas report timestamp, client waits for replica to catch up before reading). "Greatest timestamp among replicas" isn't the pattern — you stick to one or ensure it's caught up to your write, not dynamically pick the most up-to-date.

### Q3: If a follower falls far behind (lagging), what's one safe way to bring it back up without serving stale reads?

**Answer:** Ask leader for logs between latest snapshot and now.

**Feedback:** Right idea but underspecified. Key nuance: replaying old logs is slow. Correct process: (1) leader takes a new snapshot, (2) ship snapshot to follower, (3) after follower loads snapshot, request logs from the point the snapshot was taken, (4) follower catches up by replaying just the recent logs. Snapshot gives a clean starting point instead of replaying thousands of old entries.

### Q4: Statement-based replication has a problem with non-deterministic functions like NOW() or RAND(). What's the alternative that avoids this?

**Answer:** Ship logical records in logs.

**Feedback:** Correct. Logical log records actual row-level changes (which row was inserted/updated/deleted, before/after values) rather than the SQL statement. NOW() and RAND() execute once on the leader — the resulting value is what gets shipped, so followers replay the exact same change.

### Q5: The leader crashes mid-write. How does the system choose a new leader, and what's the risk if the old leader comes back?

**Answer:** New leader is follower with latest/largest log offset. Old leader can still assume it's the leader — it would have to be quarantined.

**Feedback:** Correct on both counts. The risk is split-brain — two nodes both think they're leader, both accept writes, data diverges. Solution is fencing tokens: new leader gets a monotonically increasing token, old leader's token is invalidated. Any write from old leader with outdated token is rejected by followers/storage.

### Q6: What problem do consistent prefix reads solve, and why can't you just use timestamps?

**Answer:** Consistent prefix reads preserve the causal chain of writes. Timestamps provide a total ordering but don't provide info on which writes might be causally linked.

**Feedback:** Correct. A reader switching between followers at different lag levels might see answer 1 before question 1 — causal chain broken. Timestamps give a total order but don't tell you that write B was *caused by* write A. Solution: all replicas apply writes in the same sequence (sequence numbers, or ensuring writes from one leader are always applied in order on followers). This way, no matter which follower you read from, you see writes in the same relative order.

### Multi-Leader Replication (Jul 28)

**Q: What's the main disadvantage of multi-leader compared to single-leader?**

Conflicts.

**Q: What happens if different transactions use the same table but different rows on two leaders?**

Causal chain can be broken — write A on leader 1 happens before write B on leader 2, but B arrives at a replica before A, violating causality.

**Q: How does multi-leader ensure causal ordering?**

Version vectors track per-key causality between replicas. For cross-key causal ordering, the coordinator attaches a dependency list (or the client includes a vector timestamp of observed keys) so the receiving leader can enforce ordering.

**Q: What's the difference between conflict avoidance vs resolution, and when to pick which?**

Conflict avoidance pins a client/replica to a specific leader for all writes, trading availability/throughput for safety. Resolution is needed when writes aren't coordinated across leaders in favor of throughput — uses CRDTs, OT, or LWW to converge.

**Q: How do new leaders discover existing ones in multi-leader?**

Three mechanisms: (1) static config (restart to add), (2) gossip protocol (propagates membership in seconds), (3) external discovery (ZooKeeper/etcd). Cassandra uses gossip; Galera/InnoDB Cluster uses static seed list + gossip.

**Q: What happens on the write path when a leader fails in multi-leader?**

Clients route writes to another leader (app/load balancer failover). Other leaders continue normally. Dead leader's partners stop receiving replication from it. On recovery, the leader catches up via log replay or anti-entropy (gossip-based sync). Hinted handoff is Dynamo-style, not standard multi-leader — Cassandra uses it for temporary failures; MySQL Group Replication and PostgreSQL BDR just resync on reconnect.

---

## Chapter 5 Quiz (Jul 29, 2026)

**Score: 4.5/6**

**Q1 — Semi-sync replication timing** ❌
- The leader responds **after the sync follower acks**, not the async one.
- Semi-sync: exactly one sync follower guarantees durability on ≥1 replica; the rest are async.

**Q2 — Read-your-writes consistency** ✅ (2/3)
- Solutions: (1) read from leader, (2) pin user to same replica (monotonic reads), (3) **causal timestamps** — client attaches logical timestamp to writes, replica reports its latest, client waits until replica catches up.

**Q3 — Split-brain prevention** ✅
- Worst case: both leaders accept writes, data diverges.
- Prevention: **fencing tokens** — new leader gets monotonically increasing token, replicas reject stale-token writes.

**Q4 — Multi-leader conflict** ✅
- Two users, two DCs, last item: both writes succeed locally, conflict on replication.
- LWW silently drops one order. Single-leader trades write latency for zero conflicts.

**Q5 — Leaderless quorum (n=3, w=2, r=2)** ✅
- `w+r > n` guarantees ≥1 overlapping replica has fresh data.
- Resolution via version vectors + read repair pushes fresh data to stale replica.
- Anti-entropy (Merkle trees) catches replicas missed entirely.

**Q6 — Sloppy quorum** ✅
- Strict: requires designated n replicas, w acks from those. High consistency, lower availability during partition.
- Sloppy: any w healthy nodes. Higher availability, consistency restored via hinted handoff.

---

## DDIA Review (Jul 29, 2026)

### Ch 3 — Storage

**Q — B-tree page split** ✅
- Triggers on insert exceeding page fan-out. Expensive due to write amplification, half-empty pages, disk block movement.

**Q — Bloom filters in LSM** ❌
- User: "Never yield false positives." Correction: **Bloom filters can yield false positives but never false negatives** ("definitely no, maybe yes"). Tradeoff: memory for fewer unnecessary SSTable reads.

**Q — Statement-based replication problem** ✅
- Non-deterministic functions (NOW(), RAND()) give different results on replicas. Solution: **row-based replication** sends actual row values, not SQL statements.

**Q — B-tree vs LSM (read-heavy + nightly ETL)** ✅
- B-tree wins. Reads are O(log n) via direct page index. LSM compaction storms from batch writes can spill into peak hours.

### Ch 4 — Encoding

**Q — Backward vs forward compatibility** ❌
- User described forward issues (old code breaks on new fields) but labeled it "harder backward." Correction: **Forward compatibility** (old code reads new data) is harder — must ignore unknown fields gracefully.

**Q — Slow consumer in stream processor** ✅
- Two strategies: (1) **buffer/spill to disk** (Kafka's approach — log persists, consumer reads at own pace), (2) **backpressure** — signal producer to slow. Kafka does not drop messages.

**Q — Exactly-once semantics** ❌
- User mixed at-least-once + idempotent consumer with Kafka EOS. Kafka EOS = **idempotent producer** (producer ID + seq number for dedup) + **atomic transaction** (commit produce + offset update atomically).

### Ch 5 — Replication

**Q — Read-your-writes consistency** ✅
- 3 solutions: (1) read from leader, (2) monotonic reads (same replica), (3) causal timestamps.

**Q — Consistent prefix reads** ❌
- User missed clock skew. Timestamps don't work because clocks drift — timestamp ordering ≠ causal ordering. Need hybrid logical clocks or version vectors.

**Q — Multi-leader disadvantage** ✅
- Write conflicts and resolution complexity. Async replication between leaders also causes read-your-writes and causal consistency issues.

**Q — Multi-leader causal ordering across keys** ❌
- User said "version vectors." Correction: version vectors handle **per-key** causality. Cross-key causal ordering needs **dependency lists** (vector timestamp of all observed keys).

**Q — New leader discovery** ❌ (partial)
- User said "coordinator tracks leaders." Three mechanisms: **static config**, **gossip protocol**, **external discovery** (ZK/etcd).

### Ch 5 Review — Jul 30, 2026

**Q — Semi-sync replication** ✅
- User: "the semi-sync follower — of all the followers, typically all but one are async"
- Mine: One synchronous follower acks before leader responds. Remaining followers are async. Ensures durability on ≥1 replica.

**Q — Read-your-writes (third solution)** ✅
- User: "read from leader for a given user or a given op"
- Mine: Causal timestamps — client tracks logical timestamp of its last write, replica reports its latest, client waits until replica catches up.

**Q — Split-brain prevention via fencing token** ✅
- User: "replicas get back a token that's greater than what the rogue leader might have"
- Mine: Coordination service issues monotonically increasing token to new leader. Storage layer rejects writes from old leader bearing a stale token.

**Q — Multi-leader LWW tradeoff** ✅
- User: "no app code conflict resolution needed, but intermediate writes get lost"
- Mine: LWW guarantees convergence without app logic, but silently drops concurrent writes that lose the timestamp comparison.

**Q — Leaderless quorum (w+r>n)** ✅
- User: "if 2/3 are written to, reading from 2 nodes guarantees at least one of the three nodes will have fresh data. if one of the two nodes delivers stale data, read repair and anti-entropy fix it"
- Mine: n=3,w=2,r=2 ensures w+r>n guarantees overlap on every read. Stale reads fixed by read repair (on read) and anti-entropy (background).

**Q — Sloppy vs strict quorum** ✅
- User: "sloppy quorum doesn't require majority"
- Mine: Strict = designated n replicas must respond. Sloppy = any w healthy nodes. Sloppy trades consistency for availability; hinted handoff replays writes to designated nodes.

### Ch 7 — Transactions

**Q — SSI vs SI** ✅
- SSI detects **rw-dependency cycles** (write skew). SI only detects conflicts on same row via MVCC.

**Q — Write skew scenario** ✅
- Two doctors check on-call roster, each takes themselves off. Each writes different row — SI misses the cross-row rw-dependency.

### Ch 11 — Stream Processing

**Q — Windows in stream processing** ✅
- Four types: **tumbling** (fixed, non-overlapping), **sliding** (overlapping, size + slide), **session** (gap-based), **fixed** (= tumbling). Key: event time vs processing time — processing time gives wrong results for late events.

### Ch 7 — Multi-Leader Replication (Review Jul 31, 2026)

**Q — How do new leaders discover existing ones?** ✅
- User: "not sure if Kafka uses Zookeeper? gossip protocol or external discovery for dynamic"
- Mine: Three mechanisms: **static config** (restart to add), **gossip** (Cassandra, propagates in seconds, no central dependency), **external discovery** (ZooKeeper/etcd; Galera/InnoDB Cluster). Dynamic clusters → gossip or external discovery. Kafka: used ZooKeeper but is single-leader-per-partition, and dropped ZK in KRaft (3.3+) for its own Raft metadata quorum (neither gossip nor external). ZooKeeper: HDFS/HBase/Pulsar/Solr/Trino; etcd: Kubernetes/TiDB; third option: Consul. Gossip = eventually-consistent membership (AP-friendly), external discovery = strong membership (needs coordination). This is the same CAP-flavored tradeoff.

**Q — How does multi-leader ensure causal ordering?** ✅
- User: "version vectors per key, prefixes across keys"
- Mine: Version vectors capture per-key causality (each replica counts its own writes). Across keys, no single node sees all writes, so the client/coordinator attaches a **prefix** — a vector of versions of every observed key — and a node receiving a write can detect "missing a write this depends on" and refuse/queue rather than reorder.

**Q — Same table, different rows, different leaders?** ✅
- User: "rw-rw cycles can go undetected"
- Mine: **Write skew**: two transactions write *different* rows (no same-key WW conflict → async conflict detection sees nothing), yet each read a row the other wrote → a WR/RW **cycle** in the dependency graph. No single node sees both writes, so the nonserializable schedule goes undetected. Conflict detection only flags same-key WW; RW/cycle edges are invisible. (Two-doctors-on-call example.)

**Q — Write path when a leader fails?** ✅
- User: "alternate leader(s) write atomically to their tables/WALs and propagate — replacement leader elected"
- Mine: No election in multi-leader — clients/app/LB re-route to an available leader (failover = client retargeting). Peers keep accepting writes; the dead leader just stops receiving their streams. On recovery the failed leader **catches up** by replaying missed writes. Election is single-leader/Raft vocabulary (ZK/etcd, Galera primary-component quorum). "Write atomically to WAL and propagate" is just how any leader commits + ships its stream.

**Q — Conflict avoidance vs resolution?** ✅
- User: "avoidance only achievable in single-leader, resolution needed in multi-leader and leaderless; single for consistency, others for availability"
- Mine: Avoidance is **automatic** in single-leader (one write point) but also *possible* in multi-leader via **pinning** a client/replica to one leader — trades availability/throughput for safety; if the pinned leader fails, writes fail or conflicts return. Leaderless has no avoidance option (any replica accepts writes) → resolution mandatory. CAP summary: single-leader → consistency; multi/leaderless → availability.

**Q — Main disadvantage of multi-leader?** ✅
- User: "causal, read-your-writes, and monotonic conflicts"
- Mine: Conflicts is the headline; concrete failures are the three guarantees that weaken: **read-your-writes** (read lands on a replica that hasn't received your write), **monotonic reads** (later read hits older replica), **causal ordering** (dependent writes arrive out of order). Avoided by pinning a user's reads+writes to one leader — loops back to the avoidance Q.

## Review — 2026-08-01 (5 items)

**Q — Ch 4: JSON schema evolution (add field V1→V2, old code/new code)?** ✅
- User: "old code can read new schema in both as long as JSON configured to ignore unknowns and Protobuf fields not reordered; renames/type changes break both; new code can read old schema in both, but Protobuf needs defaults for new fields"
- Mine: Field *numbers* are Protobuf's versioning contract — unknown numbers skipped on read, missing read as zero/default. JSON has no wire schema; depends on tool config (Jackson can ignore, others error). Protobuf required is writer-side only; renames/type changes break both since field numbers / JSON keys are immutable contracts.

**Q — Ch 4: Adding a required field in Protobuf V2 (asymmetry)?** ✅
- User: "favors backward compatibility — old code should keep working but break when writing new data"
- Mine: `required` is enforced only by the **writer** at encode time (field must be set before serializing, else marshal fails). Readers never enforce it — a new reader can't reject old data. Old reader + new msg: skips unknown field. Old writer + new schema: fails loudly. This unsafe asymmetry is why `required` was **removed in Protobuf 3**; use optional + defaults, enforce in app.

**Q — Ch 5: How does a lagging follower catch up?** ✅
- User: "requests that the leader slow down" ❌
- Mine: Leader doesn't pace itself for followers. Follower **fetches missing log entries** between its last applied position and the leader's current position (or restores from snapshot, then replays the post-snapshot log). Async replication: follower pulls the gap on its own.

**Q — Ch 9: Conflict avoidance vs resolution?** ✅
- User: "avoidance only possible in single-leader; multi-leader should resolve read-your-writes, causal, monotonic conflicts" — partially mixed
- Mine: Avoidance = **pinning** a client's writes to one leader (works in multi-leader too; trades availability/throughput for safety). Resolution = LWW / merge (CRDT) / app-level conflict handling when writes hit different leaders. read-your-writes/monotonic/causal are Ch 5 *replication-lag consistency* guarantees — a separate axis from Ch 9 write-write conflicts.

**Q — Ch 5: causal timestamps for RYW on followers?** ✅
- User: "reads from the leader in the meantime, or sends a timestamp with every read"
- Mine: Read-from-leader guarantees RYW but kills read scaling. Scalable version: client tags each write with a **logical timestamp**, remembers last-write ts, sends it with reads; follower compares its last applied ts and **waits/retries until caught up** (or errors → retry another replica).
