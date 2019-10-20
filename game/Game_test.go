package game
import(
	"testing"
	"fmt"
)

func TestRandomCol(t *testing.T){
	i := 0
	for i<10 {
		j :=getRandomCol(7)
		if j>= 7 {
			t.Error("error")
		}
	}
}

