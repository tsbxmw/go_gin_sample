package service

import (
	"go_gin_sample/project/models"

	common "github.com/tsbxmw/gin_common"
)

type (
	UserAddResponse struct {
		common.Response
	}

	UserGetResponse struct {
		common.Response
		Data models.UserModel `json:"data"`
	}
)
