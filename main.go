package main

import (
	"fmt"
	time2 "time"
)

func main() {

	//variable declaration
	var i int = 10
	var f float64 = 1.0
	var b bool = true
	fmt.Println(i, f, b)
	//type inference
	ii := 20
	fmt.Println(ii)

	//if-elseif-else
	day := time2.Now().Weekday()
	if day == time2.Monday {
		println("It's Monday")
	} else if day == time2.Tuesday {
		println("It's Tuesday")
	} else {
		println("It's other day")
	}

	//switch
	switch day {
	case time2.Monday:
		println("It's Monday")
	case time2.Tuesday:
		println("It's Tuesday")
	case time2.Wednesday:
		println("It's Wednesday")
	default:
		println("It's other day")
	}

	//looping
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//array
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	println(arr[0], arr[1])

	//slices
	primes := [5]int{1, 3, 5, 7, 11}

	var s []int = primes[1:4]
	fmt.Println(s)

	//dynamic array using slice with make

	sl := make([]int, 0, 5)
	println(len(sl), cap(sl))
	sl = append(sl, 10)
	sl = append(sl, 20)
	fmt.Println(len(sl), cap(sl), sl)

	sl = append(sl, 30)
	sl = append(sl, 40)
	sl = append(sl, 50)
	sl = append(sl, 60)
	fmt.Println(len(sl), cap(sl), sl)

	//range with loop for slice/array

	for i, item := range sl {
		fmt.Println(i, item)
	}

	//struct
	type Point struct {
		x int
		y int
	}
	p := Point{10, 20}
	fmt.Println(p)

	//maps
	m := make(map[int]string)
	m[0] = "Zero"
	m[1] = "One"
	m[2] = "Two"
	fmt.Println(m)

	for k, v := range m {
		fmt.Print(k, "->", v, " ")
	}

	item, ok := m[4]
	fmt.Println("\n", item, ok)

	//function call

	var x = 20
	var y = 10
	fmt.Println("x=", x, "y=", y)
	//pass value by reference/pointer
	swap(&x, &y)
	fmt.Println("x=", x, "y=", y)
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
