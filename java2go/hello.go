package collection_test

import (
	"/home/baystars/go/src/realworldhttp/java2go/collection_test"
	"fmt"
)

func ExampleStack() {
	var s collection_test.Stack
	s.Push("world!")
	s.Push("Hello, ")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()
}
