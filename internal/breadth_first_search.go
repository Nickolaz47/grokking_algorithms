package internal

import (
	"fmt"
	"slices"
)

/*
popFirst removes and returns the first element from a slice of strings,
modifying the original slice.

Parameters:

	queue *[]string: A pointer to a slice of strings, from which the first
	element will be removed.

Returns:

	string: The first element that was removed from the slice.
*/
func popFirst(queue *[]string) string {
	first_element := (*queue)[0]
	*queue = (*queue)[1:]
	return first_element
}

/*
BreadthFirstSearch performs a breadth-first search (BFS) on a graph represented
as an adjacency list. It searches for a specified element starting from a given
start node and checks if there is a path from the start node to the element.

Parameters:

	graph map[string][]string: A map representing the adjacency list of the
		graph, where each key is a node and its value is a slice of adjacent
		nodes.
	startNode string: The node from which the BFS will start.
	element string: The element to search for in the graph.

Returns:

	string: A formatted string indicating whether the specified element is
		related to the start node (i.e., reachable from it) or not.
*/
func BreadthFirstSearch(graph map[string][]string, startNode, element string) string {
	var searchList []string
	var searchedElements []string

	searchList = append(searchList, graph[startNode]...)

	for len(searchList) > 0 {
		elementToSearch := popFirst(&searchList)
		if !slices.Contains(searchedElements, elementToSearch) {
			if elementToSearch == element {
				return fmt.Sprintf("%s is related to %s.", element, startNode)
			} else {
				searchedElements = append(searchedElements, elementToSearch)
				searchList = append(searchList, graph[elementToSearch]...)
			}
		}
	}
	return fmt.Sprintf("%s is not related to %s.", element, startNode)
}
