package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Hi, i'm Eyyub"
	fmt.Println("Sentence 1: ", str1)
	str2 := "This is my Repo"
	fmt.Println("Sentence 2: ", str2)

	// We are using Count() function for count
	// strings package

	count1 := strings.Count(str1, "i")
	fmt.Println("Sentence 1 Count: ", count1)
	count2 := strings.Count(str2, " ")
	fmt.Println("Sentence 2 Count: ", count2)

}
