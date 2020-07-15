package result

// PaginationResult 响应列表数据
//type PaginationResult struct {
//	list  interface{} `json:"list"`
//	total int         `json:"total"`
//	sign  string      `json:"sign"`
//}

// PaginationParam 分页查询条件
type PaginationParam struct {
	PageSign  string `query:"pageSign"`                              // 请求参数, total | list | both
	PageNo    uint   `query:"pageNo,default=1"`                      // 当前页
	PageSize  uint   `query:"pageSize,default=20" binding:"max=100"` // 页大小
	PageTotal uint   `query:"pageTotal"`                             // 上次统计的数据条数
}
