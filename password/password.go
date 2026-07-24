package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/argon2"
)

type Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var DefaultConfig = Config{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func Hash(pass string) (string, error) {
	salt := make([]byte, DefaultConfig.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(pass), salt, DefaultConfig.Iterations, DefaultConfig.Memory, DefaultConfig.Parallelism, DefaultConfig.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return b64Salt + "." + b64Hash, nil
}

func Verify(pass, encodedHash string) (bool, error) {
	var saltB64, hashB64 string
	for i := 0; i < len(encodedHash); i++ {
		if encodedHash[i] == '.' {
			saltB64 = encodedHash[:i]
			hashB64 = encodedHash[i+1:]
			break
		}
	}
	if saltB64 == "" || hashB64 == "" {
		return false, errors.New("invalid format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(saltB64)
	if err != nil {
		return false, err
	}
	expectedHash, err := base64.RawStdEncoding.DecodeString(hashB64)
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey([]byte(pass), salt, DefaultConfig.Iterations, DefaultConfig.Memory, DefaultConfig.Parallelism, uint32(len(expectedHash)))
	return subtle.ConstantTimeCompare(hash, expectedHash) == 1, nil
}
