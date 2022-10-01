package main

import (
	"container/list"
	"fmt"
	"log"
)

// 练习标准库中的链表

func main() {


	ll := &list.List{}
	e1 := ll.PushFront("e1")	// 向头部插入元素

	ll.PushFront("e2")

	ll.InsertAfter("e2", e1)

	ele := ll.Front()
	if ele == nil {
		log.Fatal("nil element!")
	}

	for {
		fmt.Println(ele.Value)
		if ele.Next() == nil {
			break
		}
		ele = ele.Next()
	}



}
