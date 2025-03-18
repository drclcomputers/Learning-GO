//Just a basic calculator. What I've learnt: functions, nested loops, recursion, working with variables, reading user input

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func min(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func power(a, b int) int {
	if b == 0 {
		return a
	}
	if b%2 == 1 {
		return a * power(a, b-1)
	}
	return power(a, b/2)
}

func main() {
	a := 0
	b := 0
	semn := 0
	for semn != -1 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("1. Add\n2. Minus\n3. Multiplication\n4. Division\n5. Power\n-1. Exit")
		fmt.Scanln(&semn)
		if semn == -1 {
			break
		}
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		if semn == 1 {
			fmt.Println(add(a, b))
		} else if semn == 2 {
			fmt.Println(min(a, b))
		} else if semn == 3 {
			fmt.Println(mult(a, b))
		} else if semn == 4 {
			fmt.Println(div(a, b))
		} else if semn == 5 {
			fmt.Println(power(a, b))
		} else {
			fmt.Println("Unknown sign!")
		}
		fmt.Println("Press enter to continue...")
		fmt.Scanln()
	}
}
