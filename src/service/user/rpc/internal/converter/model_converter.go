package converter

import (
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/model/entity"
)

func UserDBToRpc(userDB *entity.User) *user.UserInfo {

	userRpc := user.UserInfo{
		PostCount:    userDB.PostCount,
		School:       userDB.School,
		Grade:        userDB.Grade,
		Avatar:       userDB.Avatar,
		Introduction: userDB.Introduction,
		Username:     userDB.Username,
	}
	return &userRpc
}
