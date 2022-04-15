package research

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := make([]uint64, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)

	b := append(a, 5)
	c := append(a, 6)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//result:
	//[1 2 3]
	//[1 2 3 5]
	//[1 2 3 5]
}
