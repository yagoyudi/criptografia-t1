package aes128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	key := []byte{
		0x2b, 0x7e, 0x15, 0x16,
		0x28, 0xae, 0xd2, 0xa6,
		0xab, 0xf7, 0x97, 0x75,
		0x45, 0x21, 0x48, 0x8d,
	}
	aes := NewAES(key)

	plaintext := []byte("This test will be using 2 blocks")
	got, err := aes.Encrypt(plaintext)
	want := []byte{0xc, 0x1, 0x5, 0xbb, 0x70, 0xfe, 0x69, 0x94, 0x61, 0x80,
		0x17, 0xb7, 0x21, 0xee, 0xa4, 0x81, 0x57, 0x7b, 0x14, 0xcd, 0xf2, 0xc4, 0x42,
		0x50, 0x2e, 0x55, 0xba, 0xc3, 0xd7, 0x83, 0xe4, 0x9d}
	assert.NoError(t, err)
	assert.Equal(t, got, want)

}
