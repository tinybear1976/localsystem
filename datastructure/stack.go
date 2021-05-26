package datastructure

import (
	"errors"
)

type Stack struct {
	Capacity int
	Top      int
	Data     []string
}

//入栈
func (this *Stack) Push(a string) (err error) {
	if this.isFull() {
		err = errors.New("stack full")
		return
	}
	this.Top++
	this.Data[this.Top] = a
	return
}

//出栈
func (this *Stack) Pop() (val string, err error) {
	if this.isEmpty() {
		err = errors.New("stack empty")
		return
	}
	val = this.Data[this.Top]
	this.Top--
	return
}

func (this *Stack) SeeTop() (val string, err error) {
	if this.isEmpty() {
		err = errors.New("stack empty")
		return
	}
	val = this.Data[this.Top]
	return
}

func (this *Stack) isFull() bool {
	return this.Top+1 >= this.Capacity
}

func (this *Stack) isEmpty() bool {
	return this.Top == -1
}

func (this *Stack) clear() {
	this.Top = -1
	for i := 0; i < len(this.Data); i++ {
		this.Data[i] = ""
	}

}

func NewStack(capacity int) *Stack {
	return &Stack{
		Capacity: capacity,
		Top:      -1,
		Data:     make([]string, capacity),
	}
}

func NewStackDefaultSize() *Stack {
	return &Stack{
		Capacity: 128,
		Top:      -1,
		Data:     make([]string, 128),
	}
}
