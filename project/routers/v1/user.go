package v1

import (
	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
	"go_gin_sample/project/service"
)

func UserAdd(c *gin.Context) {
	common.InitKey(c)
	req := service.UserAddRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.UserAdd(&req)

	c.JSON(200, &res)
}

func UserGet(c *gin.Context) {
	common.InitKey(c)
	req := service.UserGetRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}

	cps := service.NewServiceMgr(c)
	res := cps.UserGet(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}
