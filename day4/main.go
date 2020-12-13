package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func getData() []string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	seperatedString := strings.Split(string(file), "\n\n")
	return seperatedString
}

func main() {
	var valid bool
	var validCount int
	data := getData()
	fmt.Println(len(data))
	keys := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "cid", "hgt"}

	for index, da := range data {
		for _, key := range keys {
			if strings.Contains(da, key) {
				valid = true
			} else if key == "cid" {
				valid = true
			} else {
				valid = false
				break
			}
		}
		if valid{
			validCount++
			fmt.Println("Passport", index+1, "is valid: ", valid)
		}
	}
	fmt.Println(validCount)
}
