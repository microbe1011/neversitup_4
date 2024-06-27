package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countSmileys(strArr []string) int {
	testString := ""

	// for handle long text case with space inside text
	// if this case can ignore this code can use testString := strings.Join(strArr, "\t")
	for i := 0; i < len(strArr); i++ {
		if len(strArr[i]) <= 3 {
			testString += strings.Trim(strArr[i], " ") + "\t"
		}
	}

	r, _ := regexp.Compile("([:;][-~]?[)D])")

	match := r.FindAllString(testString, -1)

	fmt.Println("args: ", testString)
	fmt.Println("match: ", match)
	fmt.Println("count: ", len(match))

	return len(match)
}

func main() {
	countSmileys([]string{":)", ";(", ";}", ":-D"})       // should return 2;
	countSmileys([]string{";D", ":-(", ":-)", ";~)"})     // should return 3;
	countSmileys([]string{";]", ":[", ";*", ":$", ";-D"}) // should return 1;
	countSmileys([]string{""})                            // should return 0;
	countSmileys([]string{})                              // should return 0;
}
