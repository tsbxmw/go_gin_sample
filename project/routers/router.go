package routers

import (
	v1 "go_gin_sample/project/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	GroupV1 := r.Group("/v1")
	{
		user := GroupV1.Group("/user")
		{
			user.POST("/", v1.UserAdd)
			user.GET("/", v1.UserGet)
		}

	}
}
