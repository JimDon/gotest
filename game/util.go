package game
import (
	"fmt"
	"math/rand"
)



func getRandomCol(cols int) int{
	fmt.Println("cols:",cols)
	return rand.Intn(cols)
}

func getLastSlotRow(col int,se sentinel) int {
	return se[col]

}
 
