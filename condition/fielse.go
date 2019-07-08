package main

import (
	"fmt"
	"strings"
)

func main() {
	var flag bool = false;
	if flag {
		print("True")
	} else {
		print("False")
	}

	var first int = 10
	var cond int

	if first <= 0 {

		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {

		fmt.Printf("first is between 0 and 5\n")
	} else {

		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {

		fmt.Printf("cond is greater than 10\n")
	} else {

		fmt.Printf("cond is not greater than 10\n")
	}
	fmt.Println(cond)

	//var orig string = "ABC"
	//// var an int
	//var newS string
	//// var err error
	//
	//fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	//// anInt, err = strconv.Atoi(origStr)
	//an, err := strconv.Atoi(orig)
	//if err != nil {
	//	fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
	//	return
	//}
	//fmt.Printf("The integer is %d\n", an)
	//an = an + 5
	//newS = strconv.Itoa(an)
	//fmt.Printf("The new string is: %s\n", newS)

	for ix := 1; ix < 10; ix++ {
		fmt.Println(strings.Repeat("G", ix))
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	aaaaa()
	f()
		fmt.Println()
	b()

	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
}

func aaaaa() {
	i := 5
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}