/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/29
 * @contact frankstarye@tencent.com
 **/

package graph

import (
	"fmt"
)

type Graph struct {
	Vertex []Vertex
	Vm     map[string][]Vertex
}

type Vertex struct {
	Name     string
	Row      int
	Col      int
	Neighbor []Vertex //相邻的节点
}

func NewGraph(grid [][]int) Graph {
	n := len(grid)
	vs := make([]Vertex, 0)
	gm := make(map[string][]Vertex)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == 0 {
				v := Vertex{
					Name: fmt.Sprintf("r%d-c%d", row, col),
					Row:  row,
					Col:  col,
				}

				neighbors := make([]Vertex, 0)
				// 八个方向 获取neighbors
				//西北
				if row-1 >= 0 && col-1 >= 0 && grid[row-1][col-1] == 0 {
					westNorth := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row-1, col-1),
						Row:  row - 1,
						Col:  col - 1,
					}
					neighbors = append(neighbors, westNorth)
				}
				//上
				if row-1 >= 0 && grid[row-1][col] == 0 {
					north := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row-1, col),
						Row:  row - 1,
						Col:  col,
					}
					neighbors = append(neighbors, north)
				}
				//东北
				if row-1 > 0 && col+1 < n && grid[row][col+1] == 0 {
					eastNorth := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row, col+1),
						Row:  row,
						Col:  col + 1,
					}
					neighbors = append(neighbors, eastNorth)
				}
				//西
				if col-1 >= 0 && grid[row][col-1] == 0 {
					west := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row, col-1),
						Row:  row,
						Col:  col - 1,
					}
					neighbors = append(neighbors, west)
				}
				//东
				if col+1 < n && grid[row][col+1] == 0 {
					east := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row, col+1),
						Row:  row,
						Col:  col + 1,
					}
					neighbors = append(neighbors, east)
				}
				//西南
				if row+1 < n && col-1 >= 0 && grid[row+1][col-1] == 0 {
					westSouth := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row+1, col-1),
						Row:  row + 1,
						Col:  col - 1,
					}
					neighbors = append(neighbors, westSouth)
				}
				//南
				if row+1 < n && grid[row+1][col] == 0 {
					south := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row+1, col),
						Row:  row + 1,
						Col:  col,
					}
					neighbors = append(neighbors, south)
				}
				//东南
				if row+1 < n && col+1 < n && grid[row+1][col+1] == 0 {
					south := Vertex{
						Name: fmt.Sprintf("r%d-c%d", row+1, col+1),
						Row:  row + 1,
						Col:  col + 1,
					}
					neighbors = append(neighbors, south)
				}
				v.Neighbor = neighbors
				gm[v.Name] = neighbors
				vs = append(vs, v)
			}
		}
	}
	return Graph{Vertex: vs, Vm: gm}
}

func (g Graph) Traverse() []string {
	if len(g.Vertex) == 0 {
		return nil
	}
	result := make([]string, 0)
	root := g.Vertex[0]
	queue := make([]Vertex, 0)
	queue = append(queue, root)
	visited := make(map[string]bool)
	var ele Vertex
	for len(queue) > 0 {
		ele, queue = dequeue(queue)
		if !visited[ele.Name] {
			visited[ele.Name] = true
			result = append(result, ele.Name)
			for _, n := range g.Vm[ele.Name] {
				if visited[n.Name] {
					continue
				}
				queue = enqueue(queue, n)
			}
		}
	}
	return result
}

func (g Graph) ShortPath() int {
	if len(g.Vertex) == 0 {
		return -1
	}

	root := g.Vertex[0]
	queue := make([]Vertex, 0)
	queue = append(queue, root)
	//遍历queue
	var ele Vertex
	visited := make(map[string]bool)
	distinct := -1
	for len(queue) > 0 {
		distinct++
		size := len(queue)
		for i := 0; i < size; i++ {
			ele, queue = dequeue(queue)
			visited[ele.Name] = true
			for _, n := range g.Vm[ele.Name] {
				if visited[n.Name] {
					continue
				}
				queue = enqueue(queue, n)
			}
		}
	}

	return distinct
}

func (g Graph) last(ele Vertex) bool {
	n := len(g.Vertex)
	last := g.Vertex[n-1]
	return last.Name == ele.Name
}

func enqueue(queue []Vertex, ele Vertex) []Vertex {
	queue = append(queue, ele)
	return queue
}

func dequeue(queue []Vertex) (Vertex, []Vertex) {
	ele := queue[0] //return top
	if len(queue) == 1 {
		return ele, []Vertex{}
	}
	return ele, queue[1:]
}
