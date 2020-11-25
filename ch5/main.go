package ch5

import "fmt"

func main() {
	var primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes[0], primes[2])

	primes2 := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes2)

	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	for index, note := range notes {
		fmt.Println(index, note)
	}
	for _, note := range notes {
		fmt.Println(note)
	}
}
