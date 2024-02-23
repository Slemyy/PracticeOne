package Stack

import "fmt"

func InfixToPostfix(infix string) *Stack {
	outputStack := Create()    // Создаем стек выхода
	operationStack := Create() // Создаем стек операций
	infix = "(" + infix + ")"

	operators := map[rune]bool{
		'+': true,
		'-': true,
		'*': true,
		'/': true,
	}

	for i := 0; i < len(infix); i++ {
		char := rune(infix[i])
		switch {
		case char == '(':
			operationStack.Push(char)
			break

		case char == ')':
			for pop, _ := operationStack.Pop(); pop != '('; {
				outputStack.Push(string(pop.(int32)))
				pop, _ = operationStack.Pop()
			}
			break

		case operators[char]:
			if operationStack.isEmpty() {
				operationStack.Push(char)
			} else {
				peek, err := operationStack.Peek()
				if err != nil {
					return nil
				}

				for operationPriority(peek.(rune)) >= operationPriority(char) {
					pop, err := operationStack.Pop()
					if err != nil {
						return nil
					}

					outputStack.Push(string(pop.(int32)))
					peek, err = operationStack.Peek()
					if err != nil {
						return nil
					}
				}

				operationStack.Push(char)
			}
			break

		case char >= '0' && char <= '9':
			num := char - '0'
			for j := i + 1; j < len(infix); j++ {
				nextChar := rune(infix[j])
				if nextChar >= '0' && nextChar <= '9' {
					num = num*10 + (nextChar - '0')
					i = j
				} else {
					outputStack.Push(num)
					break
				}
			}

			if i == len(infix)-1 {
				outputStack.Push(num)
			}

			break
		}
	}

	reverse, _ := outputStack.Reverse()
	return reverse
}

func operationPriority(char rune) int {
	switch char {
	case '+', '-':
		return 1

	case '*', '/':
		return 2
	}

	return 0
}

func PostfixCount(postfixCount *Stack) *Stack {
	result := Create()

	_ = postfixCount.Print()
	_ = result.Print()

	operations := map[string]func(int32, int32) int32{
		"+": func(a, b int32) int32 { return a + b },
		"-": func(a, b int32) int32 { return a - b },
		"*": func(a, b int32) int32 { return a * b },
		"/": func(a, b int32) int32 { return a / b },
	}

	for !postfixCount.isEmpty() {
		peek, err := postfixCount.Peek()
		if err != nil {
			return nil
		}

		switch peek.(type) {
		case int32:
			value, _ := postfixCount.Pop()
			peek, _ = postfixCount.Peek()
			if fmt.Sprintf("%v", peek) == "-" && result.isEmpty() {
				value = value.(int32) * -1
				_, _ = postfixCount.Pop()
			}

			result.Push(value)
			break

		case string:
			operandTwo, _ := result.Pop() // Достаем из стека второе число
			operandOne, _ := result.Pop() // Достаем из стека первое число

			operation := operations[peek.(string)]
			if operandOneFloat, ok := operandOne.(int32); ok {
				if operandTwoFloat, ok := operandTwo.(int32); ok {
					result.Push(operation(operandOneFloat, operandTwoFloat))
				}
			}

			_, err = postfixCount.Pop()
			if err != nil {
				return nil
			}
		}

		fmt.Println()
		_ = postfixCount.Print()
		_ = result.Print()
	}

	fmt.Println()
	return result
}
