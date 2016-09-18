package fileutil

import (
	"testing"
	"fmt"
)


func TestGetAllDecreasing3(t *testing.T) {
	fmt.Println(GetAllDecreasing3())
}

func TestReverseInt(t *testing.T) {
	fmt.Println(reverseInt(123))
	fmt.Println(reverseInt(9001))
	fmt.Println(reverseInt(123456789))
	fmt.Println(reverseInt(1))
	fmt.Println(reverseInt(3333))
}

func TestDoNumberTrick(t *testing.T) {
	fmt.Println("Answers for 3 digits")
	fmt.Println(DoNumberTrick(GetAllDecreasing3()))
	fmt.Println("Answers for 4 digits")
	fmt.Println(DoNumberTrick(GetAllDecreasing4()))

}
