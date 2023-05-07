// package main

// import "fmt"


// func main() {
// 	fmt.Print("Enter Vertex: ")
// 	var v int
// 	fmt.Scan(&v)
// 	var e int 

// 	fmt.Print("enter edge: ")
// 	fmt.Scan(&e)

// 	adj := make([][] int,v)
// 	// initialize 
// 	for i := range adj {
// 		adj[i] = make([]int,0)
// 	}

// 	addEdge(adj,0,1)
// 	addEdge(adj,0,2)
// 	addEdge(adj,1,2)
// 	addEdge(adj,1,3)

// 	bfs(adj)

// 	// fmt.Println(adj)

// }



// func print(adj [][]int) {
// 	// print 
// 	for _,v := range adj {
// 		for j := range v {
// 			fmt.Print(j, " ")
// 		}
// 		fmt.Println()
// 	}
// }



// func addEdge(adj [][] int, vertex, edge int) {
// 	adj[vertex] = append(adj[vertex], edge)
// 	adj[edge] = append(adj[edge],vertex )
// }

// func bfs(adj [][]int) {
// 	v := len(adj)
// 	visited := make ([] bool, v)

// 	var start int

// 	queue := [] int {start}

// 	for len(queue) >0 {
// 		vertex := queue[len(queue)-1]
// 		fmt.Print(vertex," ")
// 		visited[vertex] = true
// 		queue = queue[:len(queue)-1]

// 		for  val := range adj[vertex] {
// 			if !visited[val] {
// 				queue = append(queue, val)
// 			}
			
// 		}
// 	}


// }