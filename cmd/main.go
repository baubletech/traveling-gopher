package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/baubletech/traveling-gopher/matrix"
)

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}

func getIntFromUser(prompt string) int {
	return int(getFloat64FromUser(prompt))
}

func getFloat64FromUser(prompt string) float64 {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	// Size of matrix
	n := getIntFromUser("Enter matrix size: ")
	bound := getFloat64FromUser("Percent of average weight bound ?: ")

	if a := getUserInput("Would you like to run testing? (y/n): "); a == "y" {
		times := getIntFromUser("How many iterations?: ")
		matrix.TestAlgo(n, bound, times)
	} else {
		matrix.OutputFullReport(n, bound)
	}
}
