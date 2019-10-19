package game
import "fmt"
func displayline(row []int){
	for _,value := range row{
		fmt.Print(value)
	}
	fmt.Println()
}