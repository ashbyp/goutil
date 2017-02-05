package homework

import (
	"fmt"
	"github.com/ashbyp/goutil/mathutil"
)

func GetDecreasing(n int) (numbers []int, err error) {
	if n > 9 {
		return numbers, fmt.Errorf("max numbers can be is 9")
	}
	var a = make([]int, n)

	// how

	return a, nil

}

func GetAllDecreasing3() (numbers []int) {
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				numbers = append(numbers, 100*i+10*j+k)
			}
		}
	}
	return
}

func GetAllDecreasing4() (numbers []int) {
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				for l := k - 1; l >= 0; l-- {
					numbers = append(numbers, 1000*i+100*j+10*k+l)
				}
			}
		}
	}
	return
}

func GetAllDecreasing5() (numbers []int) {
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				for l := k - 1; l >= 0; l-- {
					for m := l - 1; m >= 0; m-- {
						numbers = append(numbers, 10000*i+1000*j+100*k+10*l+m)
					}
				}
			}
		}
	}
	return
}

func GetAllDecreasing6() (numbers []int) {
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				for l := k - 1; l >= 0; l-- {
					for m := l - 1; m >= 0; m-- {
						for n := m - 1; n >= 0; n-- {
							numbers = append(numbers, 100000*i+10000*j+1000*k+100*l+10*m+n)
						}
					}
				}
			}
		}
	}
	return
}

func DoNumberTrick(numbers []int) (answers map[int]bool) {
	answers = make(map[int]bool)
	for _, n := range numbers {
		nv := mathutil.ReverseInt(n)
		var s int
		if n > nv {
			s = n - nv
		} else {
			s = nv - n
		}
		a := s + mathutil.ReverseInt(s)
		answers[a] = true
	}
	return
}
