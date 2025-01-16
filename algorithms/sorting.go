package algorithms

// We will iterate over an array, comparing each element with the next one, swapping them if the next one
// is smaller than the previous one, this ensures the last element in the array is the largest one after each iteration.
// That means, after each iteration N is going to be lowet (n, n-1, n-2, n-3, etc...), There is a formula for calculating the
// sum of a range of numvers (n+1)n/2. If we factorize that into O notation we get:
//
// n(n+1)(1/2) = 0
// Drop contstants:
// n(n+1)
// n^2 + n
// Drop insignificant terms:
// O(n^2)
func BubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
