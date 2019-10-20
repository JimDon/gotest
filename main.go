package main
import (
	game "github.com/JimDon/gotest/game"
)

func main() {
	newBoard := makeBoard(6,7)
	sentinel := makeSentinel(6,7)
	newGame := &game.Game{
		Rows: 6,
		Cols: 7,
		Board: newBoard,
		Sentinel: sentinel,
	}
	newGame.InitDisplay()
	p1 := game.Player{
		Id:1,
		Symbol: "X",
	}
	p2 := game.Player{
		Id:2,
		Symbol: "O",
	}
	
	finished,_ := newGame.PlayforFinished(p1,p2)
	if finished {
		newGame.Display(p1,p2)
	}

}

func makeBoard(row int,col int) [][]int{
	newBoard := make([][]int, row)
    for i := 0; i < row; i++ {
        newBoard[i] = make([]int, col)
        for j := 0; j < col; j++ {
            newBoard[i][j] = 0
        }
	}
	return newBoard
}

func makeSentinel(row int, col int) []int{
	sentinel := make ([]int,col)
	for k :=0; k<col; k++{
		sentinel[k] = row -1
	}
	return sentinel
}

