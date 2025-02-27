package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/model"
	"focus-single/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

// Reply 回复控制器
var Reply = cReply{}

type cReply struct{}

func (a *cReply) GetListContent(ctx context.Context, req *v1.ReplyGetListContentReq) (res *v1.ReplyGetListContentRes, err error) {
	out, err := service.Reply().GetList(ctx, model.ReplyGetListInput{
		Page:       req.Page,
		Size:       req.Size,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
	})
	if err != nil {
		return nil, err
	}
	request := g.RequestFromCtx(ctx)
	service.View().RenderTpl(ctx, "index/reply.html", model.View{Data: out})
	tplContent := request.Response.BufferString()
	request.Response.ClearBuffer()
	return &v1.ReplyGetListContentRes{Content: tplContent}, nil
}

func (a *cReply) Create(ctx context.Context, req *v1.ReplyCreateReq) (res *v1.ReplyCreateRes, err error) {
	err = service.Reply().Create(ctx, model.ReplyCreateInput{
		Title:      req.Title,
		ParentId:   req.ParentId,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		Content:    req.Content,
		UserId:     service.Session().GetUser(ctx).Id,
	})
	return
}

func (a *cReply) Delete(ctx context.Context, req *v1.ReplyDeleteReq) (res *v1.ReplyDeleteRes, err error) {
	err = service.Reply().Delete(ctx, req.Id)
	return
}
