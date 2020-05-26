package main

import (
	"fmt"
	"strings"
)

func solution(seoul []string) string {
	index := 0

	for i := 0; i < len(seoul); i++ {
		if seoul[i] == "Kim" {
			index = i
			break
		}
	}

	return fmt.Sprint("김서방은 ", index, "에 있다")
}

func solution2(s string) string {
	index := 0
	var resultString string
	for i := 0; i < len(s); i++ {

		str := s[i : i+1]

		if str == " " {
			index = 0
		} else if index%2 == 0 {
			str = strings.ToUpper(str)
			index++

		} else {
			str = strings.ToLower(str)
			index++

		}

		resultString = resultString + str
	}

	return resultString
}

func main() {
	fmt.Println(solution2("try hello world"))
}
