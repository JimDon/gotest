package game
import (
	"fmt"
	"errors"
)
// board 
type board [][]int
// board slice for the board play log
type boardStates []board

// for the check the finished status
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

//display the board
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

//display the initial board
func (g *Game) Print(){
	
	for i :=0; i<g.Rows; i++ {
		fmt.Println()
		for j:=0; j<g.Cols; j++{
				fmt.Print(defaultSymbol)
			}	
		
		}
	fmt.Println()
}
	

//set the board states,save the old one into the log
func (g *Game) SetBoardStates(newBoard board) {
	g.BoardStates =append(g.BoardStates,newBoard)
	g.Board = newBoard
}

//get the new board
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

//check the board finish status
func (g *Game)	CheckforFinished() bool{
	senLength := len(g.Sentinel)
	for i :=0; i<senLength; i++ {
		if g.Sentinel[i] != UnavailableSlot {
			return false
		}
	}
	return true

}

//check the four connected 
func (g *Game) CheckforFourConnected(coord coordinate) bool{
		if coord.x == UnavailableSlot {
			return false
		}
		r := coord.x
		c := coord.y
	    t := g.Board[r][c]
		countHorizon := 0
		countVertical := 0
		countLeftD := 0
		countRightD := 0
		for i:=1;i<4;i++ {
			
			if(c-i>=0 && g.Board[r][c-i]==t||(c+i<g.Cols && g.Board[r][c+i]==t)){
				countHorizon += 1;
			}
			if(r-i>=0 && g.Board[r-i][c]==t)||(r+i<g.Rows && g.Board[r+i][c]==t){
				countVertical += 1;
			}
			if(r-i>=0 && c-i>=0 && g.Board[r-i][c-i]==t)||(r+i<g.Rows && c+i<g.Cols && g.Board[r+i][c+i]==t){
				countLeftD += 1;
			}
			if(r-i>=0 && c+i<g.Cols && g.Board[r-i][c+i]==t)||(r+i<g.Rows && c-i>=0 && g.Board[r+i][c-i]==t){
				countRightD += 1;
			}
		}
		if(countHorizon==3 || countVertical==3 || countLeftD==3 || countRightD==3){
			return true;
		}
		return false;

}


// player next step
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

//play the game
func (g *Game) Play(p Player)(bool,Player,coordinate){
	coord,err := g.NextStep()
	if err != nil {
		return false,p,coord
	}
	newBoard :=g.GetNewBoard(coord,p)
	g.SetBoardStates(newBoard)
	g.SetSentinel(coord.y,coord.x-1)
	if g.CheckforFinished() {
		return true,p,coord
	}
	return false,p,coord
	
}


// play and fill in the board
func (g *Game) PlayforFinished(p1 Player,p2 Player) (bool,Player){
	finished := false
	for !finished {
		finished, p,_ := g.Play(p1)
		if finished {
			return finished, p
		}
		finished, p,_ = g.Play(p2)
		if finished {
			return finished, p
		}
		

	} 
	return false, p1
}

//play for the four connected 
func (g *Game) PlayforFourConnected(p1 Player,p2 Player) (bool,Player,error){
	finished := false
	for !finished {

		finished, p,coord := g.Play(p1)
		if g.CheckforFourConnected(coord) {
			return true,p1,nil
		}

		if finished {
			return finished, p,errors.New("finished and equal")
		}
		finished, p,coord = g.Play(p2)
		if g.CheckforFourConnected(coord) {
			return true,p2,nil
		}
		if finished {
			return finished, p,errors.New("finished and equal")
		}
		

	} 
	return false, p1,nil

}	


	


