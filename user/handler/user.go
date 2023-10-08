package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"

	user "user/proto"
)

type User struct {
	UserDataService service.IUserDataService
}

// 注册
func (u *User) Register(ctx context.Context, urq *user.UserRegisterRequest, urp *user.UserRegisterResponse) error {
	userRegister := &model.User{
		Username:     urq.UserName,
		FristName:    urq.FristName,
		HashPassword: urq.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	urp.Message = "添加成功"
	return nil
}

// 登录
func (u *User) Login(ctx context.Context, ulq *user.UserLoginRequest, ulp *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(ulq.UserName, ulq.Pwd)
	if err != nil {
		return err
	}
	ulp.IsSuccess = isOk
	return nil
}
func (u *User) GetUserInfo(ctx context.Context, uiq *user.UserInfoRequest, uip *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(uiq.UserName)
	if err != nil {
		return err
	}
	uip = UserForResponse(userInfo)
	return nil
}

// 类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.Username
	response.FristName = userModel.FristName
	response.UserId = userModel.ID
	return response
}
