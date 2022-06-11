package main

import "fmt"

var newRna = map[rune]byte{
	'A': 'U',
	'T': 'A',
	'C': 'G',
	'G': 'C',
}

func rnaConv(dna string) string {
	rna := make([]byte, len(dna))
	for index, value := range dna {
		rna[index] = newRna[value]
	}
	return string(rna)
}

func main() {
	names := "AGTC AATC"
	fmt.Println(rnaConv(names))
}
