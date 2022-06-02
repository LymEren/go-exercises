package main

import (
	"bytes"
	"fmt"
)

func main() {

	slice1 := []byte{'E', 'Y', 'Y', 'U', 'B'}
	slice2 := []byte{'G', 'O', 'O', 'O', 'L', 'A', 'N', 'G'}

	// Count function for bytes (bytes package)

	example1 := bytes.Count(slice1, []byte{'Y'})
	fmt.Println("First: ", example1)

	example2 := bytes.Count(slice2, []byte{'O'})
	fmt.Println("Second: ", example2)

}
