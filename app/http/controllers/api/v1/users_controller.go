package v1

import (
	"goblog/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) Hello(c *gin.Context) {

	response.JSON(c, "你好")
}
