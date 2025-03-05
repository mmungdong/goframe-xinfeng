package rotation

import (
	"context"

	"goframe-xinfeng/internal/service"

	"goframe-xinfeng/internal/dao"
	"goframe-xinfeng/internal/model"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/database/gdb"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建内容
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: uint(lastInsertID)}, err
}

func (s *sRotation) Delete(ctx context.Context, id uint) error {

	// 软删除
	_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
		dao.RotationInfo.Columns().Id: id,
	}).Delete()

	// 忽视 deleted_at 字段， 直接硬删除，如果表中没有 deleted_at 字段， 则使用上面的软删除效果和这个硬删除效果一致
	// _, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
	// 	dao.RotationInfo.Columns().Id: id,
	// }).Unscoped().Delete()

	return err
}

func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id). // goframe的字段过滤， 排除id字段
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}
