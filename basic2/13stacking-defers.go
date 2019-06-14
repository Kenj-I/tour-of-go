// +build ignore

package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		//  LIFO(last-in-first-out)
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
