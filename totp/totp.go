package totp

import (
	"crypto/hmac"
	"crypto/sha1"
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
