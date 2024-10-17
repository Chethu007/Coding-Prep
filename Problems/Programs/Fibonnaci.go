package main

func fib(n int) int {
	if n > 0 {
		if n == 1 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	}
	return 0
}

func main() {
	println("Res:", fib(10))
}
