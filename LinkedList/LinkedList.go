package LinkedList

import (
	"errors"
	"fmt"
	"time"
)

// BlockedIP структура заблокированных айпи адрессов.
type BlockedIP struct {
	IPAddress string
	Reason    string
	BlockedAt time.Time
}

type node struct {
	data BlockedIP
	next *node
	prev *node
}

type LinkedList struct {
	head   *node
	tail   *node
	length int
}

func Create() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) AddToFront(value BlockedIP) {
	newNode := &node{data: value}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}

	list.length++
}

func (list *LinkedList) AddToBack(value BlockedIP) {
	newNode := &node{data: value}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}

	list.length++
}

func (list *LinkedList) DeleteAtIndex(index int) error {
	if list.head == nil {
		return errors.New("the list is empty")
	}

	if index < 0 || index > list.length {
		return errors.New("index out of range")
	}

	if index == 0 {
		list.head = list.head.next

		if list.head == nil {
			list.tail = nil
		} else {
			list.head.prev = nil
		}
	} else if index == list.length-1 {
		list.tail = list.tail.prev
		list.tail.next = nil
	} else if index > list.length/2 {
		curr := list.tail
		for i := list.length; i > index; i-- {
			curr = curr.prev
		}

		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	} else {
		curr := list.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}

		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}

	list.length--

	return nil
}

func (list *LinkedList) DeleteFromFront() error {
	err := list.DeleteAtIndex(0)
	if err != nil {
		return err
	}

	return nil
}

func (list *LinkedList) Search(value BlockedIP) bool {
	curr := list.head
	for curr != nil {
		if curr.data == value {
			return true
		}
		curr = curr.next
	}

	return false
}

func (list *LinkedList) Print() error {
	if list.length == 0 {
		return errors.New("the list is empty")
	}

	fmt.Println("\t=== Список заблокированных айпи адресов ===\t")

	curr := list.head
	for i := 0; curr != nil; i++ {
		fmt.Printf("%d) %s (%s) | %s\n", i, curr.data.IPAddress, curr.data.Reason,
			curr.data.BlockedAt.Format("2006-01-02 15:04:05"))

		curr = curr.next
	}
	fmt.Println()

	return nil
}

func (list *LinkedList) GetValue(index int) (BlockedIP, error) {
	if list.head == nil {
		return BlockedIP{}, errors.New("the list is empty")
	}

	if index < 0 || index > list.length {
		return BlockedIP{}, errors.New("index out of range")
	}

	if index == 0 {
		return list.head.data, nil
	} else if index == list.length-1 {
		return list.tail.data, nil
	} else if index > list.length/2 {
		curr := list.tail
		for i := list.length; i > index; i-- {
			curr = curr.prev
		}

		return curr.data, nil
	} else {
		curr := list.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}

		return curr.data, nil
	}
}
