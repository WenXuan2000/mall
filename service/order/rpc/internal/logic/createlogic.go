package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		_, err1 := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			Id: in.Uid,
		})
		if err1 != nil {
			return fmt.Errorf("用户不存在")
		}
		newOrder := model.Order{
			Uid:    in.Uid,
			Pid:    in.Pid,
			Amount: in.Amount,
			Status: 0,
		}
		_, err1 = l.svcCtx.OrderModel.TxInsert(l.ctx, tx, &newOrder)
		if err1 != nil {
			return fmt.Errorf("订单创建失败")
		}
		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &order.CreateResponse{}, nil
}
