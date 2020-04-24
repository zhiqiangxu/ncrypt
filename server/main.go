package main

import (
	"fmt"
	"os"

	"github.com/zhiqiangxu/ncrypt/pkg/instance"
	"github.com/zhiqiangxu/ncrypt/pkg/rest"
	"go.uber.org/zap"
)

const (
	addr = ":9999"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("at least two parties!")
		return
	}
	err := instance.Secret().DecideSecret(os.Args[1:])
	if err != nil {
		instance.Logger().Fatal("DecideSecret", zap.Error(err))
	}

	s := rest.NewServer()
	s.Start(addr)
}
