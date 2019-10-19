package game
import(
	"testing"
)

func TestDisplay(t *testing.T){
	var g GameClass
	if g.Display() == 3{
		t.Error("error occued")
	}
}

