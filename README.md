# sativas-lib (v5.0 Enterprise Defensive Core)

Comprehensive, multi-language security & utility library removing 95%+ of external dependency overhead.

## Added Core Modules Overview (30+ Total Packages)

1. **`crypto/`**: AES-256-GCM authenticated encryption, HMAC-SHA256, constant-time comparisons.
2. **`password/`**: Argon2id password hashing & verification (`golang.org/x/crypto/argon2`).
3. **`totp/`**: RFC 6238 Time-Based One-Time Password (2FA) generation engine.
4. **`auth/`**: HS256 JWT generation and validation with constant-time verification.
5. **`ratelimit/`**: Bounded Token Bucket IP rate-limiter with automatic TTL memory cleanup.
6. **`cache/`**: Thread-safe in-memory cache engine with background Janitor TTL worker.
7. **`circuit/`**: Circuit Breaker pattern preventing cascade failures across external endpoints.
8. **`retry/`**: Exponential backoff retry engine with jitter calculation.
9. **`env/`**: Type-safe environment variable parsing (`Get`, `GetInt`, `GetBool`).
10. **`audit/`**: Cryptographically tamper-proof hash-chained append-only audit log chain.
11. **`sandbox/`**: Context-bounded command execution engine tracking process timeouts and exit codes.
12. **`js/`**: Node.js module providing `constantTimeCompare`, `encryptAES256GCM`, `decryptAES256GCM`, and JS template string escaping.
13. **`py/`**: Python module providing `compare_digest`, `AESGCM`, `sanitize_input`, and chess FEN validation.
14. **`wasm/`**: WebAssembly bridge for browser environments exporting timing-safe comparators.
