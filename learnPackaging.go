package main

import (
	"fmt"
	"learnPackaging/SumFourNums"
	"learnPackaging/SumThreeNums"
	"learnPackaging/utils"
)

func main() {
	fmt.Println("test")
	fmt.Println(SumThreeNums.SumThree(1, 2, 3))
	fmt.Println(SumFourNums.SumFour(1, 2, 3, 4))
	fmt.Println(utils.SumFiveFunc(1, 2, 3, 4, 5))
	fmt.Println(utils.SumSixFunc(1, 2, 3, 4, 5, 6))
}
