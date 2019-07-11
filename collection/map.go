package main

import (
	"fmt"
	"go/types"
	"unsafe"
)

func main() {
	var test map[string]int = map[string]int{"12": 134, "12321": 1213}
	fmt.Println(test["12"])
	delete(test, "werew")

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)

	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

	fmt.Println()
	fmt.Println(unsafe.Sizeof(types.Int))

}
