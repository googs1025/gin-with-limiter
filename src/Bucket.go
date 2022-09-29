package src

import (
	"log"
	"sync"
)

type Bucket struct {
	cap int
	tokens int
	lock sync.Mutex
}

func NewBucket(cap int) *Bucket {
	if cap <= 0 {
		log.Fatal("cap can't == 0!")
	}
	return &Bucket{
		cap: cap,
		tokens: cap,
	}
}

func (b *Bucket) IsAccept() bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.tokens > 0 {
		return true
	}

	return false
}


