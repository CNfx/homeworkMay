//*****************************
//用数组模拟QUEUE队列（先进先出）
//思路:
//1.定义BP，存储PUSH位置指针
//2.定义XP，存储POP位置指针
//3.自定义一个Queue数组类型,用切片来修改BP，XP,完成Queue数组
//*****************************
//疑问:是否利用了slice头部引用的堆其实和头部属性(ptr\len\cap)是一段连续的内存的特性？


package main

import (
	"errors"
	"fmt"
	)

type Queue []int


var (
	ErrQueueFull = errors.New("对列满")
	ErrQueueEmpty = errors.New("对列空")

)

func NewQueue(x int) Queue {
	s := make([]int,x+2) //创建一个切片，模拟Queue数组	
				
	s[0] = x + 1 	//[0]作为BP-用于记录当前push位置
	s[1] = x + 1	//[1]用于记录pop位置
	return s
}

func (self Queue) Push(v int) error {
	bp := self[0]	
	xp := self[1]
	

	//当xp小于栈容量时，认为xp下面内存是空的，可写入。push时以xp+1为地址

	caps := cap(self) - 1
	//println("push",bp, xp,caps)
	if xp < caps {
		xp += 1
		self[xp] = v
		self[1] = xp

		return nil
	}	

	//当bp、xp超出栈顶(数组下标1)时，报错
	if bp == 1 || xp == 1 {
		return	ErrQueueFull 
	}

	//push 
	self[bp] = v
	bp = bp -1
	self[0] = bp
	return nil
		
}

func (self Queue) Pop() (int, error) {
	bp := self[0]
	xp := self[1]

	//当xp大于总容量时，或者xp小于bp时，认为是空队列
	caps := cap(self) - 1
	if xp > caps || xp <= bp {
		return 0,ErrQueueEmpty
	}

	//pop 
	xp = xp - 1 
	v := self[xp]
	self[1] = xp	//调整xp
	//println("pop",bp, xp,caps)
	return v, nil
}

func main() {
	p := NewQueue(3)
	//caps := cap(p)
	//println(caps)
	fmt.Println(p.Push(4), p)
	fmt.Println(p.Push(3), p)
	fmt.Println(p.Push(2), p)
	fmt.Println(p.Push(1), p)
	println(p[4],&p[4])	//打印出p[4]的值和地址
	fmt.Println(p.Pop())
	fmt.Println(p.Push(3), p)
	println(p[4],&p[4])	//打印出p[4]的值和地址/
	fmt.Println(p.Pop())
	fmt.Println(p.Pop())
	fmt.Println(p.Pop())
	fmt.Println(p.Pop())


}
