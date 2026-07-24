package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

func Generate(secret []byte, t time.Time) string {
	ts := uint64(t.Unix() / 30)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, ts)

	mac := hmac.New(sha1.New, secret)
	mac.Write(buf)
	h := mac.Sum(nil)

	offset := h[len(h)-1] & 0x0f
	code := (int32(h[offset]&0x7f) << 24) |
		(int32(h[offset+1]&0xff) << 16) |
		(int32(h[offset+2]&0xff) << 8) |
		(int32(h[offset+3] & 0xff))

	otp := code % int32(math.Pow10(6))
	return fmt.Sprintf("%06d", otp)
}

func Verify(secret []byte, code string, t time.Time, skewWindow int) bool {
	if len(code) != 6 {
		return false
	}

	for i := -skewWindow; i <= skewWindow; i++ {
		checkTime := t.Add(time.Duration(i*30) * time.Second)
		gen := Generate(secret, checkTime)
		if subtle.ConstantTimeCompare([]byte(gen), []byte(code)) == 1 {
			return true
		}
	}
	return false
}
