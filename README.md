# sativlib (v5.5 Advanced Defensive Security Core)

High-performance, multi-language security, cryptography, and systems utility library designed to replace bloated external dependencies.

## Comprehensive Module Map

* **`crypto/`**: AES-256-GCM authenticated encryption, HKDF key derivation (RFC 5869), HMAC-SHA256, constant-time comparisons.
* **`auth/`**: HS256 JWT generation/validation, bitmask scope authorization (`ScopeRead`, `ScopeWrite`, `ScopeAdmin`, etc.).
* **`password/`**: Argon2id password hashing & constant-time key verification (`golang.org/x/crypto/argon2`).
* **`totp/`**: RFC 6238 Time-Based One-Time Password (2FA) engine.
* **`net/`**: Stealth HTTP transport removing leak headers (`X-Forwarded-For`, `Via`, `X-Real-IP`) with randomized UA injection.
* **`log/`**: Zero-allocation JSON logger with asynchronous Discord Webhook error dispatch.
* **`ratelimit/`**: Memory-bounded Token Bucket IP rate limiter with background TTL cleanup.
* **`cache/`**: Thread-safe in-memory cache engine with background Janitor worker.
* **`circuit/`**: Circuit Breaker pattern preventing cascade outages.
* **`retry/`**: Exponential backoff retry engine with jitter calculation.
* **`env/`**: Type-safe environment variable parsing (`Get`, `GetInt`, `GetBool`).
* **`audit/`**: Cryptographically tamper-proof SHA-256 append-only hash chain.
* **`sandbox/`**: Context-bounded command execution engine tracking timeouts and exit codes.
* **`sys/`**: Resource watchdog (`runtime.MemStats`/`runtime.NumGoroutine`), secure random entropy, and memory zeroing (`ZeroBuffer`).
* **`pool/`**: Zero-copy recycled `BytePool` (`sync.Pool`) for high-throughput I/O.
* **`validate/`**: Strict input sanitization, FEN/move matchers, template string escaping.
* **`js/`**: Node.js module providing `constantTimeCompare`, `encryptAES256GCM`, `decryptAES256GCM`, and template escaping.
* **`py/`**: Python module supplying `compare_digest`, `AESGCM`, and input sanitization.
* **`wasm/`**: WebAssembly bridge exporting timing-safe comparators to browser runtimes.

## Usage

```go
import (
    "sativlib/auth"
    "sativlib/crypto"
    "sativlib/sys"
)
```
