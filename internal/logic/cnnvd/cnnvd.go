package cnnvd

import (
	"context"
	"focus-single/internal/dao"
	"focus-single/internal/model"
	"focus-single/internal/model/entity"
	"focus-single/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
)

type sCnnvd struct{}

func init() {
	service.RegisterCnnvd(New())
}

func New() *sCnnvd {
	return &sCnnvd{}
}

// GetList 查询内容列表
func (s *sCnnvd) GetList(ctx context.Context, in model.CnnvdGetListInput) (out *model.CnnvdGetListOutput, err error) {
	var (
		m = dao.Cnnvd.Ctx(ctx)
	)
	out = &model.CnnvdGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.Cnnvd
	if err := listModel.OrderDesc("modified").Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Cnnvd
	if err := listModel.OrderDesc("modified").ScanList(&out.List, "Cnnvd"); err != nil {
		return out, err
	}
	return
}

// GetDetail 查询详情
func (s *sCnnvd) GetDetail(ctx context.Context, id string) (out *model.CnnvdGetDetailOutput, err error) {
	out = &model.CnnvdGetDetailOutput{}
	if err := dao.Cnnvd.Ctx(ctx).WherePri(id).Scan(&out.Cnnvd); err != nil {
		return nil, err
	}
	return out, err
}

// Create 创建内容
func (s *sCnnvd) Create(ctx context.Context, in model.ContentCreateInput) (out model.ContentCreateOutput, err error) {
	if in.UserId == 0 {
		in.UserId = service.BizCtx().Get(ctx).User.Id
	}
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.Content.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ContentCreateOutput{ContentId: uint(lastInsertID)}, err
}

// Update 修改
func (s *sCnnvd) Update(ctx context.Context, in model.ContentUpdateInput) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.Content.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Content.Columns().Id).
			Where(dao.Content.Columns().Id, in.Id).
			Where(dao.Content.Columns().UserId, service.BizCtx().Get(ctx).User.Id).
			Update()
		return err
	})
}

// Delete 删除
func (s *sCnnvd) Delete(ctx context.Context, id uint) error {
	return dao.Content.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		user := service.BizCtx().Get(ctx).User
		// 管理员直接删除文章和评论
		if user.IsAdmin {
			_, err := dao.Content.Ctx(ctx).Where(dao.Content.Columns().Id, id).Delete()
			if err == nil {
				_, err = dao.Reply.Ctx(ctx).Where(dao.Reply.Columns().TargetId, id).Delete()
			}
			return err
		}
		// 删除内容
		_, err := dao.Content.Ctx(ctx).Where(g.Map{
			dao.Content.Columns().Id:     id,
			dao.Content.Columns().UserId: service.BizCtx().Get(ctx).User.Id,
		}).Delete()
		// 删除评论
		if err == nil {
			err = service.Reply().DeleteByUserContentId(ctx, user.Id, id)
		}
		return err
	})
}
