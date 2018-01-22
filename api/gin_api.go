package api

import "github.com/gin-gonic/gin"

type PingApi struct {

}

func (PingApi) Run() {
	r := gin.Default()

	r.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}