package research

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnsignedOverflow(t *testing.T) {
	var a uint32 = 2147483648
	var b uint32 = 2147483648
	var sum = a + b
	fmt.Println(reflect.TypeOf(sum))
	fmt.Printf("Sum is : %d", sum)
	//uint32
	//Sum is : 0
}

func TestSignedOverflow(t *testing.T) {
	var a int8 = 127
	var b int8 = 1
	var sum = a + b
	fmt.Println(reflect.TypeOf(sum))
	fmt.Printf("Sum is : %d", sum)
	//int8
	//-128
}

// https://bycsec.top/2021/02/07/golang%E7%9A%84%E4%B8%80%E4%BA%9B%E5%AE%89%E5%85%A8%E9%97%AE%E9%A2%98/
