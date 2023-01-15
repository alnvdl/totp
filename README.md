# totp

This library generates a TOTP for the given key at time t, in the manner
commonly used by most websites (SHA1-based 6-digit code padded with 0s). The
implementation is based on [RFC 4226](https://www.ietf.org/rfc/rfc4226.txt) and
[RFC 6238](https://www.ietf.org/rfc/rfc6238.txt).

There's only one function, which generates the TOTP code for `key` at time `t`:
```go
func TOTP(key []byte, t time.Time) string
```

I mainly wrote this because I wanted to see a bit more in depth how TOTP works.
There are much more complete libraries out there.
