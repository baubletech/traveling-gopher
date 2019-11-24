package segment

import (
	"math/rand"
	"time"
)

// Segment is a basic rib in the Jobs graph
type Segment struct {
	ID         int        // Segment code
	Start, End int    	  // Start and End of segment. Basically - graph points
	Weight     float64    // Segment weight, for example, time to switch a job
	Links      []*Segment // Array of pointers to segments which we can travel to from the current one
	Disabled   bool		  // Is segment already processed
}

func New(id int, start int, end int, weight ...float64) Segment{
	segment := Segment{}
	segment.ID = id
	segment.Start = start
	segment.End = end
	if (len(weight) == 1) {
		segment.Weight = weight[0]
	} else if (len(weight) == 2) {
		segment.Weight = float64(rand.Intn(int(weight[1])))
	}
	return segment
}

func (segment *Segment) Reverse(id int, weight ...float64) Segment{
	if (len(weight) == 0) {
		weight = []float64{segment.Weight}
	}
	back := New(id, segment.End, segment.Start, weight...)
	return back
}

func (segment *Segment) AddLink(link *Segment) {
	segment.Links = append(segment.Links, link);
}

func GenerateMatrix(amount int, maxWeight float64) []Segment {
	rand.Seed(time.Now().UnixNano())
	var result []Segment
	var counter int = 1
	for j := 1; j <= amount - 1; j++ {
		for i := j + 1; i <= amount; i++ {
			segment := New(counter, j, i, 0, maxWeight)
			result = append(result, segment)
			counter++

			back := segment.Reverse(counter, 0, maxWeight)
			result = append(result, back)
			counter++
		}
	}

	for vid, value := range result {
		for _, next := range result {
			if (next.Start == value.End && next.End != value.Start) {
				result[vid].AddLink(&next);
			}
		}
	}

	return result
}