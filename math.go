package main

func add(x int, y int) int {

	return x + y
}

func prod(x int, y int) int {

	return x * y
}

func fact(x int) int {

	if x > 0 {
		return x * fact(x-1)
	}
	return 1
}
