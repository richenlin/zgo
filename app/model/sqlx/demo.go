package sqlx

import "time"

// Demo 示例对象
type Demo struct {
	ID     int    `db:"id" json:"id"`                                         // 唯一标识
	Parent int    `db:"parent" json:"parent"`                                 // 父节点
	Code   string `db:"code" json:"code" binding:"required"`                  // 编号
	Name   string `db:"name" json:"name" binding:"required"`                  // 名称
	Memo   string `db:"memo" json:"memo"`                                     // 备注
	Status int    `db:"status"  json:"status" binding:"required,max=2,min=1"` // 状态(1:启用 2:停用)

	Creator   string    `db:"creator" json:"creator"`       // 创建者
	Updator   string    `db:"updator" json:"updator"`       // 创建者
	CreatedAt time.Time `db:"created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"` // 更新时间
}
