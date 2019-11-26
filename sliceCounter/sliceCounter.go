package main

import "fmt"

func main() {
	sample := []int{2, 3, 5, 7, 11, 13, 100, 100, 100, 100, 100}
	needle := 100
	fmt.Println(CountSlices(sample, needle))

}

// CountSlices count how many times the needle value is found in the haystack slice
func CountSlices(haystack []int, needle int) int {
	needleCount := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			needleCount++
		}
	}

	return needleCount
}