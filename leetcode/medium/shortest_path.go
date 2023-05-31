/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/26
 * @contact frankstarye@tencent.com
 **/

package medium

func shortestPathBinaryMatrix(grid [][]int) int {

	n := len(grid)

	moves := [][]int{
		{-1, -1}, //西北
		{-1, 0},  //正北
		{-1, 1},  //东北
		{0, -1},  //正西
		{0, 1},   //正东
		{1, -1},  //西南
		{1, 0},   //正南
		{1, 1},   //东南
	}

	if grid[0][0] != 0 {
		return -1
	}

	queue := make([][]int, 0)
	queue = enqueue(queue, []int{0, 0})
	dist := 1
	var ele []int
	for len(queue) > 0 {
		dist++
		size := len(queue)
		for i := 0; i < size; i++ {
			queue, ele = dequeue(queue) //
			grid[ele[0]][ele[1]] = 2
			for _, move := range moves {
				r := move[0] + ele[0]
				c := move[1] + ele[1]
				if r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0 {
					if r == n-1 && c == n-1 {
						return dist
					}
					queue = enqueue(queue, []int{r, c})
				}
			}
		}

	}
	return -1

}

func contains(queue [][]int, ele []int) bool {
	for _, e := range queue {
		if e[0] == ele[0] && e[1] == ele[1] {
			return true
		}
	}
	return false
}

func enqueue(queue [][]int, ele []int) [][]int {
	queue = append(queue, ele)
	return queue
}

func dequeue(queue [][]int) ([][]int, []int) {
	ele := queue[0]
	if len(queue) == 1 {
		return [][]int{}, ele
	}
	return queue[1:], ele
}
