package controller

import (
	"context"
	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

// Cnnvd 漏洞库管理
var Cnnvd = cCnnvd{}

type cCnnvd struct{}

// Index article list
func (a *cCnnvd) Index(ctx context.Context, req *v1.CnnvdIndexReq) (res *v1.CnnvdIndexRes, err error) {
	getListRes, err := service.Cnnvd().GetList(ctx, model.CnnvdGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	service.View().Render(ctx, model.View{
		ContentType: consts.CnnvdTypeArticle,
		Data:        getListRes,
	})
	return
}

// Detail .article details
func (a *cCnnvd) Detail(ctx context.Context, req *v1.CnnvdDetailReq) (res *v1.CnnvdDetailRes, err error) {
	getDetailRes, err := service.Cnnvd().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if getDetailRes == nil {
		service.View().Render404(ctx)
		return nil, nil
	}
	service.View().Render(ctx, model.View{
		ContentType: consts.CnnvdTypeArticle,
		Data:        getDetailRes,
	})
	return
}
