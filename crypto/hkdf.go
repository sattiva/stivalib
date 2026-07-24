package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
)

func HKDFExpand(prk, info []byte, outLen int) ([]byte, error) {
	mac := hmac.New(sha256.New, prk)
	hashLen := mac.Size()
	if outLen > 255*hashLen {
		return nil, errors.New("requested length too long")
	}

	var okm []byte
	var T []byte
	counter := byte(1)

	for len(okm) < outLen {
		mac.Reset()
		mac.Write(T)
		mac.Write(info)
		mac.Write([]byte{counter})
		T = mac.Sum(nil)
		okm = append(okm, T...)
		counter++
	}

	return okm[:outLen], nil
}

func HKDFExtract(secret, salt []byte) []byte {
	if len(salt) == 0 {
		salt = make([]byte, sha256.Size)
	}
	mac := hmac.New(sha256.New, salt)
	mac.Write(secret)
	return mac.Sum(nil)
}

func HKDFDeriveKey(secret, salt, info []byte, keyLen int) ([]byte, error) {
	prk := HKDFExtract(secret, salt)
	return HKDFExpand(prk, info, keyLen)
}
