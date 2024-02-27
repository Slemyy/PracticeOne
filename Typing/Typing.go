package Typing

import "fmt"

func FirstTask() {
	dict := make(map[string][]int)

	dict["nstu.com"] = []int{404, 0, 202}
	dict["google.com"] = []int{200, 200, 200}
	dict["vk.com"] = []int{200, 404, 200}

	for key, value := range dict {
		fmt.Printf("Сайт: %s, Результат запросов: %v\n", key, value)
	}
}

func SpamStringer(num int, str string) string {
	result := ""

	for i := 0; i < num; i++ {
		result += str
	}

	return result
}
