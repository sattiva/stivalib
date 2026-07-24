package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"hash"
	"io"
)

func Equal(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}

func EqualString(a, b string) bool {
	return Equal([]byte(a), []byte(b))
}

func EncryptGCM(key, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes for AES-256")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func DecryptGCM(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes for AES-256")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ns := gcm.NonceSize()
	if len(ciphertext) < ns+gcm.Overhead() {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ct := ciphertext[:ns], ciphertext[ns:]
	return gcm.Open(nil, nonce, ct, nil)
}

func HMACSign(key, msg []byte, h func() hash.Hash) []byte {
	mac := hmac.New(h, key)
	mac.Write(msg)
	return mac.Sum(nil)
}

func HMACVerify(key, msg, expectedSig []byte, h func() hash.Hash) bool {
	actualSig := HMACSign(key, msg, h)
	return subtle.ConstantTimeCompare(actualSig, expectedSig) == 1
}
