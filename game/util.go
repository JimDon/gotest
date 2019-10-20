package game
import (
	"math/rand"
	"time"
)



func getRandomCol(cols int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(cols)
}

func getLastSlotRow(col int,se sentinel) int {
	return se[col]

}
 
