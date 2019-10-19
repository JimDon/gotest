package main
import (
	game "github.com/JimDon/gotest/game"
)

func main() {
	row := 6
	col := 7
	// newGame :=&game.GameClass{
	// 	Board: newBoard,
	// 	Sentinel: newSentinel,
	// }
	newBoard := make([][]string, row)
    for i := 0; i < row; i++ {
        innerLen := col
        newBoard[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            newBoard[i][j] = '0'
        }
	}
	sentinel := make ([]int,col)
	for k :=0; k<col; k++{
		sentinel[k] = 0
	}
	newGame := &game.GameClass{
		Board: newBoard,
		Sentinel: sentinel,
	}
	newGame.Display()
	
}


