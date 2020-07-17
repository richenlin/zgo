package helper

// 定义上下文中的键
const (
	prefix      = "zgo"
	UserInfoKey = prefix + "/user-info"
	TraceIDKey  = prefix + "/tract-id"
	ReqBodyKey  = prefix + "/req-body"
	ResBodyKey  = prefix + "/res-body"
)

const (
	// ShowNone 静音
	ShowNone = 0
	// ShowWarn 消息警告
	ShowWarn = 1
	// ShowError 消息错误
	ShowError = 2
	// ShowNotify 通知；
	ShowNotify = 4
	// ShowPage 页
	ShowPage = 9
)

// Result 正常请求结构体
type Result struct {
	Success bool        // 请求成功, true
	Data    interface{} // 响应数据
	TraceID string      // 方便进行后端故障排除：唯一的请求ID
}

func (e *Result) Error() string {
	return "success"
}

// ErrorInfo 异常的请求结果体
type ErrorInfo struct {
	Success      bool        `json:"success"`        // 请求成功, false
	Data         interface{} `json:"data,omitempty"` // 响应数据
	ErrorCode    string      `json:"errorCode"`      // 错误代码
	ErrorMessage string      `json:"errorMessage"`   // 向用户显示消息
	ShowType     int         `json:"showType"`       //错误显示类型：0静音； 1条消息警告； 2消息错误； 4通知； 9页
	TraceID      string      `json:"traceId"`        // 方便进行后端故障排除：唯一的请求ID
}

func (e *ErrorInfo) Error() string {
	return e.ErrorMessage
}

// PaginationResult 响应列表数据
//type PaginationResult struct {
//	list  interface{} `json:"list"`
//	total int         `json:"total"`
//	sign  string      `json:"sign,omitempty" binding:"required"`
//}

// PaginationParam 分页查询条件
type PaginationParam struct {
	PageSign  string `query:"pageSign"`                              // 请求参数, total | list | both
	PageNo    uint   `query:"pageNo,default=1"`                      // 当前页
	PageSize  uint   `query:"pageSize,default=20" binding:"max=100"` // 页大小
	PageTotal uint   `query:"pageTotal"`                             // 上次统计的数据条数
}

// UserInfo 用户信息
type UserInfo interface {
	GetUserID() string
	GetRoleID() string
}

// UserInfoFunc user
type UserInfoFunc interface {
	GetUserInfo() (UserInfo, bool)
	SetUserInfo(UserInfo)
}
