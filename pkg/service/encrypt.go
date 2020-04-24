package service

import (
	"encoding/base64"
	"net/http"

	"github.com/zhiqiangxu/ncrypt/pkg/crypto"
	"github.com/zhiqiangxu/ncrypt/pkg/instance"
	"github.com/zhiqiangxu/ncrypt/pkg/io"
	"github.com/zhiqiangxu/util"
)

// Encrypt impl
func (s *Service) Encrypt(input io.EncryptInput) (output io.EncryptOutput) {

	cipherBytes, err := crypto.AESEncrypt(util.Slice(input.PlainText), instance.Secret().GetSecret())
	if err != nil {
		output.Code = http.StatusInternalServerError
		output.Msg = err.Error()
		return
	}

	output.CipherText = base64.StdEncoding.EncodeToString(cipherBytes)
	return
}
