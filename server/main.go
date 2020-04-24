package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zhiqiangxu/ncrypt/pkg/instance"
	"github.com/zhiqiangxu/ncrypt/pkg/rest"
	"go.uber.org/zap"
)

func main() {

	app := &cli.App{
		Name:  "ncrypt server",
		Usage: "start ncrypt server",
		Action: func(c *cli.Context) (err error) {
			if c.Args().Len() < 2 {
				fmt.Println("at least two parties!")
				return
			}

			s := rest.NewServer()
			listenAddr := c.String("listen")
			fmt.Println("listenAddr", listenAddr)
			if listenAddr == "" {
				fmt.Println("please specify listen address!")
				return
			}

			err = instance.Secret().DecideSecret([]string(c.Args().Slice()))
			if err != nil {
				instance.Logger().Fatal("DecideSecret", zap.Error(err))
			}

			instance.Logger().Info("ncrypt ready to go!")
			s.Start(listenAddr)
			return
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "listen", Aliases: []string{"l"}},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
