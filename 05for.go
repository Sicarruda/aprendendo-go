package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for sem uma condição será repetido várias vezes até que você dê um break no loop ou return da função de fechamento.
	// o for abaixo geraria um loop caso não ouvesse um break
	for {
		fmt.Println("loop")
		break
	}

	for i := 0; i < 10; i++ {
		if i > 5 {
			var a string = "lala"
			fmt.Println(a)
			return
		}
		fmt.Println(i)
	}

}
