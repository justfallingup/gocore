package fibo

//Num вычисляет число Фибоначчи для неотрицательного n
func Num(n int) int {
	if n == 0 {
		return 0
	}
	x1, x2 := 0, 1
	for i := 1; i < n; i++ {
		x1, x2 = x2, x1 + x2
	}
	return x2
}
