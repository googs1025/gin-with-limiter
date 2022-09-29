package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//ratelimter := src.NewBucket(10)

	r := gin.New()
	// 不可能每个请求都写判断，所以使用装饰器模式。
	//r.GET("/", func(c *gin.Context) {
	//	if ratelimter.IsAccept() {
	//		c.JSON(200, gin.H{"message": "ok"})
	//	} else {
	//		c.AbortWithStatusJSON(400, gin.H{"message":"fail by rate limit"})
	//	}
	//})

	defer func() {
		r.Run(":8081")
	}()

}


