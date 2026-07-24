package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Claims struct {
	Sub  string `json:"sub"`
	Exp  int64  `json:"exp"`
	Iat  int64  `json:"iat"`
	Nbf  int64  `json:"nbf"`
	Role string `json:"role"`
}

func IssueJWT(key []byte, sub, role string, ttl time.Duration) (string, error) {
	if len(key) < 32 {
		return "", errors.New("key must be at least 32 bytes for HS256 security")
	}
	h := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	now := time.Now().Unix()
	c, err := json.Marshal(Claims{
		Sub:  sub,
		Role: role,
		Iat:  now,
		Nbf:  now,
		Exp:  now + int64(ttl.Seconds()),
	})
	if err != nil {
		return "", err
	}
	p := base64.RawURLEncoding.EncodeToString(c)
	sigInput := fmt.Sprintf("%s.%s", h, p)

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(sigInput))
	sig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s.%s", sigInput, sig), nil
}

func VerifyJWT(key []byte, token string) (*Claims, error) {
	if len(key) < 32 {
		return nil, errors.New("key must be at least 32 bytes")
	}
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid jwt format")
	}

	hBuf, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, errors.New("invalid header encoding")
	}

	var header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	if err := json.Unmarshal(hBuf, &header); err != nil {
		return nil, errors.New("invalid header json")
	}

	if header.Alg != "HS256" {
		return nil, errors.New("unsupported or algorithm confusion attack detected")
	}

	sigInput := fmt.Sprintf("%s.%s", parts[0], parts[1])
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(sigInput))
	expectedSig := mac.Sum(nil)

	actualSig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, errors.New("invalid signature encoding")
	}

	if subtle.ConstantTimeCompare(expectedSig, actualSig) != 1 {
		return nil, errors.New("jwt signature mismatch")
	}

	cBuf, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, errors.New("invalid payload encoding")
	}

	var c Claims
	if err := json.Unmarshal(cBuf, &c); err != nil {
		return nil, errors.New("corrupt payload json")
	}

	now := time.Now().Unix()
	if c.Nbf > 0 && now < c.Nbf {
		return nil, errors.New("jwt token not valid yet")
	}
	if now > c.Exp {
		return nil, errors.New("jwt token expired")
	}

	return &c, nil
}
