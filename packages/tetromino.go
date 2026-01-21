package packages

func Rotate(t Tetromino) Tetromino {
	rotated := Tetromino{
		ID:     t.ID,
		Letter: t.Letter,
		Blocks: make([]Point, len(t.Blocks)),
	}

	for i, block := range t.Blocks {
		rotated.Blocks[i] = Point{
			Row: block.Col,
			Col: -block.Row,
		}
	}

	rotated.Blocks = Normalize(rotated.Blocks)

	return rotated
}

func GetAllRotations(t Tetromino) []Tetromino {
	var uniqueRotations []Tetromino
	current := t

	for i := 0; i < 4; i++ {
		if !containsRotation(uniqueRotations, current) {
			uniqueRotations = append(uniqueRotations, current)
		}
		current = Rotate(current)
	}

	return uniqueRotations
}

func containsRotation(rotations []Tetromino, t Tetromino) bool {
	for _, r := range rotations {
		if areBlocksEqual(r.Blocks, t.Blocks) {
			return true
		}
	}
	return false
}

func areBlocksEqual(blocks1, blocks2 []Point) bool {
	if len(blocks1) != len(blocks2) {
		return false
	}

	// Create copies to avoid modifying original slices
	b1 := make([]Point, len(blocks1))
	b2 := make([]Point, len(blocks2))
	copy(b1, blocks1)
	copy(b2, blocks2)

	// Sort both slices for comparison
	sortBlocks(b1)
	sortBlocks(b2)

	for i := range b1 {
		if b1[i].Row != b2[i].Row || b1[i].Col != b2[i].Col {
			return false
		}
	}

	return true
}

func sortBlocks(blocks []Point) {
	// Simple bubble sort (sufficient for 4 elements)
	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {
			if blocks[i].Row > blocks[j].Row ||
				(blocks[i].Row == blocks[j].Row && blocks[i].Col > blocks[j].Col) {
				blocks[i], blocks[j] = blocks[j], blocks[i]
			}
		}
	}
}
