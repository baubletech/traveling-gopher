package segment

import (
	"fmt"
	"math/rand"
	"time"
)

// Segment is a basic rib in the Jobs graph
type Segment struct {
	ID         int        // Segment code
	Start, End int        // Start and End of segment. Basically - graph points
	Weight     float64    // Segment weight, for example, time to switch a job
	Links      []*Segment // Array of pointers to segments which we can travel to from the current one
	Disabled   bool       // Is segment already processed
}

func New(id int, start int, end int, weight ...float64) Segment {
	segment := Segment{}
	segment.ID = id
	segment.Start = start
	segment.End = end
	if len(weight) == 1 {
		segment.Weight = weight[0]
	} else if len(weight) == 2 {
		segment.Weight = GetRandom(weight[1])
	}
	return segment
}

func (segment *Segment) Reverse(id int, weight ...float64) Segment {
	if len(weight) == 0 {
		weight = []float64{segment.Weight}
	}
	back := New(id, segment.End, segment.Start, weight...)
	return back
}

func (segment *Segment) AddLink(link *Segment) {
	segment.Links = append(segment.Links, link)
}

func GetRandom(maxWeight float64) float64 {
	return float64(rand.Intn(int(maxWeight))) + 1
}

func GenerateInitialMatrix(amount int, maxWeight float64) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	var result [][]float64
	for i := 0; i < amount; i++ {
		row := make([]float64, amount)
		for j := 0; j < amount; j++ {
			if j == i {
				row[j] = -1
			} else {
				row[j] = GetRandom(maxWeight)
			}
		}
		result = append(result, row)
	}
	return result
}

func ShowInitialMatrix(initalMatrix [][]float64) {
	amount := len(initalMatrix)
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
		fmt.Printf("%3d", i + 1)
		fmt.Print("|")
		for j := 0; j < amount; j++ {
			fmt.Printf("%3.0f",initalMatrix[i][j])
		}
		fmt.Print("\n")
	}
}

func ShowMatrix(segments []Segment) {
	var normalized [][]float64
	amount := len(segments)

	for _, seg := range segments {
		row := make([]float64, amount + 1)
		for _, link := range seg.Links {
			row[link.ID] = link.Weight 
		}
		normalized = append(normalized, row)
	}

	fmt.Print("           |")
	for k := 1; k <= amount; k++ {
		fmt.Printf("%3d", k)
	}
	fmt.Print("\n")
	fmt.Print("_________")
	for k := 0; k <= amount; k++ {
		fmt.Print("___")
	}
	fmt.Print("\n")
	for i := 0; i < amount; i++ {
		fmt.Printf("%3d (%2d,%2d)", i + 1, segments[i].Start, segments[i].End)
		fmt.Print("|")
		for j := 1; j <= amount; j++ {
			if ((j - 1) == i) {
				fmt.Print(" -1")
				continue
			}
			if (normalized[i][j] == 0) {
				fmt.Print("   ")
			} else {
				fmt.Printf("%3.0f",normalized[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func GenerateMatrix(initalMatrix [][]float64) []Segment {
	var result []Segment
	var counter int = 1
	var amount = len(initalMatrix)

	for j := 1; j <= amount-1; j++ {
		for i := j + 1; i <= amount; i++ {
			segment := New(counter, j, i, initalMatrix[j-1][i-1])
			result = append(result, segment)
			counter++

			back := segment.Reverse(counter, initalMatrix[i-1][j-1])
			result = append(result, back)
			counter++
		}
	}

	for vid, value := range result {
		for nid, next := range result {
			if next.Start == value.End && next.End != value.Start {
				result[vid].AddLink(&result[nid])
			}
		}
	}

	return result
}
