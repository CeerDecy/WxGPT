package main

import (
	"WxGPT/internal/router"
)

func main() {
	engine := router.Engine()
	_ = engine.Run(":80")

}
