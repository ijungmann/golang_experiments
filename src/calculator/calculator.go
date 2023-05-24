package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter an integer")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	result, _ := strconv.Atoi(reader.Text())
OperandLoop:
	for {
		fmt.Println("Enter an operand")
		reader.Scan()
		operand := reader.Text()
		fmt.Println("Enter an integer")
		reader.Scan()
		affector, _ := strconv.Atoi(reader.Text())

		switch operand {
		case "+":
			fmt.Println("Adding")
			result += affector
		case "-":
			fmt.Println("Subtracting")
			result -= affector
		case "*":
			fmt.Println("Multiplying")
			result *= affector
		case "/":
			fmt.Println("Dividing")
			result /= affector
		default:
			fmt.Println("Breaking")
			break OperandLoop
		}
		fmt.Printf("result is %d\n", result)
	}
}
