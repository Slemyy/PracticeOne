package Stack

func AnalyzePassword(password string) *Stack {
	resultStack := &Stack{}

	if len(password) >= 8 {
		resultStack.Push("Длина пароля: Проверено")
	} else {
		resultStack.Push("Длина пароля: Слишком короткий (минимум 8 символов)")
	}

	hasDigits := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigits = true
			break
		}
	}
	if hasDigits {
		resultStack.Push("Наличие цифр: Проверено")
	} else {
		resultStack.Push("Наличие цифр: Отсутствует")
	}

	hasUpper := false
	for _, char := range password {
		if char >= 'A' && char <= 'Z' || char >= 'А' && char <= 'Я' {
			hasUpper = true
			break
		}
	}
	if hasUpper {
		resultStack.Push("Наличие символов верхнего регистра: Проверено")
	} else {
		resultStack.Push("Наличие символов верхнего регистра: Отсутствует")
	}

	hasLower := false
	for _, char := range password {
		if char >= 'a' && char <= 'z' || char >= 'а' && char <= 'я' {
			hasLower = true
			break
		}
	}
	if hasLower {
		resultStack.Push("Наличие символов нижнего регистра: Проверено")
	} else {
		resultStack.Push("Наличие символов нижнего регистра: Отсутствует")
	}

	return resultStack
}
