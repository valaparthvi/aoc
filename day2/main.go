package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getData() []string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	seperatedString := strings.Split(string(file), "\n")
	return seperatedString
}

func main() {
	var policy, limit []string
	var password, subject string
	var position1, position2, subjectCount, counter int
	// inputs := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	inputs := getData()
	for _, input := range inputs {
		getPolicyPW := strings.SplitAfter(input, ": ")
		policy, password = strings.Split(getPolicyPW[0], " "), strings.TrimSpace(getPolicyPW[1])
		limit, subject = strings.Split(policy[0], "-"), strings.Trim(policy[1], ": ")
		position1, _ = strconv.Atoi(limit[0])
		position2, _ = strconv.Atoi(limit[1])

		subjectCount = strings.Count(password, subject)
		position1char := string(password[position1-1])
		position2char := string(password[position2-1])
		if subjectCount == 1 {
			if position1char == subject || position2char == subject {
				counter++
			}
		} else if subjectCount > 1 {
			if (position1char == subject && position2char != subject) || (position1char != subject && position2char == subject) {
				counter++
			}
		}
	}
	fmt.Println(counter)

}
