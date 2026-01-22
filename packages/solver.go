package packages

import "math"

func CalculateMinSize(numPieces int) int {
	totalBlocks := numPieces * 4
	size := int(math.Sqrt(float64(totalBlocks)))
	if size*size < totalBlocks {
		size++
	}
	return size
}

func Solve(tetrominoes []Tetromino) *Board {
	if len(tetrominoes) == 0 {
		return nil
	}
	minSize := CalculateMinSize(len(tetrominoes))
	maxSize := minSize * 2
	for size := minSize; size <= maxSize; size++ {
		board := NewBoard(size)
		if solveRecursive(board, tetrominoes, 0) {
			return board
		}
	}
	return nil
}

func solveRecursive(board *Board, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			for _, rotation := range GetAllRotations(tetrominoes[index]) {
				if CanPlace(rotation, i, j, board) {
					Place(rotation, i, j, board)

					if solveRecursive(board, tetrominoes, index+1) {
						return true
					}
					Remove(rotation, i, j, board)
				}
			}
		}
	}
	return false
}
