package io

// DecryptInput for input
type DecryptInput struct {
	CipherText string `json:"cipher_text"`
}

// DecryptOutput for output
type DecryptOutput struct {
	BaseResp
	PlainText string `json:"plain_text"`
}
