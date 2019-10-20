package main
import (
	game "github.com/JimDon/gotest/game"
	"fmt"
)

func main() {
	//Task1 execution
	Task1()

	//Task2 execution
	Task2()

	//Task3 execution
	Task3()
	
}

//initialize the board
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

//initialize the sentinel
func makeSentinel(row int, col int) []int{
	sentinel := make ([]int,col)
	for k :=0; k<col; k++{
		sentinel[k] = row -1
	}
	return sentinel
}

func Task1(){
	const rows =6
	const cols =7
	newBoard := makeBoard(rows,cols)
	sentinel := makeSentinel(rows,cols)
	newGame := &game.Game{
		Rows: rows,
		Cols: cols,
		Board: newBoard,
		Sentinel: sentinel,
	}
	p1 := game.Player{
		Id:1,
		Symbol: "X",
	}
	p2 := game.Player{
		Id:2,
		Symbol: "O",
	}
	newGame.Print()
	newGame.Board[0][3] =1
	newGame.Board[3][2] =1
	newGame.Board[2][3] =2
	newGame.Board[3][5] =2
	newGame.Display(p1,p2)

}

func Task2(){
	const rows =6
	const cols =7
	newBoard := makeBoard(rows,cols)
	sentinel := makeSentinel(rows,cols)
	newGame := &game.Game{
		Rows: rows,
		Cols: cols,
		Board: newBoard,
		Sentinel: sentinel,
	}
	newGame.Print()
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

func Task3(){
	const rows =6
	const cols =7
	newBoard := makeBoard(rows,cols)
	sentinel := makeSentinel(rows,cols)
	newGame := &game.Game{
		Rows: rows,
		Cols: cols,
		Board: newBoard,
		Sentinel: sentinel,
	}
	newGame.Print()
	p1 := game.Player{
		Id:1,
		Symbol: "X",
	}
	p2 := game.Player{
		Id:2,
		Symbol: "O",
	}
	
	finished,p,err := newGame.PlayforFourConnected(p1,p2)
	if err !=nil {
		fmt.Println("Equal,no one win!")
		return 
	}
	
	if finished {
		newGame.Display(p1,p2)
	}
	fmt.Println("winner is",p.Id)

}