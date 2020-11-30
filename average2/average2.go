package average2

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Average2() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

func Average3() {
	argments := os.Args[1:]
	var numbers []float64
	for _, argument := range argments {
		argument, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, argument)
	}

	sampleCount := float64(len(numbers))
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
