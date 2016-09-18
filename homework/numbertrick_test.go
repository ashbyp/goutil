package fileutil

import (
	"testing"
	"fmt"
)


func TestGetAllDecreasing3(t *testing.T) {
	fmt.Println(GetAllDecreasing3())
}

func TestDoNumberTrick(t *testing.T) {
	fmt.Println("Answers for 3 digits")
	fmt.Println(DoNumberTrick(GetAllDecreasing3()))
	fmt.Println("Answers for 4 digits")
	fmt.Println(DoNumberTrick(GetAllDecreasing4()))
	fmt.Println("Answers for 5 digits")
	fmt.Println(DoNumberTrick(GetAllDecreasing5()))

}

func TestGetDecreasing(t *testing.T) {
	_, err := GetDecreasing(10)
	if err == nil {
		t.Error("GetDecreasing accepted 10")
	}
}
