package search

func WordSearch(board [][]byte, word string) bool {
	directions := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, -1},
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// fmt.Printf("start row: %d, col: %d\n", i, j)
			if dfs(board, word, i, j, directions, visited) {
				return true
			}
			// fmt.Print("not found\n", i, j)
		}
	}
	return false

}
func dfs(board [][]byte, word string, row, col int, directions [][]int, visited [][]bool) (exist bool) {
	if len(word) == 0 {
		return true
	}
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	if visited[row][col] {
		return false
	}
	if board[row][col] == word[0] {
		visited[row][col] = true
		exist := false
		for _, pair := range directions {
			if dfs(board, word[1:], row+pair[0], col+pair[1], directions, visited) {
				exist = true
				break
			}
		}
		visited[row][col] = false
		return exist
	}
	return false
}
