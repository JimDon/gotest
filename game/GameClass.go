package game
import (
)
type GameClass struct {
	Board [][]string
	Sentinel []int
}
var states [][][]string

func (g *GameClass) Display(){
	for i,_ := range g.Board {
		displayline(g.Board[i])
	}
}

func (g *GameClass) NewBoardState(coor coordinate,mark string) [][]string{
	newBoard := g.Board
	newBoard[coor.x][coor.y] = mark
	append(states,newBoard)
}
	


