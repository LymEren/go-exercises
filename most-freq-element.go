package main

import "fmt"

func enCokTekrar(dizi []int) int {

	var makSayici int
	var tekrar int
	deneme := map[int]int{}

	for _, value := range dizi {
		deneme[value]++
		if deneme[value] > makSayici {
			makSayici = deneme[value]
			tekrar = value
		}
	}
	return tekrar

}

func main() {

	dizim := []int{5, 5, 6, 6, 6, 7, 8, 8, 8, 8}
	fmt.Println(enCokTekrar(dizim))

}
