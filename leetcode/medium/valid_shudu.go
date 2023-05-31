/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/20
 * @contact frankstarye@tencent.com
 * @desc id:36 有效的数独
 **/

package medium

//Item 可能存在的元素
type Item struct {
	Row int
	Col int
	Ele []int
}


//isValidSudoku 有效的数独
func isValidSudoku(board [][]byte) bool {
	size := len(board)
	row := [9][9]int{}
	col := [9][9]int{}
	circle := [3][3][9]int{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == '.' {
				continue
			} else {
				num := board[i][j] - '1'
				row[i][num] ++
				col[j][num] ++
				circle[i/3][j/3][num] ++

				if row[i][num] > 1 || col[j][num] > 1 ||  circle[i/3][j/3][num] > 1 {
					return false
				}
			}

		}
	}
	return true

}
