// Package routes 注册路由
package routes

import (
	controllers "goblog/app/http/controllers/api/v1"
	"goblog/pkg/config"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	uc := new(controllers.UsersController)

	// 获取当前用户
	usersGroup := v1.Group("/users")
	{
		usersGroup.GET("/hello", uc.Hello)
	}
}
