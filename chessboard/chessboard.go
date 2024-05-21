package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (total int) {
	for _, v := range cb[file] {
		if v {
			total++
		}
	}
	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (total int) {
	if rank < 1 || rank > 8 {
		return total
	}
	for _, v := range cb {
		if v[rank-1] {
			total++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (total int) {
	for _, v := range cb {
		total += len(v)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (total int) {
	for k, _ := range cb {
		total += CountInFile(cb, k)
	}
	return
}
