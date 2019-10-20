package game

type board [][]int
type boardStates [][][]int
type sentinel []int
type coordinate struct {
	x int
	y int
}

type Game struct {
	Board board
	BoardStates boardStates
	Sentinel sentinel
	
}


func (g *Game) Display(){
	for i,_ := range g.Board {
		displayline(g.Board[i])
	}
}

func (g *Game) SetBoardStates(newBoard board) boardStates{
	newstates :=append(g.BoardStates,newBoard)
	return newstates
}
func (g *Game) GetNewBoard(coord coordinate,player player) board{
	newBoard := g.Board
	newBoard[coord.x][coord.y]= player.id
	return newBoard
}
func (g *Game)	SetSentinel(col int,slot int,se sentinel){
	g.Sentinel[col] = slot

}

func (g *Game)	CheckforFirstFinished() bool{
	senLength := len(g.Sentinel)
	for i :=0; i<senLength; i++ {
		if g.Sentinel[i] != -1 {
			return false
		}
	}
	return true

}

func (g *Game) CheckforFourConnected() bool{
	newBoard :=
}




func (g *Game) PlayforFirstFinished(p1 player,p2 player){
	
}
func (g *Game) PlayforFourConnected(p1 player,p2 player){


}
	


