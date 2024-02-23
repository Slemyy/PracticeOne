package main

import (
	"PracticeOne/Stack"
	"fmt"
)

func main() {
	infixExpression := "-5+(6+10-4)/(1+1*2)+1"
	//infixExpression := "5*2+10"
	postfixExpression := Stack.InfixToPostfix(infixExpression)

	fmt.Println("\nInfix Expression:", infixExpression)
	fmt.Print("Postfix Expression: ")
	err := postfixExpression.Print()
	if err != nil {
		return
	}

	fmt.Println()

	fmt.Println("Infix Result:", (6+10-4)/(1+1*2)+1)

	result := Stack.PostfixCount(postfixExpression)
	fmt.Print("Postfix Result: ")
	err = result.Print()
	if err != nil {
		return
	}

}
