package main

import (
	"PracticeOne/HashTable"
	"PracticeOne/Queue"
	"PracticeOne/Recursion"
	"PracticeOne/Stack"
	"PracticeOne/Typing"
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	/*
		STACK
		ОБРАТНАЯ ПОЛЬСКАЯ ЗАПИСЬ + СКОБКИ + ХРАНЕНИЕ ИНФОРМАЦИИ
		(подсчет переферии, которую нужно докупить для отдела кибербезопасности)
	*/

	fmt.Print("\n\t=== STACK ===")

	//infixExpression := "5*(2+10)"
	infixExpression := "-5+(6+10-4)/(1+1*2)+1"
	if !Stack.IsValidHooks(infixExpression) {
		fmt.Println("Ошибка: неверно расставлены скобки")
		return
	}

	postfixExpression := Stack.InfixToPostfix(infixExpression)

	fmt.Println("\nEnter an expression to calculate the number of vulnerabilities:", infixExpression)
	fmt.Print("Postfix Expression: ")
	err := postfixExpression.Print()
	if err != nil {
		return
	}

	fmt.Println()

	fmt.Println("Infix Result:", -5+(6+10-4)/(1+1*2)+1)

	result := Stack.PostfixCount(postfixExpression)
	fmt.Print("Postfix Result: ")
	err = result.Print()
	if err != nil {
		return
	}

	/*
		АНАЛИЗ ПАРОЛЯ
		(используем стек для хранения информации о вызове функций и т.п.)
	*/

	fmt.Println()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nВведите пароль для анализа: ")
	password, _ := reader.ReadString('\n')
	password = password[:len(password)-1] // Удаление символа новой строки

	// Вызов функции анализа безопасности пароля
	result = Stack.AnalyzePassword(password)

	fmt.Println("=== Результаты анализа безопасности пароля ===")
	for !result.IsEmpty() {
		v, _ := result.Pop()
		fmt.Println(v.(string))
	}

	/*
		СКОБКИ
		(помощь в написании отчетов по киберугрозам)
	*/

	badText := "(Введение)):\n[[Киберугрозы] продолжают оставаться одним из наиболее значимых вызовов для организаций " +
		"и частных лиц по всему миру. В последнее время наблюдается увеличение сложности и разнообразия таких " +
		"угроз, что требует постоянного (внимания) и адаптации со стороны защищаемых систем.\n\n" +
		"(Заключение):\nУчитывая постоянное развитие киберугроз и методов их противодействия, важно оставаться " +
		"бдительным, актуальным и предпринимать соответствующие меры для защиты информации и систем от " +
		"потенциальных атак."

	normalText := "(Введение):\n[Киберугрозы] продолжают оставаться одним из наиболее значимых вызовов для " +
		"организаций и частных лиц по всему миру. В последнее время наблюдается увеличение сложности и разнообразия  " +
		"таких угроз, что требует постоянного (внимания) и адаптации со стороны защищаемых систем.\n\n" +
		"(Заключение):\n[Учитывая] постоянное развитие киберугроз и методов их противодействия, важно оставаться " +
		"бдительным, актуальным и предпринимать соответствующие меры для защиты информации и систем от " +
		"потенциальных атак."

	fmt.Println()
	fmt.Println("Первый текст:", Stack.IsValidHooks(badText))
	fmt.Println("Второй текст:", Stack.IsValidHooks(normalText))

	BST := &Stack.BinarySearchTree{}

	BST.Insert(6)
	BST.Insert(5)
	BST.Insert(3)
	BST.Insert(16)
	BST.Insert(14)
	BST.Insert(13)
	BST.Insert(15)

	fmt.Println()
	BST.DepthFirstSearch()
	fmt.Println()
	BST.Print()

	fmt.Println()
	passwordManager := Stack.Create()
	passwordManager.Push("edfgerdg")
	passwordManager.Push("354534")
	passwordManager.Push("AD534")
	passwordManager.Push("&$)@#_)_@_)")
	passwordManager.Println()

	delPassword, _ := passwordManager.Pop()
	fmt.Println("Удаленный пароль:", delPassword)
	passwordManager.Println()

	redo, _ := passwordManager.Redo()
	fmt.Println("Повторное удаление пароля:", redo)
	passwordManager.Println()

	passwordManager.Undo()
	fmt.Print("Удаленный пароль востановлен: ")
	passwordManager.Println()

	redo, _ = passwordManager.Redo()
	fmt.Println("Повторное удаление пароля:", redo)
	passwordManager.Println()

	/*
			QUEUE
		(1) Первое задание
	*/

	fmt.Println("\n\t=== QUEUE ===")
	requests := Queue.Create()

	reqOne := Queue.Request{ID: 0, Title: "Проверка безопасности сайта."}
	reqTwo := Queue.Request{ID: 1, Title: "Разработка программного обеспечения."}
	reqThree := Queue.Request{ID: 2, Title: "Написание документации по антивирусной программе."}

	requests.Enqueue(reqOne)
	requests.Enqueue(reqTwo)
	requests.Enqueue(reqThree)

	for !requests.IsEmpty() {
		request, _ := requests.Dequeue()
		fmt.Printf("Обрабатывается заявка %d: %s\n", request.(Queue.Request).ID, request.(Queue.Request).Title)
	}

	/*
			QUEUE
		(2) Второе задание
	*/

	tasks := Queue.Create()

	taskOne := Queue.Task{ID: 0, Name: "Доделать практику по программированию."}
	taskTwo := Queue.Task{ID: 1, Name: "Завершить пинтест сайта заказчика."}
	taskThree := Queue.Task{ID: 2, Name: "Разработать программу для подбора паролей."}

	fmt.Println()
	tasks.Enqueue(taskOne)
	tasks.Enqueue(taskTwo)
	tasks.Enqueue(taskThree)

	for !tasks.IsEmpty() {
		task, _ := tasks.Dequeue()
		fmt.Printf("Выполнение задачи %d: %s\n", task.(Queue.Task).ID, task.(Queue.Task).Name)
	}

	/*
			QUEUE
		(3) Третье задание
	*/

	fmt.Println()
	queue := Queue.Create()

	tasksSec := []Queue.Task{
		{ID: 1, Name: "Сканировать на наличие вредоносного ПО"},
		{ID: 2, Name: "Обновить определения антивируса"},
		{ID: 3, Name: "Мониторинг сетевого трафика"},
		{ID: 4, Name: "Анализ системных журналов"},
		{ID: 5, Name: "Исправьте уязвимости программного обеспечения"},
		{ID: 6, Name: "Тестовое задание"},
		{ID: 7, Name: "Проверочное задание"},
	}

	var wg sync.WaitGroup
	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Queue.SecurityWorker(i, queue, &wg)
	}

	for _, task := range tasksSec {
		queue.Enqueue(task)
	}

	wg.Wait()

	/*
			QUEUE
		(4) Четвертое задание
	*/

	events := Queue.Create()

	eventOne := Queue.Event{ID: 0, Name: "Обнаружено вторжение"}
	eventTwo := Queue.Event{ID: 1, Name: "Атака на веб-приложение"}
	eventThree := Queue.Event{ID: 2, Name: "Утечка конфиденциальных данных"}

	err = eventOne.SetDescription("Система обнаружила аномальную активность, указывающую на возможное " +
		"вторжение в корпоративную сеть. Это может включать в себя обнаружение несанкционированных попыток доступа " +
		"к системным ресурсам, необычные сетевые запросы или другие аномалии в поведении сетевого трафика.")
	if err != nil {
		return
	}

	err = eventTwo.SetDescription("Система зафиксировала серию подозрительных запросов к веб-приложению, " +
		"которые могут указывать на попытку эксплойтации уязвимостей в программном обеспечении. Атаки могут включать " +
		"в себя SQL-инъекции, кросс-сайтовый скриптинг (XSS) или попытки некорректной аутентификации.")
	if err != nil {
		return
	}

	err = eventThree.SetDescription("Система обнаружила подозрительную активность, свидетельствующую о " +
		"возможной утечке конфиденциальных данных. Это может включать в себя попытки несанкционированного доступа к " +
		"базам данных, передачу конфиденциальной информации через нестандартные каналы связи или наблюдение за " +
		"нормальным потоком данных с целью кражи информации.")
	if err != nil {
		return
	}

	fmt.Println()
	events.Enqueue(eventOne)
	events.Enqueue(eventTwo)
	events.Enqueue(eventThree)

	for !events.IsEmpty() {
		event, _ := events.Dequeue()
		fmt.Printf("Описание евента %d: %s (%s)\n", event.(Queue.Event).ID, event.(Queue.Event).Description,
			event.(Queue.Event).Name)
	}

	/*
			QUEUE
		(5) Пятое задание
	*/
	fmt.Println()
	data := Queue.Create()
	Queue.GenerateCyberData(data)
	time.Sleep(1 * time.Second)
	Queue.GenerateCyberData(data)
	time.Sleep(1 * time.Second)
	Queue.GenerateCyberData(data)
	data.Print()

	data.ProcessCyberData()

	/*
		HASHTABLE
	*/

	fmt.Println("\n=== HASHTABLE ===")
	// Создаем новый экземпляр ассоциативного массива
	roles := HashTable.Create()

	// Вставляем роли доступа для пользователей
	err = roles.Insert("Slemyy", "admin")
	if err != nil {
		return
	}
	err = roles.Insert("Dister", "editor")
	if err != nil {
		return
	}
	err = roles.Insert("Holy", "viewer")
	if err != nil {
		return
	}
	err = roles.Insert("az", "viewer")
	if err != nil {
		return
	}
	err = roles.Insert("za", "viewer")
	if err != nil {
		return
	}

	// Получаем роли доступа для пользователей
	foundSlemyy, err := roles.Get("Slemyy")
	fmt.Println("Роль пользователя \"Slemyy\":", foundSlemyy)

	foundDister, err := roles.Get("Dister")
	fmt.Println("Роль пользователя \"Dister\":", foundDister)

	foundHoly, err := roles.Get("Holy")
	fmt.Println("Роль пользователя \"Holy\":", foundHoly)

	attackTypes := []string{"DDoS", "Phishing", "Malware", "DDoS", "Phishing", "Phishing", "Ransomware", "DDoS"}
	attackFrequency := HashTable.CountFrequency(attackTypes)

	fmt.Println("\nЧастота типов атак в кибербезопасности:")
	err = attackFrequency.Println()
	if err != nil {
		return
	}

	BD := HashTable.Create()
	fmt.Println()

	// Примеры добавления аутентификаций
	err = BD.Insert("Slemyy", HashTable.User{Password: "98761", Group: "admin"})
	if err != nil {
		fmt.Println("Error insert user:", err)
	}

	err = BD.Insert("Dister", HashTable.User{Password: "1338", Group: "user"})
	if err != nil {
		fmt.Println("Error insert user:", err)
	}

	// Проверка аутентификации
	isAuthenticated, err := BD.Authenticate("Slemyy", "98761")
	if err != nil {
		fmt.Println("Auth error:", err)
	} else {
		fmt.Println("Результат авторизации (Slemyy):", isAuthenticated)
	}

	isAuthenticated, err = BD.Authenticate("Dister", "wrongpassword")
	if err != nil {
		fmt.Println("Auth error:", err)
	} else {
		fmt.Println("Результат авторизации (Dister):", isAuthenticated)
	}

	isAuthenticated, err = BD.Authenticate("Dister", "1338")
	if err != nil {
		fmt.Println("Auth error:", err)
	} else {
		fmt.Println("Результат авторизации (Dister):", isAuthenticated)
	}

	/*
		РЕКУРСИЯ
	*/

	pass := []rune("1234") // пароль
	var results [][]rune
	Recursion.BruteForcePasswords(pass, 0, &results)
	fmt.Println("\nКоличество перестановок:", len(results))

	for _, permutation := range results {
		fmt.Println(string(permutation))
	}

	/*
		ТИПИЗАЦИЯ
	*/

	fmt.Println()
	Typing.FirstTask()

	fmt.Println()
	spam := Typing.SpamStringer(5, "Spam")
	fmt.Println(spam)
}
