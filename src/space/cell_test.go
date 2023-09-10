package space_test

import "testing"
import "github.com/jprajwal/birdgame/space"
import "fmt"


func TestCellStringerInterfaceForInt(t * testing.T) {
	cell := space.NewCell(1)
	if cell.String() != "1" {
		t.Fatalf("Cell failed to hold integer")
	}
}

type MyStringer int
func (ms MyStringer) String() string {
	return fmt.Sprintf("%v", int(ms))
}

func TestCellStringerInterfaceForCustomStringer(t * testing.T) {
	cell := space.NewCell(MyStringer(1))
	if cell.String() != "1" {
		t.Fatalf("Cell failed to hold a custom stringer")
	}
}

type MyNonStringer int

func TestCellStringerInterfaceForNonStringer(t * testing.T) {
	cell := space.NewCell(MyNonStringer(1))
	if cell.String() != "" {
		t.Fatalf("Cell failed to hold a non-stringer")
	}
}
