# Tetris Optimizer - Test Data Guide

This directory contains test files for the audit process.

## Directory Structure

```
testdata/
├── valid/              # Valid test cases that should produce solutions
│   ├── good_example_00.txt  # Should have 0 empty spaces
│   ├── good_example_01.txt  # Should have 9 empty spaces
│   ├── good_example_02.txt  # Should have 4 empty spaces
│   ├── good_example_03.txt  # Should have 5 empty spaces
│   └── hard_example.txt     # Should have 1 empty space
│
├── invalid/            # Invalid test cases that should output "ERROR"
│   ├── bad_example_00.txt
│   ├── bad_example_01.txt
│   ├── bad_example_02.txt
│   ├── bad_example_03.txt
│   ├── bad_example_04.txt
│   └── bad_format.txt
│
└── sample.txt          # Sample file from requirements

```

## Running Tests

To run all audit tests:

```bash
./test.sh
```

## Test File Format

Each tetromino should be a 4x4 grid:

- Use `#` for blocks
- Use `.` for empty spaces
- Separate tetrominoes with an empty line

### Valid Example:

```
...#
...#
...#
...#

....
....
....
####
```

### Invalid Examples:

- Not exactly 4 blocks per tetromino
- Blocks not connected
- Wrong grid dimensions (not 4x4)
- Invalid characters
- Missing blank line separators

## Expected Results

### Valid Files

- Should produce a square output
- Each tetromino labeled with uppercase letters (A, B, C, ...)
- Empty spaces shown as `.`
- Smallest possible square

### Invalid Files

- Should output: `ERROR`
