package game
import (
	"fmt"
	"errors"
)
type board [][]int
type boardStates []board
type sentinel []int
const UnavailableSlot = -1
const defaultSymbol = "."
type Player struct{
	Id int
	Symbol string
}

type coordinate struct {
	x int
	y int
}

type Game struct {
	Rows int
	Cols int
	Board board
	BoardStates boardStates
	Sentinel sentinel
	
}


func (g *Game) Display(p1 Player,p2 Player){
	for i :=0; i<g.Rows; i++ {
		fmt.Println()
		for j:=0; j<g.Cols; j++{
			switch g.Board[i][j] {
			case p1.Id:
				fmt.Print(p1.Symbol)
			case p2.Id:
				fmt.Print(p2.Symbol)
			default:
				fmt.Print(defaultSymbol)
			}	
		
		}
			 	 
	}
	fmt.Println()
}
func (g *Game) Print(){
	
	for i :=0; i<g.Rows; i++ {
		fmt.Println()
		for j:=0; j<g.Cols; j++{
				fmt.Print(defaultSymbol)
			}	
		
		}
	fmt.Println()
}
	


func (g *Game) SetBoardStates(newBoard board) {
	g.BoardStates =append(g.BoardStates,newBoard)
	g.Board = newBoard
}
func (g *Game) GetNewBoard(coord coordinate,p Player) board{
	if coord.x == UnavailableSlot {
		return g.Board
	}
	newBoard := g.Board

	newBoard[coord.x][coord.y]= p.Id
	return newBoard
}
func (g *Game)	SetSentinel(col int,slot int){
	g.Sentinel[col] = slot

}

func (g *Game)	CheckforFinished() bool{
	senLength := len(g.Sentinel)
	for i :=0; i<senLength; i++ {
		if g.Sentinel[i] != UnavailableSlot {
			return false
		}
	}
	return true

}

func (g *Game) CheckforFourConnected() bool{
	// newBoard := g.BoardStates[len(g.BoardStates)-1]
	// rows := len(newBoard)-1
	// cols := len(newBoard[0])-1
	
	return false

}

func (g *Game)NextStep()(coordinate,error) {
	y := getRandomCol(g.Cols)
	
	x := getLastSlotRow(y,g.Sentinel)
	var err error
	if x == UnavailableSlot {
		err = errors.New("no slot to apply")
	}
	return coordinate{
		x: x,
		y: y,
	}, err
}

func (g *Game) Play(p Player)(bool,Player){
	coord,err := g.NextStep()
	if err != nil {
		return false,p
	}
	newBoard :=g.GetNewBoard(coord,p)
	g.SetBoardStates(newBoard)
	g.SetSentinel(coord.y,coord.x-1)
	if g.CheckforFinished() {
		return true,p
	}
	return false,p
	
}
func (g *Game) PlayforFinished(p1 Player,p2 Player) (bool,Player){
	finished := g.CheckforFinished()
	for !finished {
		finished, p := g.Play(p1)
		if finished {
			return finished, p
		}
		finished, p = g.Play(p2)
		if finished {
			return finished, p
		}
		

	} 
	return false, p1
}
func (g *Game) PlayforFourConnected(p1 Player,p2 Player){


}
	


