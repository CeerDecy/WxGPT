package conf

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OpenAI OpenAI `json:"openAI"`
}

var OpenAIConfig Config

//go:embed config.yml
var config string

func init() {
	err := yaml.Unmarshal([]byte(config), &OpenAIConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	//fmt.Printf("config → %+v\n", OpenAIConfig)
}
