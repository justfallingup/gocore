package main

import (
	"fmt"
	"github.com/justfallingup/gocore/hw01-fibo/pkg"
)

func main()  {
	n := []int{-3, 0, 1, 7, 15, 27}
	for _, i := range n {
		if i < 0 {
			fmt.Printf("Передано отрицательное число %d\n", i)
			continue
		}
		if i > 20 {
			fmt.Printf("Передано %d, нужно число меньше 20\n", i)
			continue
		}
		fmt.Printf("Число Фибоначчи от %d равно %d\n", i, fibo.Num(i))
	}
}
