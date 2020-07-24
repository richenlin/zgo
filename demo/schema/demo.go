package schema

// DemoSet 示例对象
type DemoSet struct {
	Code string `json:"code" binding:"required"` // 编号
	Name string `json:"name" binding:"required"` // 名称
	Memo string `json:"memo"`                    // 备注
}
