package bucket

import (
	"log"
	"sync"
	"time"
)

//type Bucket struct {
//	cap int
//	tokens int
//	lock sync.Mutex
//	rate int // 每秒加入的令牌数
//}

type Bucket struct {
	cap      int64
	tokens   int64
	lock     sync.Mutex
	rate     int64 // 每秒加入的令牌数
	lastTime int64
}

func NewBucket(cap int64, rate int64) *Bucket {
	if cap <= 0 || rate <= 0 {
		log.Fatal("config wrong!")
	}
	b := &Bucket{
		cap:    cap,
		tokens: cap,
		rate:   rate,
	}

	return b
}

//// 使用goroutine的方式定时加入tokens
//// start 加入tokens
//func (b *Bucket) start() {
//	for {
//		time.Sleep(time.Second)
//		go b.addToken()
//	}
//}

//// addToken 加入token的方法
//func (b *Bucket) addToken() {
//	b.lock.Lock()
//	defer b.lock.Unlock()
//	// 如果 tokens:4 + rate:3 > cap:5 ，那就无法加入
//	if b.tokens + b.rate <= b.cap {
//		b.tokens += b.rate
//	} else {
//		b.tokens = b.cap
//	}
//
//
//}

// IsAccept 是否接受请求
func (b *Bucket) IsAccept() bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	now := time.Now().Unix()
	b.tokens = b.tokens + (now-b.lastTime)*b.rate

	if b.tokens >= b.cap {
		b.tokens = b.cap
	}
	b.lastTime = now

	if b.tokens > 0 {
		b.tokens--
		return true
	}

	return false
}
