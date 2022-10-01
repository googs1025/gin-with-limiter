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
	//r.GET("/", src.Limiter(10)(test))
	r.GET("/", src.ParamLimiter(5, 1, "name")(test))	// 基于令牌桶实现的对请求参数的限流
	r.GET("/aaa", src.IpLimiter(4, 1)(test))	// 令牌桶实现的ip限流
	r.GET("/aaa", src.IpLimiter1(4, 1)(test))  // 基于LRU实现的ip限流

	defer func() {
		r.Run(":8081")
	}()

}


