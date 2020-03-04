package matrix

import "fmt"

type Matrix []Segment

func New(m InitialMatrix) (result Matrix) {
	var counter int = 1
	var amount = len(m)

	for j := 1; j <= amount-1; j++ {
		for i := j + 1; i <= amount; i++ {
			segment := NewSegment(counter, j, i, m[j-1][i-1])
			result = append(result, segment)
			counter++

			back := segment.Reverse(counter, m[i-1][j-1])
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

func (m Matrix) Show() {
	var normalized [][]float64
	amount := len(m)

	for _, seg := range m {
		row := make([]float64, amount+1)
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
		fmt.Printf("%3d (%2d,%2d)", i+1, m[i].Start, m[i].End)
		fmt.Print("|")
		for j := 1; j <= amount; j++ {
			if (j - 1) == i {
				fmt.Print(" -1")
				continue
			}
			if normalized[i][j] == 0 {
				fmt.Print("   ")
			} else {
				fmt.Printf("%3.0f", normalized[i][j])
			}
		}
		fmt.Print("\n")
	}
}
