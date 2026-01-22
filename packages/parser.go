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

func ParseFile(filepath string) ([]Tetromino, error) {
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
	id := 0

	for i := 0; i < len(lines); {
		if i+4 > len(lines) {
			return nil, fmt.Errorf("ERROR: incomplete tetromino at line %d", i+1)
		}

		blockLines := lines[i : i+4]
		tetromino, err := parseTetromino(blockLines, id)
		if err != nil {
			return nil, err
		}
		tetrominoes = append(tetrominoes, *tetromino)
		id++
		i += 4

		if i < len(lines) {
			if lines[i] != "" {
				return nil, fmt.Errorf("ERROR: expected empty line separator at line %d", i+1)
			}
			i++
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
	if !IsConnected(blocks) {
		return nil, fmt.Errorf("ERROR: Tetromino blocks are not connected")
	}

	blocks = Normalize(blocks)

	letter := rune('A' + id)
	return &Tetromino{
		ID:     id,
		Blocks: blocks,
		Letter: letter,
	}, nil
}

func IsConnected(blocks []Point) bool {
	visited := make(map[Point]bool)

	var dfs func(Point)
	dfs = func(p Point) {
		if visited[p] {
			return
		}
		visited[p] = true

		neighbors := []Point{
			{p.Row - 1, p.Col},
			{p.Row + 1, p.Col},
			{p.Row, p.Col - 1},
			{p.Row, p.Col + 1},
		}

		for _, neighbor := range neighbors {
			for _, block := range blocks {
				if block == neighbor {
					dfs(neighbor)
				}
			}
		}
	}

	dfs(blocks[0])
	return len(visited) == 4
}

func Normalize(blocks []Point) []Point {

	minRow := blocks[0].Row
	minCol := blocks[0].Col

	for _, block := range blocks {
		if block.Row < minRow {
			minRow = block.Row
		}
		if block.Col < minCol {
			minCol = block.Col
		}
	}

	for i, block := range blocks {
		blocks[i] = Point{Row: block.Row - minRow, Col: block.Col - minCol}
	}

	return blocks
}
