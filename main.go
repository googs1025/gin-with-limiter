package main

import (
	"gin-with-limiting/pkg"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func main() {

	r := gin.New()
	// 使用装饰器模式。
	//r.GET("/", pkg.Limiter(10)(test))
	r.GET("/", pkg.ParamLimiter(5, 1, "name")(test)) // 基于令牌桶实现的对请求参数的限流
	r.GET("/aaa", pkg.IpLimiter(4, 1)(test))         // 令牌桶实现的ip限流
	r.GET("/aaa", pkg.IpLimiter1(4, 1)(test))        // 基于LRU实现的ip限流

	defer func() {
		r.Run(":8081")
	}()

}
