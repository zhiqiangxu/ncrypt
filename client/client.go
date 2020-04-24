package client

import (
	"encoding/json"
	"fmt"

	"github.com/zhiqiangxu/ncrypt/pkg/forward"
	"github.com/zhiqiangxu/ncrypt/pkg/io"
)

// Client for ncrypt
type Client struct {
	host string
}

// NewClient is ctor for Client
func NewClient(host string) *Client {
	return &Client{host: host}
}

// Encrypt plainText
func (c *Client) Encrypt(plainText string) (cipherText string, err error) {
	input := io.EncryptInput{PlainText: plainText}
	reqBytes, err := json.Marshal(input)
	if err != nil {
		return
	}
	url := "http://" + c.host + "/ncrypt/encrypt"
	code, _, bodyBytes, err := forward.PostJSONRequest(url, reqBytes)
	if err != nil {
		return
	}

	if code != 200 {
		err = fmt.Errorf("code:%d bodyBytes:%s", code, bodyBytes)
		return
	}

	var output io.EncryptOutput
	err = json.Unmarshal(bodyBytes, &output)
	if err != nil {
		return
	}

	cipherText = output.CipherText
	return
}

// Decrypt cipherText
func (c *Client) Decrypt(cipherText string) (plainText string, err error) {
	input := io.DecryptInput{CipherText: cipherText}
	reqBytes, err := json.Marshal(input)
	if err != nil {
		return
	}
	url := "http://" + c.host + "/ncrypt/decrypt"
	code, _, bodyBytes, err := forward.PostJSONRequest(url, reqBytes)
	if err != nil {
		return
	}

	if code != 200 {
		err = fmt.Errorf("code:%d bodyBytes:%s", code, bodyBytes)
		return
	}

	var output io.DecryptOutput
	err = json.Unmarshal(bodyBytes, &output)
	if err != nil {
		return
	}
	plainText = output.PlainText
	return
}
