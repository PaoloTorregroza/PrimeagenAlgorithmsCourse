package algorithms

import (
	"fmt"
	"math"
)

// Search algorithms
func Search() {
	fmt.Println("ARRAY SEARCH ALGORITHMS")

	testArr := []int{2, 3, 69, 420, 469, 1424, 12321}
	fmt.Println("Searching for 420 in testArr:", BinarySearch(testArr, 420))
	fmt.Println("Searcing for 12312 in testArr:", BinarySearch(testArr, 12321))
	fmt.Println("Searching for 42 in testArr:", BinarySearch(testArr, 42))

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
	l := 0
	h := len(haystack)

	for {
		m := l + (h-l)/2
		v := haystack[m]

		if needle == v {
			return true
		} else if v > needle {
			h = m
		} else if v < needle {
			l = m + 1
		}

		if l > h || l == h {
			break
		}
	}

	return false
}

func TwoCrystalBalls(breaks []bool) int {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jmpAmount
	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jmpAmount

	for j := 0; j <= jmpAmount && i < len(breaks); {
		if breaks[i] {
			return i
		}

		i++
		j++
	}

	return -1
}
