// +build document

package sqlx

/*
 如果是简单sql直接处理以及对sql有一定的能力,推荐使用sqlx处理
*/
import (
	"github.com/jmoiron/sqlx"
)

// Demo 用例, 没有任何意义
func Demo() {
	sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
}
