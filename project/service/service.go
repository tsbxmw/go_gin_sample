package service

import (
	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
)

type (
	ProjectService struct {
		common.BaseService
	}
)

func NewServiceMgr(c *gin.Context) *ProjectService {
	return &ProjectService{
		BaseService: common.BaseService{
			Ctx: c,
		},
	}
}
