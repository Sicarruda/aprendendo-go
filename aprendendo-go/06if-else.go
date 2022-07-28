package main

import "fmt"

//Não há if ternário em Go, então você precisa usar uma instrução completa do if mesmo em condições básicas.

func main() {
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 é divisivel por 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "tem 1 digito")
	} else {
		fmt.Println(num, "tem varios digitos")
	}

}
