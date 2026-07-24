import hmac
import hmac as hmac_mod
import hashlib
import os
import re
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

def constant_time_compare(val1: str, val2: str) -> bool:
    return hmac_mod.compare_digest(val1.encode('utf-8'), val2.encode('utf-8'))

def sanitize_input(val: str, max_len: int = 256) -> str:
    if not isinstance(val, str):
        return ""
    val = val.strip().replace('\x00', '').replace('\r', '')
    return val[:max_len]

def encrypt_aes256_gcm(key: bytes, plaintext: bytes) -> bytes:
    if len(key) != 32:
        raise ValueError("Key must be 32 bytes")
    aesgcm = AESGCM(key)
    nonce = os.urandom(12)
    return nonce + aesgcm.encrypt(nonce, plaintext, None)

def decrypt_aes256_gcm(key: bytes, ciphertext: bytes) -> bytes:
    if len(key) != 32:
        raise ValueError("Key must be 32 bytes")
    if len(ciphertext) < 28:
        raise ValueError("Ciphertext too short")
    aesgcm = AESGCM(key)
    nonce = ciphertext[:12]
    ct = ciphertext[12:]
    return aesgcm.decrypt(nonce, ct, None)

def is_valid_fen(fen: str) -> bool:
    pat = r'^([rnbqkbnrRNBQKBNR1-8]{1,8}/){7}[rnbqkbnrRNBQKBNR1-8]{1,8}\s+[wb]\s+([KQkqA-Ha-h]{1,4}|-)\s+([a-h][36]|-)\s+\d+\s+\d+$'
    return bool(re.match(pat, fen.strip()))
