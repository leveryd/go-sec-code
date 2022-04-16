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
	// a = append(a, 4)

	b := append(a, 5)
	c := append(a, 6)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//result:
	//[1 2 3]
	//[1 2 3 6]
	//[1 2 3 6]
}

// https://bycsec.top/2021/02/07/golang%E7%9A%84%E4%B8%80%E4%BA%9B%E5%AE%89%E5%85%A8%E9%97%AE%E9%A2%98/
