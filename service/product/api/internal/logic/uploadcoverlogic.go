package logic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10MB
type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCoverLogic) UploadCover(req *http.Request) (resp *types.UploadCoverResponse, err error) {
	// todo: add your logic here and delete this line
	_ = req.ParseMultipartForm(maxFileSize)
	file, handler, err := req.FormFile("cover")
	if err != nil {
		return nil, errors.New("文件获取错误")
	}
	defer file.Close()
	name := genFilename(handler.Filename)
	_, err = l.svcCtx.Cosclient.Object.Put(context.Background(), name, file, nil)
	if err != nil {
		return nil, err
	}
	return &types.UploadCoverResponse{CoverUrl: genFileURL(name)}, nil
}
func genFilename(filename string) string {
	return fmt.Sprintf("%d_%s", time.Now().UnixMilli(), filename)
}

func genFileURL(objectKey string) string {
	return fmt.Sprintf("https://cover-1302313797.cos.ap-nanjing.myqcloud.com/%v", objectKey)
}
