package sqlxm

import "time"

// Demo 示例对象
type Demo struct {
	ID     int    `db:"id"`                                    // 唯一标识
	Parent int    `db:"parent"`                                // 父节点
	Code   string `db:"code" binding:"required"`               // 编号
	Name   string `db:"name" binding:"required"`               // 名称
	Memo   string `db:"memo" `                                 // 备注
	Status int    `db:"status" binding:"required,max=2,min=1"` // 状态(1:启用 2:停用)

	Creator   string    `db:"creator"`    // 创建者
	Updator   string    `db:"updator"`    // 创建者
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
}
