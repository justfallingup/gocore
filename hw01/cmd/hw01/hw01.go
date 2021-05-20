package main

import (
	"fmt"

	"github.com/justfallingup/gocore/hw01/pkg/Fibonacci"
)

func main()  {
	n := []int{-3, 0, 1, 2, 5, 12, 15, 27}
	for _, i := range n {
		if i > 20 {
			fmt.Println("Выберите число меньше 20")
			continue
		}
		if i < 0 {
			fmt.Println("Выберите положительное число меньше 20")
			continue
		}
		f := Fibonacci.Number(i)
		fmt.Printf("Число Фибоначчи номер %d равно %d\n", i, f)
	}
}