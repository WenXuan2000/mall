package logic

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	var res = &model.User{}
	if err := l.svcCtx.UserModel.Where("id=?", in.Id).Take(res).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(100, "用户不存在")
	}

	return &user.UserInfoResponse{
		Id:     int64(res.Model.ID),
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil

}
