# sativas-lib

Core utility modules library for Sativa's tooling ecosystem.

## Modules

* **`crypto`**: Constant-time token comparisons (`Equal`), AES-256-GCM encryption/decryption, HMAC signing.
* **`net`**: Stealth HTTP Transport removing leak headers (`X-Forwarded-For`, `Via`) and rotating User-Agents.
* **`log`**: Zero-allocation JSON logger with built-in Discord webhook error dispatch.
* **`validate`**: Strict input validation (FEN chess strings, move syntax, alphanumeric IDs) and JS template string escaping.

## Import

```go
import (
    "sativas-lib/crypto"
    "sativas-lib/net"
    "sativas-lib/log"
    "sativas-lib/validate"
)
```
