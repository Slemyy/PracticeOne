package Stack

import (
	"errors"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

type Operation struct {
	Type  OperationType
	Value interface{}
}

type OperationType int

const (
	Push OperationType = iota
	Pop
	Peek
)

type Stack struct {
	head          *node
	lastOperation Operation
}

func Create() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value interface{}) {
	newNode := &node{data: value}
	stack.lastOperation = Operation{Type: Push, Value: value}

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
	stack.lastOperation = Operation{Type: Pop, Value: value}

	return value, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.head == nil {
		return nil, errors.New("stack is empty")
	}

	stack.lastOperation = Operation{Type: Peek, Value: stack.head.data}

	return stack.head.data, nil
}

func (stack *Stack) IsEmpty() bool {
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

func (stack *Stack) Redo() (interface{}, error) {
	if stack.lastOperation.Type == 0 {
		return nil, errors.New("nothing to redo")
	}

	switch stack.lastOperation.Type {
	case Push:
		value := stack.lastOperation.Value
		stack.Push(value)
		return value, nil

	case Pop:
		value, err := stack.Pop()
		if err != nil {
			return nil, err
		}
		return value, nil

	case Peek:
		value, err := stack.Peek()
		if err != nil {
			return nil, err
		}
		return value, nil
	}

	return nil, nil
}

func (stack *Stack) Undo() error {
	if stack.lastOperation.Type == 0 {
		return errors.New("nothing to undo")
	}

	switch stack.lastOperation.Type {
	case Push:
		_, err := stack.Pop()
		if err != nil {
			return err
		}
		stack.lastOperation.Type = Push

	case Pop:
		value := stack.lastOperation.Value
		stack.Push(value)
		stack.lastOperation.Type = Pop
	}

	return nil
}
