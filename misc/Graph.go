package main

import "fmt"

type Vertex struct {
	Val       int
	Neighbors []*Vertex
}

type Graph struct {
	undirectedGraph map[int]*Vertex
}

func Constructor(adjacencyList [][]int) *Graph {
	graph := Graph{
		map[int]*Vertex{},
	}
	for i := range adjacencyList {
		val := i + 1
		graph.undirectedGraph[val] = &Vertex{
			Val: val,
		}
	}

	for i, neighbors := range adjacencyList {
		vertex := graph.undirectedGraph[i+1]
		neighborsForVertex := []*Vertex{}
		for _, val := range neighbors {
			neighborsForVertex = append(neighborsForVertex, graph.undirectedGraph[val])
		}
		vertex.Neighbors = neighborsForVertex
	}
	return &graph
}

func (g *Graph) AddVertex(v int) {
	if _, ok := g.undirectedGraph[v]; !ok {
		g.undirectedGraph[v] = &Vertex{
			Val: v,
		}
	}
}

func (g *Graph) AddEdge(u, v int) {
	fmt.Printf("Adding vertices u:%d,v:%d\n", u, v)
	var vertexU, vertexV *Vertex
	vertexU = g.undirectedGraph[u]
	if vertexU == nil {
		vertexU = &Vertex{
			Val: u,
		}
		g.undirectedGraph[u] = vertexU
	}
	vertexV = g.undirectedGraph[v]
	if vertexV == nil {
		vertexV = &Vertex{
			Val: v,
		}
		g.undirectedGraph[v] = vertexV
	}

	neighborsOfU := vertexU.Neighbors
	if neighborsOfU == nil {
		vertexU.Neighbors = []*Vertex{vertexV}
	} else {
		foundVinU := false
		for _, neighborOfU := range neighborsOfU {
			if neighborOfU.Val == vertexV.Val {
				foundVinU = true
				break
			}
		}
		if !foundVinU {
			vertexU.Neighbors = append(vertexU.Neighbors, vertexV)
		}
	}

	neighborsOfV := vertexV.Neighbors
	if neighborsOfV == nil {
		vertexV.Neighbors = []*Vertex{vertexU}
	} else {
		foundUinV := false
		for _, neighborOfV := range neighborsOfV {
			if neighborOfV.Val == vertexU.Val {
				foundUinV = true
				break
			}
		}
		if !foundUinV {
			vertexV.Neighbors = append(vertexV.Neighbors, vertexU)
		}
	}

}
func (g *Graph) BFS(start int, visit func(v int)) {
	// visit all neighbors of start first, adding to FIFO queue
	// track visited neighbors with set/map
}

func (g *Graph) DFS(start int, visit func(v int)) {
	// for each neighbor of start, visit its neighbors until all neighbors in path visited
	// track visited neighbors with set/map
}

// BFS, unweighted
func (g *Graph) ShortestPath(from, to int) []int {
	return nil
}

func (g *Graph) HasVertex(v int) bool {
	_, ok := g.undirectedGraph[v]
	return ok
}

func (g *Graph) Degree(v int) int {
	if vertex, ok := g.undirectedGraph[v]; !ok {
		return 0
	} else {
		return len(vertex.Neighbors)
	}
}

func (g *Graph) Print() {
	for val, vertex := range g.undirectedGraph {
		neighbors := []int{}
		for _, neighborVertex := range (*vertex).Neighbors {
			neighbors = append(neighbors, neighborVertex.Val)
		}
		fmt.Printf("%d: %v\n", val, neighbors)
	}
	println()
}

func main() {
	fmt.Printf("Graph for adjacencyList %v:\n", [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	graph := Constructor([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	graph.Print()
	fmt.Printf("degree for vertex: %d is %d\n", 1, graph.Degree(1))
	graph.AddEdge(1, 3)
	fmt.Printf("degree for vertex: %d is %d\n", 1, graph.Degree(1))
	graph.Print()

	fmt.Printf("does graph have vertex: 5? %v\n", graph.HasVertex(5))
	fmt.Printf("does graph have vertex: 6? %v\n\n", graph.HasVertex(6))

	fmt.Printf("degree for vertex: %d is %d\n", 6, graph.Degree(6))

	graph.AddEdge(5, 6)

	fmt.Printf("does graph have vertex: 5? %v\n", graph.HasVertex(5))
	fmt.Printf("does graph have vertex: 6? %v\n\n", graph.HasVertex(6))

	graph.Print()
}
