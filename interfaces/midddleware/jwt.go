package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Y1le/gotolist/consts"
	"github.com/Y1le/gotolist/infrastructure/auth"
	lctx "github.com/Y1le/gotolist/infrastructure/common/context"
)

// JWT token楠岃瘉涓棿浠?
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = consts.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(consts.InvalidParams, gin.H{
				"status": code,
				"msg":    consts.GetMsg(code),
				"data":   "缂哄皯Token",
			})
			c.Abort()
			return
		}
		jwtService := auth.NewJWTTokenService()
		claims, err := jwtService.ParseToken(c.Request.Context(), token)
		if err != nil {
			code = consts.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = consts.ErrorAuthCheckTokenTimeout
		}

		if code != consts.SUCCESS {
			c.JSON(consts.InvalidParams, gin.H{
				"status": code,
				"msg":    consts.GetMsg(code),
				"data":   "鍙兘鏄韩浠借繃鏈熶簡锛岃閲嶆柊鐧诲綍",
			})
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(
			lctx.NewContext(
				c.Request.Context(),
				&lctx.UserInfo{Id: claims.Id, Name: claims.Username},
			),
		)
		c.Next()
	}
}
