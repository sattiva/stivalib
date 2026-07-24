# sativas-lib (v4.0 Multi-Language Defensive Security Core)

Unified defensive security, cryptography, AST injection prevention, and execution sandboxing library built for Go, JavaScript/Node.js, Python, and WebAssembly.

## Package Architecture & Modular Directories

1. **`crypto/` (Go)**: Constant-time authentication, AES-256-GCM authenticated encryption, HMAC verification.
2. **`auth/` (Go)**: Hardened HS256 JWT issuer/verifier enforcing constant-time signature validation and expiry checks.
3. **`net/` (Go)**: Leak-proof HTTP client stripping `X-Forwarded-For`, `Via`, `X-Real-IP`, with randomized UA headers.
4. **`log/` (Go)**: Zero-allocation JSON logger with asynchronous Discord alert dispatching for error states.
5. **`validate/` (Go)**: Strict input validation, FEN/move regex matchers, template string escaping.
6. **`ratelimit/` (Go)**: Thread-safe, memory-bounded token bucket IP rate limiter with automatic TTL cleanup.
7. **`audit/` (Go)**: Cryptographically tamper-proof append-only SHA-256 hash-chained audit log chain.
8. **`sys/` (Go)**: Secure entropy byte/int generator and memory buffer zeroing (`ZeroBuffer`).
9. **`sandbox/` (Go)**: Context-bounded command execution engine with strict timeouts and exit status tracking.
10. **`js/` (Node.js)**: `index.js` implementing timing-safe string comparison, AES-256-GCM, and JS template string escaping.
11. **`py/` (Python)**: `sativas_lib.py` supplying `hmac.compare_digest`, AES-256-GCM AEAD, and input sanitizers.
12. **`wasm/` (WebAssembly)**: High-performance Go WASM bridge exporting constant-time string comparators to browser environments.

## Integration Examples

### Go
```go
import (
    "sativas-lib/auth"
    "sativas-lib/crypto"
    "sativas-lib/sandbox"
)

token, _ := auth.IssueJWT([]byte("secret_key_32_bytes_long_123456"), "user_1", "admin", 1 * time.Hour)
```

### Node.js
```javascript
const { constantTimeCompare, encryptAES256GCM } = require('./js');
```

### Python
```python
from py.sativas_lib import constant_time_compare, encrypt_aes256_gcm
```
