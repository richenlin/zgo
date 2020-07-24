package helper

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	res := &Success{
		Success: true,
	}
	buf, _ := JSONMarshal(res)
	log.Println(string(buf))
	assert.NotNil(t, nil)
}
