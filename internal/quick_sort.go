package internal

/*
QuickSort sorts an array using the QuickSort algorithm, which is an efficient
sorting algorithm with an average-case time complexity of O(n log n).
The function recursively partitions the array into elements less than or equal
to the pivot and elements greater than the pivot, then concatenates the sorted
partitions.

Type Parameters:

	T: The type of elements in the array, which can be either string or int.

Parameters:

	array []T: A slice of elements to be sorted.

Returns:

	[]T: A new slice containing the elements of the input array, sorted in ascending order.
*/
func QuickSort[T string | int](array []T) []T {
	if len(array) < 2 {
		return array
	}
	var lowerArray, higherArray []T

	pivotIndex := len(array) / 2
	pivot := array[pivotIndex]

	// Create a temporary array without the pivot element
	tempArray := append(array[:pivotIndex], array[pivotIndex+1:]...)
	
	// Partition the array into elements less than or equal to the pivot and
	// elements greater than the pivot
	for _, e := range tempArray {
		if e <= pivot {
			lowerArray = append(lowerArray, e)
		} else if e > pivot {
			higherArray = append(higherArray, e)
		}
	}

	// Recursively sort the lower and higher arrays and concatenate them with
	// the pivot
	sortedArray := append(QuickSort(lowerArray), pivot)
	sortedArray = append(sortedArray, QuickSort(higherArray)...)
	return sortedArray
}
