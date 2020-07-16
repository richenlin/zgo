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
	success bool        // 请求成功, true
	data    interface{} // 响应数据
	traceID string      // 方便进行后端故障排除：唯一的请求ID
}

func (e *Result) Error() string {
	return "success"
}

// ErrorInfo 异常的请求结果体
type ErrorInfo struct {
	success      bool        // 请求成功, false
	data         interface{} // 响应数据
	errorCode    string      // 错误代码
	errorMessage string      // 向用户显示消息
	showType     int         //错误显示类型：0静音； 1条消息警告； 2消息错误； 4通知； 9页
	traceID      string      // 方便进行后端故障排除：唯一的请求ID
}

func (e *ErrorInfo) Error() string {
	return e.errorMessage
}

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
