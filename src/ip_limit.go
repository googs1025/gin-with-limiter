package src

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"time"
)

type LimiterCache struct {
	data sync.Map	// key 是 ip+端口  value ==>bucket
}

var IpCache *LimiterCache

func init() {

	IpCache = &LimiterCache{}
	IpCache2 = NewListCache(WithMaxSize(1000))

}

// 基于请求的ip 实现限流。 基于令牌桶算法实现
func IpLimiter(cap int64, rate int64) func(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(handler gin.HandlerFunc) gin.HandlerFunc {

		return func(c *gin.Context) {
			ip := c.Request.RemoteAddr
			var limiter *Bucket
			if v, ok := IpCache.data.Load(ip); ok {
				limiter = v.(*Bucket)
			} else {
				limiter = NewBucket(cap, rate)
				IpCache.data.Store(ip, limiter)
			}

			if limiter.IsAccept() {
				handler(c)
			} else {
				c.AbortWithStatusJSON(400, gin.H{"message":"this ip too many requests!"})
			}
		}
	}
}

// 基于LRU实现，ip限流

var IpCache2 *ListCache
func IpLimiter1(cap int64, rate int64) func(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(handler gin.HandlerFunc) gin.HandlerFunc {

		return func(c *gin.Context) {
			ip := c.Request.RemoteAddr
			var limiter *Bucket

			if v := IpCache2.Get(ip); v != nil {
				limiter = v.(*Bucket)
			} else {
				limiter = NewBucket(cap, rate)
				log.Println("from cache")
				IpCache2.Put(ip, limiter, time.Second * 5)
			}
		}
	}
}
