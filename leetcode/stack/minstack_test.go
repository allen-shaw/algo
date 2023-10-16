package stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	/**
	  MinStack minStack = new MinStack();
	  minStack.push(-2);
	  minStack.push(0);
	  minStack.push(-3);
	  minStack.getMin();   --> 返回 -3.
	  minStack.pop();
	  minStack.top();      --> 返回 0.
	  minStack.getMin();   --> 返回 -2.
	*/
	ms := NewMinStack()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())
}
