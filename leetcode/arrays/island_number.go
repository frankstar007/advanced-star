/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/16
 * @contact frankstarye@tencent.com
 * @desc 岛屿的数量  200
 **/

package arrays

import "fmt"

type Island struct {
	Top   int
	Down  int
	Left  int
	Right int
}


//numIslands 岛屿的数量
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	//step1 先找出所有岛屿所在的位置
	locations := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				coordinate := make([]int, 0)
				coordinate = append(coordinate, i)
				coordinate = append(coordinate, j)
				locations = append(locations, coordinate)
			}
		}
	}
	if len(locations) <= 0 {
		return 0
	}
	//step2 遍历每个岛屿

	//存放每个岛屿的上下左右的最值
	result := make([]Island, 0)
	exist := make(map[string]bool)
	for k := 0; k < len(locations); k++ {
		h := locations[k][0] //水平索引
		v := locations[k][1] //垂直索引
		top, down, left, right := h, h, v, v
		island := Island{top, down, left, right}
		fmt.Println(fmt.Sprintf("location: i:%d j:%d", h, v))
		//先计算上的极限 如果top = 0 则不需要进行
		for top > 0 {
			if grid[top][v] == '0' {
				break
			}
			top--
		}
		island.Top = top
		//计算下的极限 如果down = m-1 则不需要进行
		for down < m-1 {
			if grid[down][v] == '0' {
				break
			}
			down++
		}
		island.Down = down
		//计算左的极限 如果left = 0 则不需要进行
		for left > 0 {
			if grid[h][left] == '0' {
				break
			}
			left--
		}
		island.Left = left
		//计算右的极限 如果right = n - 1 则不需要进行
		for right < n-1 {
			if grid[h][right] == '0' {
				break
			}
			right++
		}
		island.Right = right
		coordinateKey := fmt.Sprintf("%d_%d_%d_%d", top, down, left, right)
		fmt.Println(fmt.Sprintf("island: top:%d down:%d left:%d right:%d", top, down, left, right))
		if _, ok := exist[coordinateKey]; ok {
			//如果是小岛的一部分
			continue
		}
		exist[coordinateKey] = true
		result = append(result, island)
	}
	return len(result)

}
