package service

import (
	common "github.com/tsbxmw/gin_common"
	"go_gin_sample/project/models"
	"time"
)

func (cps *ProjectService) UserAdd(req *UserAddRequest) *UserAddResponse {
	res := UserAddResponse{}
	userModel := models.UserModel{}
	if err := common.DB.Table(userModel.TableName()).
		Where("user_nickname=?", req.UserNickname).First(&userModel).Error; err != nil {
		if err.Error() == "record not found" {
			userModel.Remark = req.Remark
			userModel.UserNickname = req.UserNickname
			userModel.BaseModel.ModifiedTime = time.Now()
			userModel.BaseModel.CreationTime = time.Now()

			if err = common.DB.Table(userModel.TableName()).Create(&userModel).Error; err != nil {
				cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
				panic(err)
			} else {
				res.Code = common.HTTP_RESPONSE_OK
				res.Message = common.HTTP_MESSAGE_OK
				res.Data = []string{}
			}
		} else {
			cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
			panic(err)
		}
	} else {
		res.Code = 0
		res.Message = "User already exists"
	}
	return &res
}

func (cps *ProjectService) UserGet(req *UserGetRequest) *UserGetResponse {
	userModel := models.UserModel{}

	if err := common.DB.Table(userModel.TableName()).
		Where("user_nickname=?", req.UserNickname).Find(&userModel).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	res := UserGetResponse{
		Response: common.Response{
			Message: common.HTTP_MESSAGE_OK,
			Code:    common.HTTP_RESPONSE_OK,
		},

		Data: userModel,
	}
	return &res
}
