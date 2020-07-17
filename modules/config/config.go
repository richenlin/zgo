package config

import (
	"fmt"
)

var (
	// C 全局配置(需要先执行MustLoad，否则拿不到配置)
	C = new(Config)
)

// Config 配置参数
type Config struct {
	RunMode      string
	Swagger      bool
	PrintConfig  bool
	HTTP         HTTP
	Casbin       Casbin
	Logging      Logging
	RateLimiter  RateLimiter
	JWTAuth      JWTAuth
	CORS         CORS
	GZIP         GZIP
	Redis        Redis
	MySQL        MySQL
	Postgres     Postgres
	Sqlite3      Sqlite3
	WWW          WWW
	MiddleConfig MiddleConfig
}

// IsDebugMode 是否是debug模式
func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

// MiddleConfig 中间件启动和关闭
type MiddleConfig struct {
	Logger  bool
	Recover bool
}

// HTTP http配置参数
type HTTP struct {
	Host             string //`yaml:"zgo,http,host"`
	Port             int    //`yaml:"zgo,http,port"`
	CertFile         string
	KeyFile          string
	ShutdownTimeout  int
	MaxContentLength int64
	ContextPath      string
	Prefixes         []string
}

// Casbin casbin配置参数
type Casbin struct {
	Enable           bool
	Debug            bool
	Model            string //
	AutoLoad         bool
	AutoLoadInternal int
	PolicyType       string // file | mysql | sqlite3 | postgres | redis | restful
	PolicySource     string // policy.json | root:1234@tcp(127.0.0.1:3306)/yourdb | http://xxx.xxx/api/casbin/policy.rule
	PolicyTable      string // casbin_rule
}

// Logging 日志配置参数
type Logging struct {
	Level            string
	Format           string // json | text
	Output           string
	OutputFile       string
	EnableSyslogHook bool
	SyslogNetwork    string
	SyslogAddr       string
	SyslogTag        string
	//SyslogPriority   int
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

// GZIP gzip压缩
type GZIP struct {
	Enable             bool
	ExcludedExtentions []string
	ExcludedPaths      []string
}

// JWTAuth 用户认证
type JWTAuth struct {
	Enable        bool
	SigningMethod string
	SigningKey    string
	Expired       int
	Store         string
	FilePath      string
	RedisDB       int
	RedisPrefix   string
}

// RateLimiter 请求频率限制配置参数
type RateLimiter struct {
	Enable  bool
	Count   int64
	RedisDB int
}

// Redis redis配置参数
type Redis struct {
	Addr     string
	Password string
}

// WWW 静态资源
type WWW struct {
	Index   string
	RootDir string
}

//===============================================分割线
//===============================================分割线
//===============================================分割线

// MySQL mysql配置参数
type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}

// Sqlite3 sqlite3配置参数
type Sqlite3 struct {
	Path string
}

// DSN 数据库连接串
func (a Sqlite3) DSN() string {
	return a.Path
}
