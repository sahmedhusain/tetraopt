package packages

import "strings"

type Board struct {
	Size int
	Grid [][]rune
}

func NewBoard(size int) *Board {
	Slice := make([][]rune, size)
	for i := range Slice {
		Slice[i] = make([]rune, size)
		for j := range Slice[i] {
			Slice[i][j] = ' '
		}
	}
	return &Board{
		Size: size,
		Grid: Slice,
	}
}

func CanPlace(t Tetromino, row, col int, b *Board) bool {

	for _, block := range t.Blocks {
		newRow := block.Row + row
		newCol := block.Col + col
		if newRow < 0 || newRow >= b.Size || newCol < 0 || newCol >= b.Size {
			return false
		}
		if b.Grid[newRow][newCol] != ' ' {
			return false
		}
	}
	return true
}

func Place(t Tetromino, row, col int, b *Board) *Board {
	for _, block := range t.Blocks {
		newRow := block.Row + row
		newCol := block.Col + col
		b.Grid[newRow][newCol] = t.Letter
	}

	return b
}

func Remove(t Tetromino, row, col int, b *Board) *Board {
	for _, block := range t.Blocks {
		newRow := block.Row + row
		newCol := block.Col + col
		b.Grid[newRow][newCol] = ' '
	}
	return b
}

func (b *Board) String() string {
	var sb strings.Builder
	for _, row := range b.Grid {
		for _, cell := range row {
			sb.WriteRune(cell)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
