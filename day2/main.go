package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)

func getData() []string{
	file, err := ioutil.ReadFile("input.txt")
	if err!=nil {
		panic(err)
	}
	seperatedString := strings.Split(string(file), "\n")
	return seperatedString
}

func main() {
	var policy, limit []string
	var password, subject string
	var min, max, subjectCount, counter int
	// inputs := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	inputs := getData()
	for _, input := range inputs {
		getPolicyPW := strings.SplitAfter(input, ": ")
		policy, password = strings.Split(getPolicyPW[0], " "), strings.TrimSpace(getPolicyPW[1])
		limit, subject = strings.Split(policy[0], "-"), strings.Trim(policy[1], ": ")
		min, _ = strconv.Atoi(limit[0])
		max, _ = strconv.Atoi(limit[1])

		subjectCount = strings.Count(password, subject)
		if subjectCount >= min && subjectCount <= max {
			counter++
		}
	}
	fmt.Println(counter)

}
