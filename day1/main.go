package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getData() []int {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var expense []int
	seperatedString := strings.Split(string(file), "\n")
	for _, number := range seperatedString {
		num, _ := strconv.Atoi(number)
		expense = append(expense, num)
	}
	return expense
}

func main() {
	expense := getData()
	// expense := []int{1721, 979, 366, 299, 675, 1456}

	for i, one := range expense {
		for j, two := range expense[i+1:] {
			if one+two < 2020 {
				for _, three := range expense[j+1:] {
					if one+two+three == 2020 {
						fmt.Println(one, two, three, one*two*three)
					}
				}
			}
		}
	}
}
