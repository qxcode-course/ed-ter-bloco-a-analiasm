package main
import (
    "fmt"
    "io"
    
)

func maximumD(bomba [][]int) int {

    n := len(bomba)
    adj := make([][]int, n)

    for i:= 0; i < n; i++{

        x1, y1, r1 := int64(bomba[i][0]), int64(bomba[i][1]), int64(bomba[i][2])
        r1Sq := r1 * r1

        for j := 0; j < n; j++ {

            if i == j {
                continue
            }

            x2, y2 := int64(bomba[j][0]), int64(bomba[j][1])

            dx := x1 - x2
            dy := y1 - y2
            distSq := dx*dx + dy*dy

            if distSq <= r1Sq{
                adj[i] = append(adj[i], j)
            }
        }
    }
          
    maxD := 0

    for i := 0; i < n; i++{
        visited := make([]bool, n)

        count := dfs(i, adj, visited)

        if count > maxD {
            maxD = count
        }
    }

    return maxD
}

func dfs(node int, adj [][]int, visited []bool) int{
    visited[node] = true
    count := 1

    for _, vizinho := range adj[node] {
        if !visited[vizinho]{
            count += dfs(vizinho, adj, visited)
        }
    }
    return count
}


func main() {

    var n, m int

    for {
        _, err := fmt.Scan(&n, &m)

        if err != nil{
            if err == io.EOF{
                break
            }
            break
        }

    exemploBombas := make([][]int, n)

    for i := 0; i < n; i++{
        exemploBombas[i] = make([]int, 3)

        fmt.Scan(&exemploBombas[i][0], &exemploBombas[i][1], &exemploBombas[i][2])
    }

    if n > 0{
    fmt.Printf("%d\n", maximumD(exemploBombas))
    
        }
    }
}