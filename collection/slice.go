package main

import "fmt"

func f(a [3]int) { fmt.Println(&a) }
func fp(a *[3]int) { fmt.Println(a) }

func main()  {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	var ar [3]int
	fmt.Println(&ar)
	f(ar) 	// passes a copy of ar
	fp(&ar) // passes a pointer to ar

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range


	var  test []int = []int{2, 3, 5, 7, 11, 13,15,17,19,21}
	slice12 := test[2:4]
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))
	slice13 := test[2:5]
	fmt.Println(len(slice13))
	fmt.Println(cap(slice13))

	items := test[2:4]
	for _, item := range items {
		item *= 2
	}

	for _ , value := range items {
		fmt.Println(value)
	}

	slice14 := test[2:2]
	fmt.Println(slice14)
}
