package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (accum int) {
	actualRank, exists := cb[rank]
	// Note we don't need to code in the requirement of "Return a count of zero
	// if the given rank cannot be found in the map", since Go returns a zero-value
	// if the key is missing from a map - but let's do it explicitly anyway to make
	// intent more legible
	if !exists {
		return
	}

	// I bet Go has a built-in "fold" function to do this more functionally
	for _, val := range actualRank {
		if val {
			accum++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (accum int) {
	if file < 1 || file > 8 {
		return
	}
	for _, rank := range cb {
		if rank[file-1] { // files are 1-indexed
			accum++
		}
	}
	return accum
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (accum int) {
	for idx, _ := range cb {
		accum += CountInRank(cb, idx)
	}
	return
}
