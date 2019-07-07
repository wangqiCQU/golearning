package main

import (
	"fmt"
	"time"
)

func main()  {
	//var i int = 1
	//g := adder()
	//for i<10{
	//	i += 1
	//	fmt.Println(g(i))
	//}

	result := 0
	start := time.Now()
	for i := 0; i <= 25; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

//func adder() func(a int) int{
//	var x int;
//	return func(a int) int {
//		x += a
//		return x
//	}
//}


func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
