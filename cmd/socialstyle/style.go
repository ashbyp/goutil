package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func readfile(filename string) (questions []string, err error) {
	var f *os.File
	if f, err = os.Open(filename); err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			questions = append(questions, scanner.Text())
		}
	}
	return
}

func response(question string) (response int, err error) {
	reader := bufio.NewReader(os.Stdin)
	var text string
	for {
		fmt.Printf("\n%s\n\n", question)
		fmt.Printf("Enter response[1 2 3 4 5 6 7] ..> ")
		if text, err = reader.ReadString('\n'); err != nil {
			fmt.Println("Try again")
		} else {
			text = strings.TrimRight(text, "\n")
			if len(text) > 0 && (text[0] == 'q' || text[0] == 'Q') {
				return 0, err
			}
			if response, err = strconv.Atoi(text); err != nil {
				fmt.Println("Try again")
			} else {
				if response >= 1 && response <= 7 {
					break
				} else {
					fmt.Println("Try again")
				}
			}
		}
	}
	return
}

func processQuestions(questions []string) {
	var dr, ex, an, am, ve, nv int

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Please answer each question on a scale of 1 to 7")
	fmt.Println("")
	fmt.Println("  1 = rarely do I exhibit this behaviour")
	fmt.Println("  4 = exhibit this behaviour from time to time")
	fmt.Println("  7 = exhibit this behaviour very often")
	fmt.Println("")
	fmt.Println(" Of course other ratings apply as well")
	fmt.Println("  2 and 3 represent varying degrees between average and little")
	fmt.Println("  5 and 6 represent varying degrees between average and often")
	fmt.Println("")
	fmt.Println("Type 'q' to quit at ant time ")
	fmt.Println("")
	fmt.Println("--------------------------------------------------------------")

	for _, line := range questions {
		q := strings.Split(line, "|")
		r, _ := response(q[0])
		if r == 0 {
			fmt.Println("You chose to quit, bye")
			return
		} else {
			switch strings.ToUpper(q[1]) {
			case "DR":
				dr = dr + r
			case "EX":
				ex = ex + r
			case "AN":
				an = an + r
			case "AM":
				am = am + r
			case "VE":
				ve = ve + r
			case "NV":
				nv = nv + r
			default:
				panic("Invalid question type " + q[0] + " " + q[1])
			}
		}
	}

	fmt.Printf("Social styles Driver=%d, Expressive=%d, Analytical=%d, Amiable=%d\n", dr-2, ex+1, an+2, am-6)
	fmt.Printf("Versatility=%d", ve-nv)
}

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Exiting...")
		os.Exit(1)
	}()

	var filename string
	if len(os.Args) != 2 {
		filename = os.Getenv("GODATA") + string(filepath.Separator) + "questions.txt"
	} else {
		filename = os.Args[1]
	}
	fmt.Printf("Will read questions from %s\n", filename)
	if questions, err := readfile(filename); err == nil {
		fmt.Printf("Read %d questions\n", len(questions))
		processQuestions(questions)
	} else {
		fmt.Println(err)
	}
}
