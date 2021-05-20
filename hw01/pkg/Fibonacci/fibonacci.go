package Fibonacci

func Number(n int)  int {
	f := 0
	if n >= 2 {
		x := 0
		y := 1
		for i := 0; i < n; i++ {
			f = x + y
			x = y
			y = f
		}
	}
	return f
}
