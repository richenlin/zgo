// +build document

package ent

/*
 如果是复杂的逻辑关系或者对sql没有接触,可以使用ent处理
*/
import (
	"github.com/facebookincubator/ent"
)

// Demo 用例, 没有任何意义
type Demo struct {
	ent.Schema
}
