package result

// 定义错误
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
var (
	Err400BadRequest       = ErrorModel{showType: ShowWarn, errorCode: "ERR-BAD-REQUEST", errorMessage: "请求发生错误"}
	Err401Unauthorized     = ErrorModel{showType: ShowWarn, errorCode: "ERR-UNAUTHORIZED", errorMessage: "用户没有权限（令牌、用户名、密码错误）"}
	Err403Forbidden        = ErrorModel{showType: ShowWarn, errorCode: "ERR-FORBIDDEN", errorMessage: "用户得到授权，但是访问是被禁止的"}
	Err404NotFound         = ErrorModel{showType: ShowWarn, errorCode: "ERR-NOT-FOUND", errorMessage: "发出的请求针对的是不存在的记录，服务器没有进行操作"}
	Err405MethodNotAllowed = ErrorModel{showType: ShowWarn, errorCode: "ERR-METHOD-NOT-ALLOWED", errorMessage: "请求的方法不允许"}
	Err406NotAcceptable    = ErrorModel{showType: ShowWarn, errorCode: "ERR-NOT-ACCEPTABLE", errorMessage: "请求的格式不可得"}
	Err429TooManyRequests  = ErrorModel{showType: ShowWarn, errorCode: "ERR-TOO-MANY-REQUESTS", errorMessage: "请求次数过多"}
	Err500InternalServer   = ErrorModel{showType: ShowWarn, errorCode: "ERR-INTERNAL-SERVER", errorMessage: "服务器发生错误"}
)
