# sativlib (v5.5 Enterprise Security Engine)

## Features & Capabilities

- **AES-256-GCM Encryption (`crypto`)**: Authenticated symmetric encryption and decryption with nonce validation (`EncryptGCM`, `DecryptGCM`).
- **HKDF Key Derivation (`crypto`)**: RFC 5869 Extract-and-Expand Key Derivation Function (`HKDFDeriveKey`).
- **Constant-Time Verification (`crypto`)**: Timing-attack resistant byte and string comparison (`Equal`, `EqualString`, `HMACVerify`).
- **Hardened JWT Authentication (`auth`)**: HS256 JWT issuer & verifier (`IssueJWT`, `VerifyJWT`) enforcing minimum key lengths, `nbf`/`exp` claims, and algorithm confusion attack prevention.
- **Bitmask Authorization (`auth`)**: Zero-lookup scope authorization (`ScopeRead`, `ScopeWrite`, `ScopeAdmin`, `ScopeDeploy`, `ScopeExecute`).
- **Argon2id Password Hashing (`password`)**: Modern memory-hard password hashing & constant-time key verification (`Hash`, `Verify`).
- **RFC 6238 TOTP 2FA Engine (`totp`)**: Time-Based One-Time Password generator & verifier with clock skew tolerance (`Generate`, `Verify`).
- **Stealth HTTP Transport (`net`)**: Leak-proof HTTP client stripping `X-Forwarded-For`, `Via`, and `X-Real-IP` while injecting dynamic browser User-Agents.
- **Bounded Token Bucket Rate Limiter (`ratelimit`)**: Memory-bounded IP rate limiter with automatic background TTL cleanup.
- **In-Memory Cache Engine (`cache`)**: Thread-safe cache with background Janitor worker clearing expired keys.
- **Circuit Breaker (`circuit`)**: Microservice fault-tolerance handling `Closed`, `HalfOpen`, and `Open` states.
- **Exponential Backoff Retry (`retry`)**: Resilient operation retries with randomized jitter.
- **Type-Safe Env Parser (`env`)**: Environment configuration extractor (`Get`, `GetInt`, `GetBool`).
- **Tamper-Proof Audit Chain (`audit`)**: SHA-256 append-only cryptographic hash chain for immutable event logging.
- **Command Execution Sandbox (`sandbox`)**: Context-bounded process execution engine with timeout enforcement.
- **Resource Watchdog (`sys`)**: Memory (`AllocBytes`) and goroutine monitoring watchdog with self-healing alert callbacks.
- **Zero-Copy BytePool (`pool`)**: Reusable `sync.Pool` buffer recycler eliminating GC latency spikes.
- **Input Sanitizer (`validate`)**: String sanitization, chess FEN/move matchers, and JS template string escaping.
- **Multi-Language Bridge (`js`, `py`, `wasm`)**: Built-in runtime bindings for Node.js, Python, and WebAssembly.

## Installation

```bash
go get github.com/sativac/sativlib
```
