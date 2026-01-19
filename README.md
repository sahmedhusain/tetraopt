# Tetris Optimizer 🧩

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)

---

<p align="center">
  🧩 <strong>Tetris Optimizer</strong><br/>
  <em>Optimal tetromino arrangement powered by Go</em>
</p>

<p align="center">
  Backtracking algorithm • File parsing • Smallest square optimization
</p>

---

<p align="center">
  <strong>Arrange tetrominoes into the smallest possible square.</strong><br/>
  <em>Smart. Optimized. Efficient.</em>
</p>

<!-- 🔗 Quick Navigation -->
<p align="center">
  <a href="#-features">Features</a> •
  <a href="#-technologies-used">Tech Stack</a> •
  <a href="#-getting-started">Getting Started</a> •
  <a href="#-how-to-use">Usage</a> •
  <a href="#-application-architecture">Architecture</a>
</p>

---

## Overview

**Tetris Optimizer** is a command-line program written in **Go** that reads tetromino pieces from a text file and arranges them into the smallest possible square using a backtracking algorithm.

Designed as an algorithmic optimization project, Tetris Optimizer demonstrates core concepts such as file parsing, shape rotation, recursive backtracking, and spatial optimization.

---

## ✨ Features

Tetris Optimizer includes the following core features:

- **File Parsing** 📄  
  Reads and validates tetromino pieces from text files.

- **Shape Validation** ✅  
  Ensures all pieces are valid (4 connected blocks, proper format).

- **Rotation Logic** 🔄  
  Automatically generates all unique rotations of each piece.

- **Backtracking Algorithm** 🎯  
  Intelligently places pieces using recursive backtracking.

- **Size Optimization** 📐  
  Finds the smallest possible square that fits all pieces.

- **Error Handling** ⚠️  
  Prints "ERROR" for invalid input files or malformed tetrominoes.

---

## 🛠️ Technologies Used

- **Go** 🐹 – Core language and standard libraries
- **File I/O** 📁 – Reading and parsing tetromino files
- **Algorithms** 🧮 – Backtracking and recursive optimization
- **Data Structures** 📊 – 2D grids and coordinate systems

<!-- 🧩 Technology Logo -->
<p align="center">
  <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" width="48"/>
</p>

---

## 🎯 Project Objective

Tetris Optimizer provides an efficient solution for arranging tetromino pieces into the smallest possible square.

### Core Responsibilities

1. **Parser** – Read and validate tetromino files
2. **Tetromino** – Handle shape representation and rotation
3. **Board** – Manage piece placement and grid operations
4. **Solver** – Find optimal arrangement using backtracking

The application uses a backtracking algorithm combined with rotation optimization to efficiently find the smallest square.

---

## 🚀 Getting Started

### Prerequisites

- Go version **1.20 or newer**
- A terminal to run the program

### Installation & Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/sayehusain/tetris-optimizer.git
   ```

2. Navigate to the project directory:

   ```bash
   cd tetris-optimizer
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Build the program:
   ```bash
   go build -o tetris-optimizer
   ```

---

## 📖 How to Use

Run the program with a text file containing tetrominoes as the argument.

### Basic Usage

```bash
go run . testdata/sample.txt
```

Or using the compiled binary:

```bash
./tetris-optimizer testdata/sample.txt
```

### Input File Format

Each tetromino must be represented as a 4x4 grid using `#` for blocks and `.` for empty spaces, separated by empty lines:

```text
...#
...#
...#
...#

....
....
....
####
```

---

## 💻 Terminal Examples

### Running the Program

```bash
$ go run . testdata/sample.txt
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

### Invalid Input

```bash
$ go run . testdata/invalid/bad_format.txt
ERROR
```

### Building and Running

```bash
$ go build -o tetris-optimizer
$ ./tetris-optimizer testdata/sample.txt
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

---

## 🛠️ Application Architecture

```text
┌─────────────┐
│  Input File │
│  (*.txt)    │
└──────┬──────┘
       │
┌──────▼──────┐
│   Parser    │ → Validate & Parse
└──────┬──────┘
       │
┌──────▼──────┐
│  Tetromino  │ → Generate Rotations
└──────┬──────┘
       │
┌──────▼──────┐
│   Solver    │ → Backtracking Algorithm
│  (Algorithm)│
└──────┬──────┘
       │
┌──────▼──────┐
│    Board    │ → Place Pieces
└──────┬──────┘
       │
       ▼
   Solution Output
```

### Design Overview

- Parser validates and reads tetromino files
- Each piece is normalized and rotations are generated
- Backtracking algorithm tries all placements
- Board size increases until a solution is found

### Algorithm Details

- **Input Validation** – Ensures valid tetromino format
- **Rotation Generation** – Creates all unique orientations
- **Backtracking** – Recursively tries all placements
- **Size Optimization** – Starts with minimum square size
- **Error Handling** – Returns "ERROR" for invalid input

### Project Structure

```
tetris-optimizer/
├── main.go              # Entry point
├── packages/            # Core packages
│   ├── parser.go       # File parsing
│   ├── tetromino.go    # Shape operations
│   ├── board.go        # Grid management
│   └── solver.go       # Backtracking algorithm
├── tests/              # Unit tests
└── testdata/           # Test files
```

---

## 📋 Algorithm Details

- Validates tetromino connectivity
- Generates unique rotations only
- Uses backtracking for optimal placement
- Incrementally tries larger board sizes
- Identifies pieces with letters (A, B, C...)

---

## 🤝 Contributing

Contributions are welcome. Fork the repository, implement improvements, and submit a pull request.

---

## 📄 License

This project is licensed under the **MIT License**. See [LICENSE.md](LICENSE.md) for details.

---

## 🙏 Acknowledgments

Built as part of a Go learning journey with a focus on algorithms, file parsing, and optimization techniques.

---

## 👥 Authors

- **Sayed Ahmed Husain** – [sayedahmed97.sad@gmail.com](mailto:sayedahmed97.sad@gmail.com)

---

## 📚 What I Learned

- File I/O and parsing in Go
- Implementing backtracking algorithms
- Shape rotation and normalization
- Recursive problem-solving
- Algorithm optimization techniques

---

## ⚠️ Limitations

- Only standard tetrominoes (4 connected blocks)
- Maximum 26 pieces (limited by alphabet)
- Text-based input/output only
- No GUI interface

---

## 🔮 Future Improvements

- Add visual output representation
- Implement piece color coding
- Support custom piece shapes
- Add performance benchmarking
- Create interactive solver visualization
