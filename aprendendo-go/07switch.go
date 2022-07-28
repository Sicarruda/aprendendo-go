package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("escreva ", i, " como ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("é final de semana")
	default:
		fmt.Println("é dia de semana")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("é antes do meio dia")
	default:
		fmt.Println("é depois do meio dia")
	}
}
