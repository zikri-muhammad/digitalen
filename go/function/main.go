// ini pembelajaran function

package main

import "fmt"

func tambah(a, b int) int {
	return a + b
}

func kurang(c, d int) int {
	return c - d
}

func perkalian(e, f int) int {
	return e * f
}

func main() {
	a := tambah(5, 6)
	fmt.Println("hasil dari fuction tambah", a)

	b := kurang(9, 5)
	fmt.Println("hasil dari fuction tambah", b)

	c := perkalian(5, 6)
	fmt.Println("hasil dari fuction tambah", c)

}
