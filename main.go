// part 1
package main

import (
	"fmt"
	"os"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}
func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("prime check requires number >= 2")
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, fmt.Errorf("negative exponents not supported")
	}
	if exponent == 0 && base == 0 {
		return 0, fmt.Errorf("MATH ERROR!! 0^0 is undefined")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

// END-PART 1
// PART 2

func makeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func makeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial

	add = func(x int) {
		total += x
	}

	subtract = func(x int) {
		total -= x
	}

	get = func() int {
		return total
	}

	return
}

// END-PART2
// PART 3
func Apply(nums []int, operation func(int) int) []int {
	returned := make([]int, len(nums))
	for i, v := range nums {
		returned[i] = operation(v)
	}
	return returned
}

func Filter(nums []int, predicate func(int) bool) []int {
	var filtered []int
	for _, v := range nums {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	result := initial
	for _, v := range nums {
		result = operation(result, v)
	}
	return result
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// PART3 -END
// part 4

func ExploreProcess() {
	fmt.Printf("Current Process ID: %d\n", os.Getpid())
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())

	data := []int{1, 2, 3, 4}
	fmt.Printf("Memory address of slice: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])
	fmt.Println("Note: Other processes can't see these memory addresses due to process isolation.")

	/*  
		* A process ID is a unique identifier assigned by the operating system kernel to distinguish each running process.
		* Process isolation is essential because it stops one process from accessing or 
			altering another process’s memory without authorization, which improves system security and stability.
		* The key difference between a slice header address and element addresses is that the slice header is stored 
			directly in memory, whereas the elements reside in a contiguous memory block that the header references. 
	*/
}

func main() {
	ExploreProcess()
}

// part 4- end