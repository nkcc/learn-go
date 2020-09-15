package hello

import "fmt"

func init() {
	fmt.Println("test")
	var (
		// arr0        [5]int    = [5]int{1, 2, 3}
		// arr1                  = [5]int{1, 2, 3, 4, 5}
		// stringArray [5]string = [5]string{1: "abcde"}
		nolen = [...]int{1, 2, 3, 4, 5, 6}
	)

	const (
		arr = iota
		n1
		n2
		n4
	)

	for k1, v1 := range nolen {
		fmt.Printf("(%d, %d)", k1, v1)
		fmt.Println()
	}

	fmt.Println(arr, n1, n2)
}
