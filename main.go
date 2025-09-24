package main

import (
	"fmt"
	"time"
)

func RechercheDichotomique(tab []int, n int) {
	temps := time.Now()
	debut := 0
	fin := len(tab) - 1
	for debut <= fin {
		index_milieu := (debut + fin) // 2
		if tab[index_milieu] == n {
			fmt.Println("Trouvé à l'index", index_milieu)
			dure2 := time.Since(temps)
			fmt.Println("Le temps d'éxécution est de", dure2)
			return
		} else if tab[index_milieu] < n {
			debut = index_milieu + 1
		} else {
			fin = index_milieu - 1
		}
	}
	fmt.Println("Non trouvé")
	dure := time.Since(temps)
	fmt.Println("le temps d'éxécution est de", dure)
}

func main() {
	tab := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println("Test 1 : élément présent")
	RechercheDichotomique(tab, 7)

	fmt.Println("Test 2 : élément absent")
	RechercheDichotomique(tab, 8)

	fmt.Println("Test 3 : premier élément")
	RechercheDichotomique(tab, 1)

	fmt.Println("Test 4 : dernier élément")
	RechercheDichotomique(tab, 15)

	fmt.Println("Test 5 : tableau vide")
	RechercheDichotomique([]int{}, 5)
}
