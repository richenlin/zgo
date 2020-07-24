package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	log.Println("===========================loading")
	MustLoad("../../configs/config.toml")
	PrintWithJSON()
	assert.NotNil(t, nil) // 异常才能显示日志
}
