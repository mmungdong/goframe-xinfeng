package model

type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   uint
}

type RotationCreateInput struct {
	RotationCreateUpdateBase
}

type RotationCreateOutput struct {
	RotationId uint `json:"rotation_id"`
}
