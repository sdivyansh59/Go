package main

func solve(A int , B [][]int )  (int) {
    
    visited := make([]bool,A)
    adj := make([][]int, A)
    for i,_ := range adj {
        adj[i] = make([]int,0)
    }
    // fill adj 
    for _,v := range B {
        v1 := v[0]-1
        v2 := v[1]-1
        adj[v1] = append(adj[v1],v2)
    }
    
    parent  := -1
    // visited[source]= true
    for i,v := range visited {
        if !v {
            if dfs(adj ,parent ,i , visited) == 1 {
                return 1
            }
        }
    }
    return 0
     
}

func  dfs(adj[][]int,parent, source int, visited []bool) int{
     visited[source]= true
     
     for _, v := range adj[source] {
         if visited[v] && v != parent {
             return 1
         }
         if !visited[v] {
            if  dfs(adj,source,v,visited) == 1 {
                return 1
            }
         }
     }
     
     return 0  
}

