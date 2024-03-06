package logic

import (
	"context"
	"fmt"
	"mall/common/jwtx"
	"mall/service/user/rpc/types/user"
	"time"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	tokenkey := fmt.Sprintf("%s%v", "api:token:", res.Id)
	v, err := l.svcCtx.Rds.GetCtx(l.ctx, tokenkey)
	if err == nil {
		token := v
		claims, err1 := jwtx.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
		if err1 == nil {
			exptime := claims.RegisteredClaims.ExpiresAt
			if time.Now().Before(exptime.Time) {
				return &types.LoginResponse{
					AccessToken:  token,
					AccessExpire: exptime.Unix() + l.svcCtx.Config.Auth.AccessExpire,
				}, nil
			}
		}
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GenToken(jwtx.JwtPayLoad{
		UserID:   res.Id,
		Username: res.Name,
	}, l.svcCtx.Config.Auth.AccessSecret, accessExpire)
	//accessToken, err := jwtx.GenToken(l.svcCtx.Config.Auth.AccessSecret, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}
	if err := l.svcCtx.Rds.Set(tokenkey, accessToken); err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
