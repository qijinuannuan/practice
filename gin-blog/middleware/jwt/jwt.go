package jwt

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := ctx.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
