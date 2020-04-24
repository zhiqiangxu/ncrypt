package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	g "github.com/zhiqiangxu/ncrypt/pkg/rest/gin"
)

// New for api
func New() *gin.Engine {
	r := g.New()

	ncrypt := r.Group("/ncrypt")
	ncrypt.POST("/encrypt", Encrypt)
	ncrypt.POST("/decrypt", Decrypt)
	return r
}

func sendoutput(c *gin.Context, code int, output interface{}) {
	if code == 0 {
		code = http.StatusOK
	}

	c.JSON(code, output)
}
