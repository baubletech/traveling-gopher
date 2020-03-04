package matrix

import (
	"math/rand"
)

// Segment is a basic rib in the Jobs graph
type Segment struct {
	ID         int        // Segment code
	Start, End int        // Start and End of segment. Basically - graph points
	Weight     float64    // Segment weight, for example, time to switch a job
	Links      []*Segment // Array of pointers to segments which we can travel to from the current one
	Disabled   bool       // Is segment already processed
}

// BestSegment is a segment with minimal weight overall value
type BestSegment struct {
	Weight   float64
	Segments []*Segment
}

func NewSegment(id int, start int, end int, weight ...float64) Segment {
	segment := Segment{}
	segment.ID = id
	segment.Start = start
	segment.End = end
	if len(weight) == 1 {
		segment.Weight = weight[0]
	} else if len(weight) == 2 {
		segment.Weight = getRandom(weight[1])
	}
	return segment
}

func (segment *Segment) Reverse(id int, weight ...float64) Segment {
	if len(weight) == 0 {
		weight = []float64{segment.Weight}
	}
	back := NewSegment(id, segment.End, segment.Start, weight...)
	return back
}

func (segment *Segment) AddLink(link *Segment) {
	segment.Links = append(segment.Links, link)
}

func getRandom(maxWeight float64) float64 {
	return float64(rand.Intn(int(maxWeight))) + 1
}
