package main

import (
	"fmt"
	"grokking_algorithms/internal"
)

func main() {
	// Binary Search Example
	runBinarySearch()

	// Selection Sort Example
	runSelectionSort()

	// Recursive Example (Factorial)
	runFactorial()

	// Quick Sort Example
	runQuickSort()

	// Breadth First Search Example
	runBreadthFirstSearch()

	// Dijkstra's Algorithm Example
	runDijkstra()
}

// runBinarySearch demonstrates the binary search algorithm
func runBinarySearch() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	item := 1
	itemIndex := internal.BinarySearch(array, item)
	fmt.Println("Binary Search Result:", itemIndex)
}

// runSelectionSort demonstrates the selection sort algorithm
func runSelectionSort() {
	array := []string{"pidgeotto", "squirtle", "charmander", "bulbasaur", "pikachu", "snorlax"}
	sortedArray := internal.SelectionSort(array)
	fmt.Println("Selection Sort Result:", sortedArray)
}

// runFactorial demonstrates the recursive factorial calculation
func runFactorial() {
	number := 5
	factorial := internal.Factorial(number)
	fmt.Println("Factorial Result:", factorial)
}

// runQuickSort demonstrates the quick sort algorithm
func runQuickSort() {
	array := []string{"N", "T", "A", "G", "L", "C", "F"}
	sortedArray := internal.QuickSort(array)
	fmt.Println("Quick Sort Result:", sortedArray)
}

// runBreadthFirstSearch demonstrates the breadth-first search algorithm
func runBreadthFirstSearch() {
	relationships := map[string][]string{
		"Alpha":   {"Bravo", "Charlie", "Delta"},
		"Bravo":   {"Echo", "Charlie", "Foxtrot"},
		"Charlie": {"Bravo", "Hotel", "India"},
		"Delta":   {"Juliet", "Kilo", "Lima"},
		"Mike":    {"November", "Oscar", "Papa"},
	}
	start := "Alpha"
	elementToFind := "Mike"
	result := internal.BreadthFirstSearch(relationships, start, elementToFind)
	fmt.Println("Breadth First Search Result:", result)
}

// runDijkstra demonstrates Dijkstra's algorithm
func runDijkstra() {
	network := map[string]map[string]uint{
		"A": {"B": 2, "C": 1},
		"B": {"D": 4, "E": 1},
		"C": {"E": 3},
		"D": {"F": 3},
		"E": {"F": 5},
		"F": {},
	}
	result := internal.Dijkstra(network, "A", "F")
	fmt.Println("Dijkstra Result:", result)
}
