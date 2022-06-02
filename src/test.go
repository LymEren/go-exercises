/*

//MAPPING

package main

import "fmt"

func main() {

	mapOrnek := map[string]string{
		"Brand":  "Ford",
		"Name":   "Ali",
		"Zorluk": "Evet"}
		"Yeni" : "urun"
	fmt.Printf("%v", mapOrnek["Brand"])

} */

// STRUCT

/* package mail

type team struct {
	name   string
	played int
	won    int
	lost   int
	drwan  int
}
*/

// POINTER

/* package main

import "fmt"

func main() {

	x := 22
	fmt.Println(x) // 22
	fmt.Println(*x)
	fmt.Println(&x)       // 22 nin adresi &x ---> *int
	fmt.Println(*(&x))    // dereferencing
	fmt.Println(&(*(&x))) // * --->> ilgili adresteki değeri gösteriri
	fmt.Println(*(&(*(&x))))
	fmt.Println(3 * 5)
	return
} */

// Part of array

/*
package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[2:4])
}
*/