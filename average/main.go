package main

import (
	"fmt"
	"log"

	"github.com/kadragon/headfirstgo/datafile"
)

func main() {
	// numbers := [3]float64{71.8, 56.2, 89.5}
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64

	for _, value := range numbers {
		sum += value
	}
	fmt.Println(sum)

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
