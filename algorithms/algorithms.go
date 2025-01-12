package algorithms

import "fmt"

// Search algorithms
func Search() {
	fmt.Println("ARRAY SEARCH ALGORITHMS")

	testArr := []int{2, 3, 1424, 12321, 222, 31, 69, 420}

	fmt.Println("Searching for 420 in testArr:", LinearSearch(testArr, 420))
	fmt.Println("Searcing for 12312 in testArr:", LinearSearch(testArr, 12312))
}

func LinearSearch(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func BinarySearch(haystack []int, needle int) bool {
	return false
}
