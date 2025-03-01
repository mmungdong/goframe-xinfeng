package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationCreateReq struct {
	g.Meta `path:"/backend/rotation" method:"post" tags:"轮播图" summary:"创建轮播图" `
	PicUrl string `json:"pic_url" v:"required#轮播图图片链接不能为空" dc:"轮播图图片链接"`
	Link   string `json:"link" v:"required#轮播图跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort" dc:"排序"`
}

type RotationCreateRes struct {
	RotationId int `json:"rotation_id" dc:"轮播图ID"`
}
