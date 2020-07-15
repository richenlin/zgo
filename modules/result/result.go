package result

// 定义上下文中的键
const (
	prefix      = "zgo"
	UserInfoKey = prefix + "/user-info"
	TraceIDKey  = prefix + "/tract-id"
	ReqBodyKey  = prefix + "/req-body"
	ResBodyKey  = prefix + "/res-body"
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
