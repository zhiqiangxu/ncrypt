package service

import (
	"encoding/base64"
	"net/http"

	"github.com/zhiqiangxu/ncrypt/pkg/crypto"
	"github.com/zhiqiangxu/ncrypt/pkg/instance"
	"github.com/zhiqiangxu/ncrypt/pkg/io"
	"github.com/zhiqiangxu/util"
)

// Decrypt impl
func (s *Service) Decrypt(input io.DecryptInput) (output io.DecryptOutput) {
	base64Decoded, err := base64.StdEncoding.DecodeString(input.CipherText)
	if err != nil {
		output.Code = http.StatusInternalServerError
		output.Msg = err.Error()
		return
	}

	plainBytes, err := crypto.AESDecrypt(base64Decoded, instance.Secret().GetSecret())
	if err != nil {
		output.Code = http.StatusInternalServerError
		output.Msg = err.Error()
		return
	}

	output.PlainText = util.String(plainBytes)

	return
}
