package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var result int = 0
	values, exists := cb[file]
	if !exists {
		return result
	}

	for _, value := range values {
		if value {
			result += 1
		}
	}

	return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var result int = 0
	if rank < 1 && rank > 8 {
		return result
	}

	for _, file := range cb {
		fmt.Println("file", file)
		for i, v := range file {
			if i+1 == rank && v {
				result += 1
			}
		}
	}

	return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	result := 0
	for _, c := range cb {
		for _, file := range c {
			if file {
				result += 1
			}
		}
	}

	return result
}
