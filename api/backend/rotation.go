package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationCreateReq struct {
	g.Meta `path:"/backend/rotation" method:"post" tags:"轮播图" summary:"创建轮播图" `
	PicUrl string `json:"pic_url" v:"required#轮播图图片链接不能为空" dc:"轮播图图片链接"`
	Link   string `json:"link" v:"required#轮播图跳转链接不能为空" dc:"跳转链接"`
	Sort   uint   `json:"sort" dc:"排序"`
}

type RotationCreateRes struct {
	RotationId uint `json:"rotation_id" dc:"轮播图ID"`
}

type RotationDeleteReq struct {
	g.Meta     `path:"/backend/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图"`
	RotationId uint `json:"rotation_id" v:"required#轮播图IDD不能为空" dc:"轮播图ID"`
}

type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/backend/rotation/update/{Id}" method:"post" tags:"内容" summary:"更新轮播图"`
	Id     uint   `json:"id" v:"required#轮播图ID不能为空" dc:"轮播图ID"`
	PicUrl string `json:"pic_url" v:"required#轮播图图片链接不能为空" dc:"轮播图图片链接"`
	Link   string `json:"link" v:"required#轮播图跳转链接不能为空" dc:"跳转链接"`
	Sort   uint   `json:"sort" dc:"排序"`
}

type RotationUpdateRes struct{}
