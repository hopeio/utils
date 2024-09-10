package linkedlist

import (
	"fmt"
	"github.com/hopeio/utils/log"
	"testing"
)

func TestList(t *testing.T) {
	// 打印链表信息
	var l LinkedList[int]
	fmt.Println("###############################################")
	fmt.Println("链表长度为：", l.Len())
	fmt.Println("链表是否为空:", l.IsEmpty())
	fmt.Print("遍历链表：")
	l.traverse(func(data int) { log.Info(data, " ") })
	fmt.Println("###############################################")

}
