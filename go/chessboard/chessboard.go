package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	if f, ok := cb[file]; ok {
		c := 0
		for _, o := range f {
			if o {
				c += 1
			}
		}
		return c
	}
	return 0
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	if rank < 1 || rank > 8 {
		return 0
	}

	count := 0
	for _, file := range cb {
		if len(file) >= rank {
			if file[rank-1] {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// iterate over the cb
	// iterate over each rank
	// add to the count

	count := 0
	for _, file := range cb {
		for range file {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, rank := range file {
			if rank {
				count++
			}
		}
	}
	return count
}
