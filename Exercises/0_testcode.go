package main

import "fmt"

func main() {

	mapOrnek := map[string]string{
		"Brand":  "Ford",
		"Name":   "Ali",
		"Zorluk": "Evet"}

	fmt.Printf("%v", mapOrnek["Brand"])

}
