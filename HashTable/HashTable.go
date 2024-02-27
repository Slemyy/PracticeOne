package HashTable

import (
	"errors"
	"fmt"
)

const arraySize = 769

type node struct {
	key   string
	value interface{}
	next  *node
}

type Hashtable struct {
	table [arraySize]*node
}

type User struct {
	Password string
	Group    string
}

func hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}

	return hash % arraySize
}

func Create() *Hashtable {
	return &Hashtable{}
}

func (HT *Hashtable) Insert(key string, value interface{}) error {
	newNode := &node{key: key, value: value}
	index := hashFunction(newNode.key)

	if HT.table[index] == nil {
		HT.table[index] = newNode
		return nil
	}

	if HT.table[index].key == key {
		return errors.New("such a key already exists")
	}

	curr := HT.table[index]
	for curr.next != nil {
		if curr.next.key == key {
			return errors.New("such a key already exists")
		}

		curr = curr.next
	}
	curr.next = newNode

	return nil
}

func (HT *Hashtable) Remove(key string) (interface{}, error) {
	index := hashFunction(key)

	if HT.table[index] == nil {
		return "", errors.New("no such meaning")
	}

	curr := HT.table[index]
	if curr.key == key {
		HT.table[index] = curr.next
		return curr.value, nil
	}

	prev := curr
	for curr != nil && curr.key != key {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		return "", nil
	}

	prev.next = curr.next
	return curr.value, nil
}

func (HT *Hashtable) Get(key string) (interface{}, error) {
	index := hashFunction(key)

	if HT.table[index] == nil {
		return "", errors.New("value not found")
	} else if HT.table[index].key == key {
		return HT.table[index].value, nil
	}

	curr := HT.table[index]

	for curr != nil {
		if curr.key == key {
			return curr.value, nil
		}
		curr = curr.next
	}

	return curr.value, nil
}

func (HT *Hashtable) Found(key string) (bool, error) {
	index := hashFunction(key)

	if HT.table[index] == nil {
		return false, errors.New("value not found")
	} else if HT.table[index].key == key {
		return true, nil
	}

	curr := HT.table[index]

	for curr != nil {
		if curr.key == key {
			return true, nil
		}
		curr = curr.next
	}

	return true, nil
}

func (HT *Hashtable) Println() error {
	for i, node := range HT.table {
		if HT.table[i] == nil {
			continue
		}
		fmt.Printf("Index %d: ", i)
		curr := node
		for curr != nil {
			switch curr.value.(type) {
			case int:
				fmt.Printf("(%s, %d) ", curr.key, curr.value.(int))
				curr = curr.next
			case string:
				fmt.Printf("(%s, %s) ", curr.key, curr.value.(string))
				curr = curr.next
			default:
				fmt.Printf("(%s, %s) ", curr.key, curr.value)
				curr = curr.next
			}
		}
		fmt.Println()
	}

	return nil
}

func CountFrequency(elements []string) *Hashtable {
	freq := Create()

	for _, element := range elements {
		if exists, _ := freq.Found(element); exists {
			// Если элемент уже встречался, увеличиваем его частоту на 1
			el, _ := freq.Remove(element)
			err := freq.Insert(element, el.(int)+1)
			if err != nil {
				return nil
			}
		} else {
			// Если элемент встречается впервые, устанавливаем его частоту в 1
			err := freq.Insert(element, 1)
			if err != nil {
				return nil
			}
		}
	}

	return freq
}

func (HT *Hashtable) Authenticate(username, password string) (bool, error) {
	auth, err := HT.Get(username)
	if err != nil {
		return false, err // Пользователь не найден
	}

	user, ok := auth.(User)
	if !ok {
		return false, errors.New("ошибка авторизации")
	}

	return user.Password == password, nil
}
