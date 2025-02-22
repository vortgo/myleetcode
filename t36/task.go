package t36

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var box = make(map[string][]byte)

	bi := 0
	for i := 0; i < 9; i++ {
		var column []byte
		var row []byte
		bj := 0
		for j := 0; j < 9; j++ {
			key := fmt.Sprintf("%d-%d", bi, bj)
			if _, ok := box[key]; !ok {
				box[key] = []byte{}
			}

			if board[i][j] != '.' {
				row = append(row, board[i][j])

				box[key] = append(box[key], board[i][j])
			}
			if board[j][i] != '.' {
				column = append(column, board[j][i])
			}

			if (j+1)%3 == 0 {
				bj++
			}
		}
		if !isUnicSlice(column) || !isUnicSlice(row) {
			return false
		}

		if (i+1)%3 == 0 {
			bi++
		}
	}

	for _, v := range box {
		if !isUnicSlice(v) {
			return false
		}
	}

	return true
}

func isUnicSlice(s []byte) bool {
	exists := make(map[byte]bool)
	for _, v := range s {
		if _, ok := exists[v]; ok != true {
			exists[v] = true
		} else {
			return false
		}
	}
	return true
}
