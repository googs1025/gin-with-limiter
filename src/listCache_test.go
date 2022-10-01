package src

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

//func TestListCache(t *testing.T) {
//	convey.Convey("Test ListCache", t, func() {
//		cache := NewListCache()
//		cache.Put("name", "jiang")
//		cache.Put("age", 134)
//		cache.Put("sex", "man")
//		fmt.Println("-------------------------")
//		cache.Print()
//		fmt.Println("-------------------------")
//		cache.Get("age")
//		cache.Print()
//		fmt.Println("-------------------------")
//		cache.RemoveOldest()
//		cache.Print()
//		fmt.Println("-------------------------")
//	})
//}

//func TestWithMaxSizeListCache(t *testing.T) {
//	convey.Convey("测试有最大缓存size的链表LRU", t, func(){
//				cache := NewListCache(WithMaxSize(4))
//				cache.Put("name", "jiang")
//				cache.Put("age", 134)
//				cache.Put("sex", "man")
//				cache.Put("sex234", "man")
//				cache.Put("sex111", "man")
//				fmt.Println("-------------------------")
//				cache.Print()
//				fmt.Println("-------------------------")
//				cache.Get("age")
//				cache.Print()
//				fmt.Println("-------------------------")
//				cache.RemoveOldest()
//				cache.Print()
//				fmt.Println("-------------------------")
//	})
//}

func TestWithMaxSizeWithTtlListCache(t *testing.T) {
	convey.Convey("测试有最大缓存size和过期时间的链表LRU", t, func() {
		cache := NewListCache(WithMaxSize(4))
		cache.Put("name", "jiang", time.Second * 7)
		cache.Put("age", 134, time.Second * 7)
		cache.Put("sex", "man", 0)
		cache.Put("sex234", "man", time.Second * 7)
		cache.Put("sex111", "man", time.Second * 7)
		fmt.Println("-------------------------")
		//cache.Print()
		fmt.Println("-------------------------")
		fmt.Println(cache.Get("age"))
		//cache.Print()
		fmt.Println("-------------------------")
		//cache.RemoveOldest()
		//cache.Print()
		fmt.Println("-------------------------")


		for {
			fmt.Println(cache.Get("name"))
			fmt.Printf("name=%v, age=%v, sex=%v", cache.Get("name"), cache.Get("age"), cache.Get("sex"))
			time.Sleep(time.Second)
		}

	})
}
