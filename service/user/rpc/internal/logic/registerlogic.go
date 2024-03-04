package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mall/common/cryptx"
	"mall/service/user/model"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	isUnique := l.IsMobileUnique(in.Mobile)

	if isUnique {
		return nil, status.Error(100, "该用户已存在")
	}

	newUser := model.User{
		Name:     in.Name,
		Gender:   in.Gender,
		Mobile:   in.Mobile,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
	}

	if res := l.svcCtx.UserModel.Create(&newUser); res.Error != nil {
		return nil, status.Error(500, res.Error.Error())
	}
	return &user.RegisterResponse{
		Id:     int64(newUser.Model.ID),
		Name:   newUser.Name,
		Gender: newUser.Gender,
		Mobile: newUser.Mobile,
	}, nil
}
func (l *RegisterLogic) IsMobileUnique(mobile string) bool {
	var userExist = &model.User{}
	if err := l.svcCtx.UserModel.Where("mobile=?", mobile).First(&userExist).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
