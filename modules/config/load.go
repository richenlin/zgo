package config

import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	"github.com/koding/multiconfig"
)

var (
	once sync.Once
)

// MustLoad 加载配置
func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			//if strings.HasSuffix(fpath, "ini") {
			//	loaders = append(loaders, &multiconfig.INILLoader{Path: fpath})
			//}
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yml") || strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		// 加载配置
		m.MustLoad(C)
	})
}

// Print 基于JSON格式输出配置
func Print() {
	if C.PrintConfig {
		PrintWithJSON()
	}
}

// PrintWithJSON 基于JSON格式输出配置
func PrintWithJSON() {
	b, err := json.MarshalIndent(C, "", " ")
	if err != nil {
		os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
		return
	}
	os.Stdout.WriteString(string(b) + "\n")
}
