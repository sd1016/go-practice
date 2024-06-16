package main

import "fmt"

// Within the function, the type of nums is equivalent
// to []int. We can call len(nums), iterate over it with range, etc.
func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
