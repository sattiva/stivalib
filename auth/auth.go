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
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
	Iat int64  `json:"iat"`
	Role string `json:"role"`
}

func IssueJWT(key []byte, sub, role string, ttl time.Duration) (string, error) {
	h := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	now := time.Now().Unix()
	c, err := json.Marshal(Claims{
		Sub: sub,
		Role: role,
		Iat: now,
		Exp: now + int64(ttl.Seconds()),
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
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid jwt format")
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

	if time.Now().Unix() > c.Exp {
		return nil, errors.New("jwt token expired")
	}

	return &c, nil
}
