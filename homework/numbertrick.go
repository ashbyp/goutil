package fileutil

func GetAllDecreasing3() (numbers []int) {	
	for i := 9; i >= 0; i-- {
		for j := i-1; j >= 0; j-- {
			for k := j-1; k >=0; k-- {
				numbers = append(numbers, 100*i+10*j+k)
			}
		}
	}
	return
}	

func GetAllDecreasing4() (numbers []int) {	
	for i := 9; i >= 0; i-- {
		for j := i-1; j >= 0; j-- {
			for k := j-1; k >=0; k-- {
				for l := k-1; l >= 0 ; l-- {
					numbers = append(numbers, 1000*i+100*j+10*k + l)
				}
			}
		}
	}
	return
}	

func DoNumberTrick(numbers []int) (answers map[int]bool) {
	answers = make(map[int]bool)
	for _, n := range numbers {
		nv := reverseInt(n)
		var s int
		if n > nv {
			s = n - nv
		} else {
			s = nv - n
		}
		a := s + reverseInt(s)
		answers[a] = true
	}
	return
}


func reverseInt(i int) (j int) {
	for ; i != 0; {
		d := i%10
        i = i / 10
        j = j * 10 + d
	}
	return
}
