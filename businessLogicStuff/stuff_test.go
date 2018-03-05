package businessLogicStuff

import (
	"fmt"
	"testing"
)

type asdf struct {
	name  string
	alter int
}

type jkl struct {
	asdf
	bla  string
	name string
}

func (a asdf) printName() {
	fmt.Println(a.name)
}

func TestDoStuff(t *testing.T) {
	a := asdf{"benni", 32}
	a.printName()
	j := jkl{asdf{"stefan", 45}, "cool", "falscher name"}
	fmt.Println(j.asdf.name)
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1 test", "1", 1},
		{"2 test", "2", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoStuff(tt.s); got != tt.want {
				t.Errorf("DoStuff() = %v, want %v", got, tt.want)
			}
		})
	}
}
