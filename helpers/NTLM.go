package helpers

import (
	"encoding/binary"
	"encoding/hex"
	"unicode/utf16"
	"github.com/sergeymakinen/go-crypt/nthash"
)

func NTHASH(plainText string) string {
	// UTF-8 to UTF-16 LE
	a := utf16.Encode([]rune(plainText))
	b := make([]byte, len(a)*2)
	for i, r := range a {
		binary.LittleEndian.PutUint16(b[i*2:], r)
	}

	// Generate NTHASH
	key, error_hash := nthash.Key(b)

	// Check if error_hash is not nil
	if error_hash != nil {
		// Return error
		return "something_error"
	}

	return hex.EncodeToString(key)
}