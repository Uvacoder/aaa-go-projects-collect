package main

import "fmt"

// customized implementation of stack to use for vertices
// main implementation in ==> https://github.com/Armingodiz/GoWorld/tree/master/Data-Structures/stack

type Stack struct {
	stack []*Vertex
	top   int
}

func CreateStack() *Stack {
	stack := &Stack{
		stack: []*Vertex{},
		top:   -1,
	}
	return stack
}

func (stack *Stack) IsEmpty() bool {
	if stack.top == -1 {
		fmt.Println("STACK IS EMPTY !")
		return true
	}
	return false
}

func (stack *Stack) Push(element *Vertex) {
	stack.stack = append(stack.stack, element)
	stack.top += 1
}

func (stack *Stack) Pop() *Vertex {
	if stack.IsEmpty() {
		return nil
	} else {
		popped := stack.stack[stack.top]
		stack.top -= 1
		stack.stack = stack.stack[:stack.top+1]
		return popped
	}
}

func (stack *Stack) Peak() *Vertex {
	if stack.IsEmpty() {
		return nil
	} else {
		return stack.stack[stack.top]
	}
}
