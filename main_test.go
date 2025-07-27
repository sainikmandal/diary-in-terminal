package main

import (
	"testing"

	"github.com/sainikmandal/diary-in-terminal/encryption"
)

func TestEncryptionCycle(t *testing.T) {
	key := []byte("12345678901234567890123456789012") // 32-byte test key
	text := "Hello secure world!"

	enc, err := encryption.Encrypt(text, key)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	dec, err := encryption.Decrypt(enc, key)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if dec != text {
		t.Fatalf("Expected %s, got %s", text, dec)
	}
}
