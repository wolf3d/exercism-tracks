package chessboard

//import "fmt"

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool
// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank
// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    var piecesInRank int = 0
    inRank, ok := cb[rank];
    if !ok {
        return piecesInRank
    }
	for i := 0; i < len(inRank); i++ {
        if inRank[i] {
            piecesInRank += 1
        }
    }
	return piecesInRank
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    var piecesInFile int = 0
	if file < 1 || file > 8 {
        return piecesInFile
    }
	for _,v := range cb {
        if v[file-1] {
            piecesInFile += 1
        }
    }
	return piecesInFile
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    var squareCount int = 0    
    for _, v := range cb {
        for i := 0; i < len(v);i++ {
			squareCount += 1
        }        
    }
	return squareCount
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    var squareCount int = 0    
    for _, v := range cb {
        for i := 0; i < len(v);i++ {
            if v[i] {
            	squareCount += 1    
            }			
        }        
    }
	return squareCount
}
