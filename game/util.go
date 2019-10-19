package game
import "fmt"
func displayline(row []int){
	for _,value := range row{
		fmt.Print(value)
	}
	fmt.Println()
}
type coordinate struct {
	x int
	y int
}
func getNextStep