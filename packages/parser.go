package packages

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	Row int
	Col int
}

type Tetromino struct {
	ID     int
	Blocks []Point
	Letter rune
}

func Parsefile(filepath string) ([]Tetromino, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var tetrominoes []Tetromino
	var currentTetromino []string
	id := 0
	for _, line := range lines {
		if line != "" {
			currentTetromino = append(currentTetromino, line)
		}
		if len(currentTetromino) == 4 {
			tetromino, err := parseTetromino(currentTetromino, id)
			if err != nil {
				log.Fatal(err)
			}
			tetrominoes = append(tetrominoes, *tetromino)
			id++
			currentTetromino = nil
		}
	}

	return tetrominoes, nil
}

func parseTetromino(lines []string, id int) (*Tetromino, error) {

	if len(lines) != 4 {
		return nil, fmt.Errorf("ERROR: Tetromino must have exactly 4 lines, got %d", len(lines))
	}

	var blocks []Point

	for row, line := range lines {
		if len(line) != 4 {
			return nil, fmt.Errorf("ERROR: Each line must have exactly 4 characters, line %d has %d", row, len(line))
		}

		for col, char := range line {
			if char == '#' {
				blocks = append(blocks, Point{Row: row, Col: col})
			}
		}
	}

	if len(blocks) != 4 {
		return nil, fmt.Errorf("ERROR: Tetromino must have exactly 4 blocks, got %d", len(blocks))
	}
	letter := rune('A' + id)
	return &Tetromino{
		ID:     id,
		Blocks: blocks,
		Letter: letter,
	}, nil
}
