package main

import "fmt"

func main() {
	// bases := map[string]number{
	// 	"A": 0,
	// 	"C": 0,
	// 	"G": 0,
	// 	"T": 0,
	// }
	counter := 0
	adn := []string{"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"}
	for _, seq := range adn {
		if string(seq[0]) == string(seq[1]) && string(seq[1]) == string(seq[2]) && string(seq[2]) == string(seq[3]) && string(seq[3]) == string(seq[4])  {
		counter += 1
	} else if string(seq[1]) == string(seq[2]) && string(seq[2]) == string(seq[3]) && string(seq[3]) == string(seq[4]) && string(seq[4]) == string(seq[5]) {
		counter += 1
	} else if string(seq[2]) == string(seq[3]) && string(seq[3]) == string(seq[4]) && string(seq[4]) == string(seq[5]) {
		counter += 1
	}
}
fmt.Println("Mutation counter", counter)
// fmt.Println("Mutation counter", string(adn[1][1]), string(adn[1][2]), string(adn[1][3]))
	// fmt.Println(string(adn[0][0]))
}