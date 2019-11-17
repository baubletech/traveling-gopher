package main

import "fmt"

// Segment is a basic rib in the Jobs graph
type Segment struct {
	ID         int        // Segment code
	Start, End string     // Start and End of segment. Basically - graph points
	Weight     float64    // Segment weight, for example, time to switch a job
	Links      []*Segment // Array of pointers to segments which we can travel to from the current one
}

// PrepareExampleSegments prepares example segments slice
func PrepareExampleSegments() []Segment {
	var segments []Segment

	segments = append(segments, Segment{
		ID:     1,
		Start:  "a1",
		End:    "a2",
		Weight: 3,
		Links:  make([]*Segment, 2),
	})

	segments = append(segments, Segment{
		ID:     9,
		Start:  "a2",
		End:    "a3",
		Weight: 4,
		Links:  make([]*Segment, 2),
	})

	return segments
}

func main() {
	segments := PrepareExampleSegments()
	fmt.Println(segments)
}
