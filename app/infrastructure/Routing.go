package infrastructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Routing struct {
	Gin          *gin.Engine
	AbsolutePath string
}

func NewRouting() *Routing {
	c, _ := NewConfig()
	r := &Routing{
		Gin:          gin.Default(),
		AbsolutePath: c.AbsolutePath,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	r.Gin.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "http://zura-chan-zura.com/")
	})
}

func (r *Routing) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return r.Gin.Run(":" + port)
}
