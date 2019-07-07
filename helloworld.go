package main

import ("fmt"
"math/rand"
"strings")

var a string

func main() {
	for i:=0; i<10; i++  {
		a := rand.Float32()
		fmt.Printf("%g   \n\r", a)
	}
	b := "123"
	c := "123"
	var test bool =strings.HasPrefix("123", "2")
	fmt.Println( b == c)
	fmt.Println(test)

	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	var str1 string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str1)
	fmt.Printf("%d\n", strings.Count(str1, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	var origS string = "Hi there! "
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)

	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)

	str3 := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str3)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str4 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str4)


	var testInt int8 = 5
	var testPtr *int8 = &testInt
	fmt.Println(*testPtr)
	fmt.Println(testPtr)


}