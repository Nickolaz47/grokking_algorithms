package internal

/*
findIndex searches for the specified element in the provided array
and returns the index of the first occurrence of the element.

If the element is found, it returns its index (0-based).
If the element is not found, the function returns -1.

Type Parameters:

	T: The type of elements in the array, which can be either a string or an int.

Parameters:

	array []T: A slice of elements where the search is performed.
	element T: The element to search for in the array.

Returns:

	int: The index of the first occurrence of the element if found, otherwise -1.
*/
func findIndex[T string | int](array []T, element T) int {
	for i, e := range array {
		if e == element {
			return i
		}
	}
	return -1
}

/*
pop removes the first occurrence of the specified element from the array
and returns the modified array.

If the element is not found, the original array is returned.

Type Parameters:

	T: The type of elements in the array, which can be either a string or an int.

Parameters:

	array []T: A slice of elements where the removal is performed.
	elementToRemove T: The element to be removed from the array.

Returns:

	[]T: A new slice with the specified element removed, or the original array if not found.
*/
func pop[T string | int](array []T, elementToRemove T) []T {
	elementToRemoveIndex := findIndex(array, elementToRemove)
	if elementToRemoveIndex < 0 {
		return array
	}
	return append(array[:elementToRemoveIndex], array[elementToRemoveIndex+1:]...)
}

/*
getLowerElement returns the smallest (or lexicographically lowest) element
from the given array.
It assumes that the array is non-empty and that the elements are comparable.

Type Parameters:

	T: The type of elements in the array, which can be either a string or an int.

Parameters:

	array []T: A slice of elements from which the smallest element is found.

Returns:

	T: The smallest element in the array.
*/
func getLowerElement[T string | int](array []T) (lowerElement T) {
	lowerElement = array[0]
	for _, element := range array[1:] {
		if element < lowerElement {
			lowerElement = element
		}
	}
	return
}

/*
SelectionSort performs a selection sort on the provided array, returning a
new sorted array in ascending order.

The function repeatedly finds the smallest element in the array, appends it to the
result, and removes it from the original array until the array is empty.

Type Parameters:

	T: The type of elements in the array, which can be either a string or an int.

Parameters:

	array []T: A slice of elements to be sorted.

Returns:

	[]T: A new slice containing the elements of the input array, sorted in ascending order.
*/
func SelectionSort[T string | int](array []T) (sortedArray []T) {
	if len(array) < 2 {
		return array
	}

	for {
		if len(array) == 0 {
			return
		}

		lowerElement := getLowerElement(array)
		sortedArray = append(sortedArray, lowerElement)
		array = pop(array, lowerElement)
	}

}
