package io

// EncryptInput for input
type EncryptInput struct {
	PlainText string `json:"plain_text"`
}

// EncryptOutput for output
type EncryptOutput struct {
	BaseResp
	CipherText string `json:"cipher_text"`
}
