/*
Package demo 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	make swagger

*/
package demo

// @title zgo-demo
// @version 0.0.1
// @description GIN + ENT/SQLX + CASBIN + WIRE
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes https http
// @basePath /api
// @contact.name suisrc
// @contact.email susirc@outlook.com
