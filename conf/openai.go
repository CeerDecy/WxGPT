package conf

import (
	_ "embed"
)

type OpenAI struct {
	AuthToken string
	BaseUrl   string
}
