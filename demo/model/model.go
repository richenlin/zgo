package model

var (
	// TableSchemaInit 是否初始化数据表结构
	TableSchemaInit = false

	// TableSchemaInitEnt 强制使用ent更新表结构, TableSchemaInit => true 无效
	TableSchemaInitEnt = true

	// TableSchemaInitSqlx 强制使用sqlx更新表结构, TableSchemaInit => true 无效
	TableSchemaInitSqlx = false
)