package totp_test

import (
	"testing"
	"time"

	"github.com/alnvdl/totp/totp"
)

func assert(t *testing.T, gotCode, wantCode string) {
	if wantCode != gotCode {
		t.Fatalf("unexpected TOTP: want %s got %s", wantCode, gotCode)
	}
}

func TestTOTP(t *testing.T) {
	assert(t, totp.TOTP([]byte("abcde"), time.Unix(1673803904, 0)), "589712")
	assert(t, totp.TOTP([]byte("vwxyz"), time.Unix(1673804234, 0)), "036054")
}
