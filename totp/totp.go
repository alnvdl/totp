// Package totp contains the function that generates a TOTP code.
package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

// TOTP generates a TOTP for the given key at time t, in the manner commonly
// used by most websites (SHA1-based 6-digit code padded with 0s).
// This implementation is based on RFC 4226 and RFC 6238.
func TOTP(key []byte, t time.Time) string {
	counter := uint64(math.Floor(float64(t.Unix()) / float64(30)))
	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, counter)

	h := hmac.New(sha1.New, key)
	h.Write(counterBytes)
	hmacResult := h.Sum(nil)

	offset := hmacResult[19] & 0xf
	var binCode uint32 = (uint32(hmacResult[offset]&0x7f))<<24 |
		uint32((hmacResult[offset+1]&0xff))<<16 |
		uint32((hmacResult[offset+2]&0xff))<<8 |
		uint32((hmacResult[offset+3] & 0xff))

	return fmt.Sprintf("%06d", int(binCode%1000000))
}
