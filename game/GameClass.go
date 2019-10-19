package game
import (
)
type GameClass struct {
	Board [][]int
	Sentinel []int
}
func (g *GameClass) Display(){
	for i,_ := range g.Board {
		displayline(g.Board[i])
	}
}
