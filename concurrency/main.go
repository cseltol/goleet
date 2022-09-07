package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSE = false
var DATA = make(map[int]bool)


func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func f(min, max int, out chan<- int) {
	for {
		if CLOSE {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func s(out chan<- int, in <-chan int) {
	for i := range in {
		fmt.Print(i, " ")
		if _, ok := DATA[i]; ok {
			CLOSE = true
		} else {
			DATA[i] = true
			out <- i
		}
	}
	fmt.Println()
	close(out)
}

func t(in <-chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters")
		return
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
	}

	rand.Seed(time.Now().UnixNano())

	A, B := make(chan int), make(chan int)
	go f(n1, n2, A)
	go s(B, A)
	t(B)
}
