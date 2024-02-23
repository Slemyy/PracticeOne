package Stack

import (
	"errors"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

type Stack struct {
	head *node
}

func Create() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value interface{}) {
	newNode := &node{data: value}

	if stack.head == nil {
		stack.head = newNode
	} else {
		newNode.next = stack.head
		stack.head = newNode
	}
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.head == nil {
		return nil, errors.New("stack is empty")
	}

	value := stack.head.data
	stack.head = stack.head.next

	return value, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.head == nil {
		return nil, errors.New("stack is empty")
	}

	return stack.head.data, nil
}

func (stack *Stack) isEmpty() bool {
	if stack.head == nil {
		return true
	} else {
		return false
	}
}

func (stack *Stack) Reverse() (*Stack, error) {
	newStack := Create()

	if stack.head == nil || stack.head.next == nil {
		return nil, errors.New("the stack is empty or contains only one element")
	}

	for {
		if stack.head == nil {
			break
		}

		value, _ := stack.Pop()
		newStack.Push(value)
	}

	return newStack, nil
}

func (stack *Stack) Print() error {
	if stack.head == nil {
		fmt.Print("{}")
		return errors.New("stack is empty")
	}

	fmt.Print("{")
	curr := stack.head
	for curr != nil {
		fmt.Printf("%v", curr.data)
		if curr.next != nil {
			fmt.Print(", ")
		}
		curr = curr.next
	}
	fmt.Print("}")

	return nil
}

func (stack *Stack) Println() error {
	if stack.head == nil {
		fmt.Println("{}")
		return errors.New("stack is empty")
	}

	fmt.Print("{")
	curr := stack.head
	for curr != nil {
		fmt.Printf("%v", curr.data)
		if curr.next != nil {
			fmt.Print(", ")
		}
		curr = curr.next
	}
	fmt.Println("}")

	return nil
}
