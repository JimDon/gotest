package game
type player struct{
	id int
	Symbol string
}
func (p *player)NextStep() coordinate{
	return coordinate{2,3}
}