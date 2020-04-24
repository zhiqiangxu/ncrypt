package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiqiangxu/ncrypt/pkg/io"
	"github.com/zhiqiangxu/ncrypt/pkg/service"
)

// Decrypt service
func Decrypt(c *gin.Context) {
	var input io.DecryptInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	output := service.Instance().Decrypt(input)
	sendoutput(c, output.Code, output)
}
