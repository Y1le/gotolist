package initialize

import (
	"github.com/gin-gonic/gin"

	api "github.com/Y1le/gotolist/interfaces/controller"
	middleware "github.com/Y1le/gotolist/interfaces/midddleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 鐢ㄦ埛鎿嶄綔
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())
		authed := v1.Group("/task/") // 闇€瑕佺櫥闄嗕繚鎶?
		authed.Use(middleware.JWT())
		{
			// 浠诲姟鎿嶄綔
			authed.POST("create", api.CreateTaskHandler())
			authed.GET("list", api.ListTaskHandler())
			authed.GET("detail", api.DetailTaskHandler())
			authed.POST("update", api.UpdateTaskHandler())
			authed.POST("search", api.SearchTaskHandler())
			authed.POST("delete", api.DeleteTaskHandler())
		}
	}
	return r
}
