package logger

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"zgo/modules/helper"

	"github.com/gin-gonic/gin"
)

// Debugf 写入调试日志
func Debugf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Debugf(format, args...)
}

// Infof 写入消息日志
func Infof(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Infof(format, args...)
}

// Printf 写入消息日志
func Printf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Printf(format, args...)
}

// Warnf 写入警告日志
func Warnf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Warnf(format, args...)
}

// Errorf 写入错误日志
func Errorf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Errorf(format, args...)
}

// Fatalf 写入重大错误日志
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	StartTrace(ctx).Fatalf(format, args...)
}

// ErrorStack 输出错误栈
func ErrorStack(ctx context.Context, err error) {
	StartTrace(ctx).WithField(StackKey, fmt.Sprintf("%+v", err)).Errorf(err.Error())
}

//=================================================================分割线
//=================================================================分割线
//=================================================================分割线

// 定义键名
const (
	TraceIDKey  = "trace_id"
	UserIDKey   = "user_id"
	RoleIDKey   = "role_id"
	VersionKey  = "version"
	HostnameKey = "hostname"
	StackKey    = "stack"
)

var (
	version     string
	pid         = os.Getpid()
	hostname, _ = os.Hostname()
)

// SetVersion 设定版本
func SetVersion(v string) {
	version = v
}

// FromTraceIDContext 从上下文中获取跟踪ID
func FromTraceIDContext(ctx context.Context) string {
	if v, ok := ctx.(*gin.Context); ok {
		return helper.GetTraceID(v)
	}
	return "main-" + strconv.Itoa(pid) // 系统上下文
}

// FromUserIDContext 从上下文中获取用户ID
func FromUserIDContext(ctx context.Context) (helper.UserInfo, bool) {
	if v, ok := ctx.(*gin.Context); ok {
		return helper.GetUserInfo(v)
	}
	return nil, false
}

// StartTrace 开始一个追踪单元
func StartTrace(ctx context.Context) *Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	fields := map[string]interface{}{
		VersionKey:  version,
		HostnameKey: hostname,
	}
	if v := FromTraceIDContext(ctx); v != "" {
		fields[TraceIDKey] = v
	}
	if v, ok := FromUserIDContext(ctx); ok {
		fields[UserIDKey] = v.GetUserID()
		fields[RoleIDKey] = v.GetRoleID()
	}

	return newEntryWithFields(fields)
}
