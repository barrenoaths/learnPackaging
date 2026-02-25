package main

import (
	"fmt"
	"learnPackaging/SumFourNums"
	"learnPackaging/SumThreeNums"
)

func main() {
	fmt.Println("test")
	fmt.Println(SumThreeNums.SumThree(1, 2, 3))
	fmt.Println(SumFourNums.SumFour(1, 2, 3, 4))
}
