package Stack

import "strings"

func IsValidHooks(line string) bool {
	stack := Stack{}

	openingBrackets := "({["
	closingBrackets := ")}]"

	for _, char := range line {
		charStr := string(char)
		if strings.Contains(openingBrackets, charStr) {
			stack.Push(charStr)
		} else if strings.Contains(closingBrackets, charStr) {
			value, err := stack.Pop()
			if err != nil {
				return false // Не хватает открывающей скобки
			}
			if (charStr == ")" && value != "(") ||
				(charStr == "}" && value != "{") ||
				(charStr == "]" && value != "[") {
				return false // Скобки не совпадают
			}
		}
	}

	return stack.head == nil // Если стек пуст, то скобки правильно вложены
}
