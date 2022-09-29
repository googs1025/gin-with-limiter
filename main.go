package main

import (
	"github.com/gin-gonic/gin"
	"golanglearning/new_project/gin-with-limiting/src"
)

func main() {

	// 最简易的限流，没有引入中间件模式
	ratelimter := src.NewBucket(10)

	r := gin.New()
	//不可能每个请求都写判断，所以使用装饰器模式。
	r.GET("/", func(c *gin.Context) {
		if ratelimter.IsAccept() {
			c.JSON(200, gin.H{"message": "ok"})
		} else {
			c.AbortWithStatusJSON(400, gin.H{"message":"fail by rate limit"})
		}
	})

	defer func() {
		r.Run(":8081")
	}()

}


