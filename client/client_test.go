package client

import (
	"testing"

	"gotest.tools/assert"
)

func TestClient(t *testing.T) {
	c := NewClient("localhost:9999")

	plainText := "plain text"
	cipherText, err := c.Encrypt(plainText)
	assert.Assert(t, err == nil)

	decipheredText, err := c.Decrypt(cipherText)
	assert.Assert(t, err == nil && decipheredText == plainText, decipheredText)
}
