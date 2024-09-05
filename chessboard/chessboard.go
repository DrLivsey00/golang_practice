package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var counter int
	files, err := cb[file]
	if !err {
		return 0
	}
	for _, x := range files {
		if x {
			counter++
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	var counter int
	for _, x := range cb {
		if x[rank-1] {
			counter++
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var counter int
	for _, x := range cb {
		for _, _ = range x {
			counter++
		}
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var counter int
	for _, x := range cb {
		for _, y := range x {
			if y {
				counter++
			}
		}
	}
	return counter
}
