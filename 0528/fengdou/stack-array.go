package main 
import (
	"errors"
)

type Stack []int

var (
	ErrStackFull = errors.New("Stack full")
	ErrStackEmpty = errors.New("stack empty")
)


func NewStack(x int) Stack {
	s := make([]int,x+1)
	s[0] = x	//充当BP寄存器,相当于编译指令：
			//C style:MOV BP,cap
			//go style:sub SP,cap
	return s
}


//模拟往栈帧加数据, 0时到栈顶，弹出错误
func (self Stack) Push(v int) error {
	i := self[0]
	if i == 0 {
		return ErrStackFull
	}
	
	self[i] = v
	self[0] = i - 1	//调整BP寄存器
	return nil

}

//模拟弹出数据
func (self Stack) Pop() (int, error) {
	i := self[0] + 1		//先设置栈底
	if i == cap(self) {	//如果己到栈底，则无数据可弹
		return 0, ErrStackEmpty
	}

	v := self[i]
	self[0] = i	//调整BP寄存器
	return v, nil
}

func main() {
	p := NewStack(3)
	println(cap(p))
	p.Push(1)
	p.Push(2)
	p.Push(3)
	p.Push(4)
	p.Push(5)

	println("pop", p.Pop)

}
