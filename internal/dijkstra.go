package internal

import (
	"math"
	"slices"
	"strings"
)

type dijkstraResult struct {
	path string
	cost uint
}

/*
getCosts initializes and returns a map of the minimum costs from the start node
to each other node in the graph. Nodes directly connected to the start node
have their initial costs set based on the graph's weights, while other nodes
start with a cost of infinity.

Parameters:

	graph map[string]map[string]uint: A map representing the graph, where each
		node points to a map of adjacent nodes with weights.
	startNode string: The starting node for calculating initial costs.

Returns:

	map[string]uint: A map of nodes to their initial minimum costs from the
	start node.
*/
func getCosts(graph map[string]map[string]uint, startNode string) map[string]uint {
	costs := make(map[string]uint)

	for key, value := range graph[startNode] {
		costs[key] = value
	}

	for key := range graph {
		if key == startNode {
			costs[key] = 0
			// Updating only the cost of the nodes away from the start node
		} else if _, alreadyInMap := costs[key]; !alreadyInMap {
			costs[key] = uint(math.Inf(1))
		}
	}

	return costs
}

/*
getParents initializes and returns a map of parent nodes for each node in the
graph. Direct children of the start node have their parent set to the start
node, while other nodes are set with empty strings.

Parameters:

	graph map[string]map[string]uint: A map representing the graph, where each
		node points to a map of adjacent nodes with weights.
	startNode string: The starting node used to initialize parent relationships.

Returns:

	map[string]string: A map of nodes to their respective parent nodes.
*/
func getParents(graph map[string]map[string]uint, startNode string) map[string]string {
	startNodeSons := graph[startNode]
	parents := make(map[string]string)

	for key := range graph {
		if _, present := startNodeSons[key]; present {
			parents[key] = startNode
		} else {
			parents[key] = ""
		}
	}

	return parents
}

/*
findCheaperNode finds the unprocessed node with the lowest cost in the costs
map.

Parameters:

	costs map[string]uint: A map of nodes to their respective costs.
	processedNodes []string: A slice of nodes that have already been processed.

Returns:

	string: The node with the lowest cost that hasn't been processed yet. If
		no such node is found, returns an empty string.
*/
func findCheaperNode(costs map[string]uint, processedNodes []string) (cheaperNode string) {
	cheaperCost := uint(math.Inf(1))

	for node, cost := range costs {
		if cost < cheaperCost && !slices.Contains(processedNodes, node) {
			cheaperNode = node
			cheaperCost = cost
		}
	}

	return
}

/*
parseDijkstraResult constructs the path from the start node to the end node by
tracing back from the end node through the parents map, calculating the total
cost of this path.

Parameters:

	parents map[string]string: A map of nodes to their respective parent nodes.
	costs map[string]uint: A map of nodes to their minimum costs from the start
		node.
	endNode string: The target node for which to retrieve the path and cost.

Returns:

	dijkstraResult: A struct containing the path from start node to end node as
		a string and the total cost.
*/
func parseDijkstraResult(parents map[string]string, costs map[string]uint, endNode string) dijkstraResult {
	var path []string
	totalCost := costs[endNode]

	currentNode := endNode
	for currentNode != "" {
		path = append(path, currentNode)
		currentNode = parents[currentNode]
	}

	slices.Reverse(path)

	return dijkstraResult{strings.Join(path, " -> "), totalCost}
}

/*
Dijkstra executes Dijkstra's algorithm on a given weighted graph to find the
shortest path from the start node to the end node. It calculates the minimum
costs and parent relationships to reconstruct the path.

Parameters:

	graph map[string]map[string]uint: A map representing the graph, where each
		node points to a map of adjacent nodes with weights.
	startNode string: The node from which to start the search.
	endNode string: The node for which to find the shortest path from the start
		node.

Returns:

	dijkstraResult: A struct containing the path from the start node to the end
		node as a string and the total cost of that path.
*/
func Dijkstra(graph map[string]map[string]uint, startNode, endNode string) dijkstraResult {
	costs := getCosts(graph, startNode)
	parents := getParents(graph, startNode)

	processedNodes := []string{startNode}
	node := findCheaperNode(costs, processedNodes)

	for node != "" {
		cost := costs[node]
		neighbors := graph[node]

		for neighborNode, neighborCost := range neighbors {
			newCost := cost + neighborCost
			if costs[neighborNode] > newCost {
				costs[neighborNode] = newCost
				parents[neighborNode] = node
			}
		}

		processedNodes = append(processedNodes, node)
		node = findCheaperNode(costs, processedNodes)
	}

	return parseDijkstraResult(parents, costs, endNode)
}
