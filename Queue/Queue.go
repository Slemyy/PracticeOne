package Queue

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Request struct {
	ID    int
	Title string
}

type Task struct {
	ID   int
	Name string
}

type Event struct {
	ID          int
	Name        string
	Description string
}

type CyberData struct {
	ID       int
	IsSecure bool // Признак безопасности
}

type node struct {
	data interface{}
	next *node
}

type Queue struct {
	head *node
	tail *node
}

func Create() *Queue {
	return &Queue{}
}

// Enqueue Добавление значения в очередь.
func (queue *Queue) Enqueue(value interface{}) {
	newNode := &node{data: value}
	if queue.head == nil {
		queue.head = newNode
		queue.tail = newNode
	} else {
		queue.tail.next = newNode
		queue.tail = newNode
	}
}

// Dequeue Удаление значения из очереди.
func (queue *Queue) Dequeue() (interface{}, error) {
	if queue.head == nil {
		return nil, errors.New("queue is empty")
	} else {
		value := queue.head.data
		queue.head = queue.head.next
		if queue.head == nil {
			queue.tail = nil
		}

		return value, nil
	}
}

func (queue *Queue) IsEmpty() bool {
	if queue.head == nil {
		return true
	} else {
		return false
	}
}

func (queue *Queue) Print() error {
	if queue.head == nil {
		return errors.New("queue is empty")
	}
	curr := queue.head
	for curr != nil {
		fmt.Printf("%v\n", curr.data)
		curr = curr.next
	}
	fmt.Println()

	return nil
}

func (event *Event) SetDescription(data string) error {
	if !(event.Description == "") {
		return errors.New("this event already has a description")
	}

	event.Description = data

	return nil
}

func GenerateCyberData(queue *Queue) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000)
	isSecure := rand.Intn(2) == 0
	queue.Enqueue(CyberData{ID: id, IsSecure: isSecure})
}

func (queue *Queue) ProcessCyberData() {
	for queue.head != nil {
		secure, err := queue.Dequeue()
		if err != nil {
			return
		}

		fmt.Printf("Обработка данных с ID %d...\n", secure.(CyberData).ID)
		if secure.(CyberData).IsSecure {
			fmt.Println("Данные безопасны")
		} else {
			fmt.Println("В данных обнаружены уязвимости!")
		}
	}
}
