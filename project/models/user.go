package models

import (
	common "github.com/tsbxmw/gin_common"
)

type (
	UserModel struct {
		common.BaseModel
		UserNickname string `json:"user_nickname"`
		Age          int    `json:"age"`
		Remark       string `json:"remark"`
	}
)

func (UserModel) TableName() string {
	return "user"
}
