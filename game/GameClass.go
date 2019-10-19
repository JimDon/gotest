package game
import (
)
type GameClass struct {
	Board [][]int
	Sentinel []int
}
var states []GameClass

func (g *GameClass) Display(){
	for i,_ := range g.Board {
		displayline(g.Board[i])
	}
}



func (g *GameClass)NewState(){
	


}