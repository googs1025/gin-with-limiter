package src

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type cacheDate struct {
	key string
	value interface{}
	expireAt time.Time
}

func newCacheData(key string, value interface{}, expireAt time.Time) *cacheDate {
	return &cacheDate{
		key: key,
		value: value,
		expireAt: expireAt,
	}
}

type ListCache struct {
	list *list.List
	listMap map[string]*list.Element
	lock sync.Mutex
	maxsize int // 限制key 的最大数量

}

//// 原版 NewListCache
//func NewListCache() *ListCache {
//	return &ListCache{
//		list: list.New(),
//		listMap: make(map[string]*list.Element),
//		maxsize: 1000,
//	}
//}

func NewListCache(opt ...ListCacheOption) *ListCache {
	cache := &ListCache{
		list: list.New(),
		listMap: make(map[string]*list.Element),
		maxsize: 0,
	}
	ListCacheOptions(opt).apply(cache)
	cache.clear()
	return cache
}

type ListCacheOption func(l *ListCache)
type ListCacheOptions []ListCacheOption
func (lOpts ListCacheOptions) apply(l *ListCache) {
	for _, f := range lOpts {
		f(l)
	}
}

func WithMaxSize(size int) ListCacheOption {
	return func(l *ListCache) {
		if size > 0 {
			l.maxsize = size
		}
	}
}



// Get 获取缓存
func (l *ListCache) Get(key string) interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()

	if v, ok := l.listMap[key]; ok {
		// 如果过期了，
		if time.Now().After(v.Value.(*cacheDate).expireAt) {
			return nil
		}
		l.list.MoveToFront(v)
		return v.Value.(*cacheDate).value
	}

	return nil
}

const NotExpireTTL = time.Hour * 24 * 5 // 默认不过期时间是一年

// Put 放入缓存
func (l *ListCache) Put(key string, newValue interface{}, ttl time.Duration) {
	l.lock.Lock()
	defer l.lock.Unlock()

	var setExpire time.Time
	if ttl == 0 {
		setExpire = time.Now().Add(NotExpireTTL)
	} else {
		setExpire = time.Now().Add(ttl)
	}

	newCache := newCacheData(key, newValue, setExpire)
	if v, ok := l.listMap[key]; ok {
		v.Value = newCache
		l.list.MoveToFront(v)
	} else {
		l.listMap[key] = l.list.PushFront(newCache)
		// 判断list长度是否溢出。末位淘汰缓存
		if l.maxsize > 0 && len(l.listMap) > l.maxsize {
			l.RemoveOldest()
		}
	}

}

func (l *ListCache) Print() {

	ele := l.list.Front()
	if ele == nil {
		fmt.Println("链表中无元素")
		return
	}

	for {
		fmt.Println(l.Get(ele.Value.(*cacheDate).key))
		//fmt.Println(ele.Value.(*cacheDate).value)
		ele = ele.Next()
		if ele == nil {
			fmt.Println("元素已遍例完毕")
			break
		}

	}

}

// 清理
func (l *ListCache) clear() {
	go func() {

		for {
			// 每隔一秒执行
			l.removeExpire()
			time.Sleep(time.Second*1)
		}
	}()
}

// RemoveOldest 最老的一个淘汰
func (l *ListCache) RemoveOldest() {
	//l.lock.Lock()
	//defer l.lock.Unlock()
	back := l.list.Back()
	if back == nil {
		return
	}
	l.removeItem(back)

}


func (l *ListCache) removeItem(ele *list.Element) {
	key := ele.Value.(*cacheDate).key
	delete(l.listMap, key)

	l.list.Remove(ele)
}

// 删除过期时间的缓存
func (l *ListCache) removeExpire() {
	l.lock.Lock()
	defer l.lock.Unlock()

	for _, v := range l.listMap {
		//tmp := v.Value.(*cacheDate)
		if time.Now().After(v.Value.(*cacheDate).expireAt) {
			l.removeItem(v)
		}
	}

}

func (l *ListCache) Len() int {
	return len(l.listMap)
}
