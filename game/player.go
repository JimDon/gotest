package game
type player struct{
	name string
	chessIcon string
	winner bool
}
func (p *player)NextStep() coordinate{
	return coordinate{2,3}
}