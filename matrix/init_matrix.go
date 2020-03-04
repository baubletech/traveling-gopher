package matrix

import (
	"fmt"
	"math/rand"
	"time"
)

type InitialMatrix [][]float64

func NewInitialMatrix(amount int, maxWeight float64) InitialMatrix {
	rand.Seed(time.Now().UnixNano())
	var result InitialMatrix
	for i := 0; i < amount; i++ {
		row := make([]float64, amount)
		for j := 0; j < amount; j++ {
			if j == i {
				row[j] = -1
			} else {
				row[j] = getRandom(maxWeight)
			}
		}
		result = append(result, row)
	}
	return result
}

//dummy 1111 table
func NewDummyInitialMatrix(amount int, maxWeight float64) InitialMatrix {
	rand.Seed(time.Now().UnixNano())
	var result InitialMatrix
	for i := 0; i < amount; i++ {
		row := make([]float64, amount)
		for j := 0; j < amount; j++ {
			if j == i {
				row[j] = -1
			} else if j == i+1 {
				row[j] = 1
			} else {
				row[j] = getRandom(maxWeight)
			}
		}
		result = append(result, row)
	}
	result[amount-1][0] = 1
	return result
}

func (m InitialMatrix) Show() {
	amount := len(m)
	fmt.Print("   |")
	for k := 1; k <= amount; k++ {
		fmt.Printf("%3d", k)
	}
	fmt.Print("\n")
	fmt.Print("_")
	for k := 0; k <= amount; k++ {
		fmt.Print("___")
	}
	fmt.Print("\n")
	for i := 0; i < amount; i++ {
		fmt.Printf("%3d", i+1)
		fmt.Print("|")
		for j := 0; j < amount; j++ {
			fmt.Printf("%3.0f", m[i][j])
		}
		fmt.Print("\n")
	}
}
