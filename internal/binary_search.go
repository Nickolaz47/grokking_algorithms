package internal

/*
BinarySearch performs a binary search on a sorted array to find the index
of the specified item. If the item is found, it returns the index of the item.
If the item is not found, it returns -1.

Type Parameters:

	T: The type of elements in the array, which can be either int or string.

Parameters:

	array []T: A sorted slice of elements to search in.
	item T: The element to search for.

Returns:

	int: The index of the item if found, otherwise -1.
*/
func BinarySearch[T int | string](array []T, item T) int {
	lower_limit := 0
	upper_limit := len(array) - 1

	for lower_limit <= upper_limit {
		// Selecionando o indíce do meio do array
		middle := (upper_limit + lower_limit) / 2

		if array[middle] == item {
			return middle
			// Se o palpite for maior do que o item
		} else if array[middle] > item {
			// O limite superior será o palpite menos um.
			upper_limit = middle - 1
			// Se o palpite for menor do que o item
		} else if array[middle] < item {
			lower_limit = middle + 1
		}
	}
	return -1
}
