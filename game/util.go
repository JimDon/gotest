package game
import (
	"fmt"
	"math/rand"
)
func displayline(row []int){
	for _,value := range row{
		fmt.Print(value)
	}
	fmt.Println()
}


func getRandomCol(colWidth int) int{
	return rand.Intn(colWidth)
}

func getLastSlotRow(col int,se sentinel) int {
	return se[col]

}
 
