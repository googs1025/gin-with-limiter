package main

import (
	"github.com/gin-gonic/gin"
	"golanglearning/new_project/gin-with-limiting/src"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func main() {



	r := gin.New()
	// 使用装饰器模式。
	r.GET("/", src.Limiter(10)(test))

	defer func() {
		r.Run(":8081")
	}()

}


