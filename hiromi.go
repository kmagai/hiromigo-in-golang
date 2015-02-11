package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter integer: ")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fizzbuzz(i)
	}
}

func isFiveIncluded(num int) bool {
	var strInt = strconv.Itoa(num)
	for k := 0; k < len([]rune(strInt)); k++ {
		if '5' == strInt[k] {
			return true
		}
	}
	return false
}

func fizzbuzz(i int) {
	exotic := "Exotic"
	japan := "Japan"

	if isFiveIncluded(i) && i%5 == 0 {
		fmt.Println(i, exotic+" "+japan)
	} else if isFiveIncluded(i) {
		fmt.Println(i, exotic)
	} else if i%5 == 0 {
		fmt.Println(i, japan)
	} else {
		fmt.Println(i)
	}
}
