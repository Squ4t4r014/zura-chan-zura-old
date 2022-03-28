package main

import (
	"github.com/gin-gonic/gin"
	"zura-chan-zura-old/app/infrastructure"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := infrastructure.NewRouting()
	_ = r.Run()
}
