package service

type (
	UserAddRequest struct {
		UserNickname string `json:"user_nickname" binding:"required"`
		Age     int `json:"age" binding:"required"`
		Remark       string `json:"remark"`
	}

	UserGetRequest struct {
		UserNickname string `form:"user_nickname" binding:"required"`
	}
)
