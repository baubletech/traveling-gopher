package matrix

import "fmt"

func getAverageWeight(initial [][]float64) float64 {
	var sum float64
	for _, i := range initial {
		for _, j := range i {
			if j != -1 {
				sum += j
			}
		}
	}

	return sum / float64(len(initial)*len(initial)-len(initial))
}

func isStartInSegments(segments []*Segment, start int) bool {
	for _, v := range segments {
		if v.Start == start {
			return true
		}
	}
	return false
}

// Check segment comparing to average
func checkSegment(segment *Segment, segments []*Segment, averageWeight float64, n int) (bool, []*Segment) {
	if n <= 0 {
		return true, segments
	}

	if isStartInSegments(segments, segment.End) {
		return false, segments
	}

	for _, seg := range segment.Links {
		if seg.Weight < averageWeight {
			res, segs := checkSegment(seg, append(segments, seg), averageWeight, n-1)
			if res {
				return true, segs
			}
		}
	}

	return false, segments
}

// Start segment check comparing to average weight
func buildChainByAverage(dotsSize int, segments []Segment, start int, averageWeight float64) (result []*Segment) {
	var dots []int
	// Get a starting segment
	seg := &segments[start]
	segLinks := seg.Links
	//result = append(result, seg)
	dots = append(dots, seg.Start)

	var freshSegs []*Segment
	freshSegs = append(freshSegs, seg)

	for _, value := range segLinks {
		if value.Weight < averageWeight {
			res, segs := checkSegment(value, append(freshSegs, value), averageWeight, dotsSize-3)
			if res {
				return segs
			}
		}
	}

	return nil
}

func getChainWeight(segments []*Segment) (sum float64) {
	for _, seg := range segments {
		sum += seg.Weight
	}

	return sum
}

// Check segment using full traversal
func checkSegmentFull(segment *Segment, segments []*Segment, best *BestSegment, n int) []*Segment {
	if n <= 0 {
		// Calculate current chain weight & get back to previous call
		weightSum := getChainWeight(segments)
		// Replace best with current if current weight is lower or if no segments are defined as best yet
		if (weightSum < best.Weight) || (best.Segments == nil) {
			best.Weight = weightSum
			best.Segments = segments
		}

		return segments
	}

	// Check that current segment won't loop until needed
	if isStartInSegments(segments, segment.End) {
		return segments
	}

	for _, seg := range segment.Links {
		checkSegmentFull(seg, append(segments, seg), best, n-1)
	}

	return segments
}

func buildChainFullTraversal(dotsSize int, segments []Segment, best *BestSegment, start int) *BestSegment {
	var dots []int
	// Get a starting segment
	seg := &segments[start]
	segLinks := seg.Links
	//result = append(result, seg)
	dots = append(dots, seg.Start)

	var freshSegs []*Segment
	freshSegs = append(freshSegs, seg)

	for _, value := range segLinks {
		checkSegmentFull(value, append(freshSegs, value), best, dotsSize-3)
	}

	return best
}

func OutputFullReport(n int, averagePercentBound float64) {
	initial := NewInitialMatrix(n, 10.0)
	segments := New(initial)

	initial.Show()
	segments.Show()

	// Get average weight
	averageWeight := getAverageWeight(initial)
	fmt.Println("Average element weight is:", averageWeight)

	// Get overall average
	fmt.Println("Overall matrix average is:", averageWeight*float64(n-1))

	// Build with average
	var builtSegments []*Segment
	for i := 0; i < len(segments); i++ {
		builtSegments = buildChainByAverage(n, segments, i, averageWeight*averagePercentBound/100)

		if builtSegments == nil {
			continue
		} else {
			break
		}
	}

	fmt.Println("\nResult found:")

	var timeOverall float64
	for _, value := range builtSegments {
		timeOverall += value.Weight
		fmt.Println(value.Start, "->", value.End)
	}
	fmt.Println("Overall built chain weight:", timeOverall)

	// Build full traversal
	best := &BestSegment{}
	for i := 0; i < len(segments); i++ {
		buildChainFullTraversal(n, segments, best, i)
	}

	fmt.Println("\nResult full traversal found:")

	for _, value := range best.Segments {
		fmt.Println(value.Start, "->", value.End)
	}
	fmt.Println("Overall built chain weight:", best.Weight)
}

func TestAlgo(n int, averagePercentBound float64, times int) {
	diffPercent := 15.0
	maxWeight := 10.0
	successfulTimes := times

	var successful int
	var averageChangeSum float64

	for i := 0; i < times; i++ {
		initial := NewInitialMatrix(n, maxWeight)
		segments := New(initial)

		// Get average weight
		averageWeight := getAverageWeight(initial)

		// Build with average
		var builtSegments []*Segment
		for i := 0; i < len(segments); i++ {
			builtSegments = buildChainByAverage(n, segments, i, averageWeight*averagePercentBound/100)

			if builtSegments == nil {
				continue
			} else {
				break
			}
		}

		var timeOverall float64
		for _, value := range builtSegments {
			timeOverall += value.Weight
		}

		// Build full traversal
		best := &BestSegment{}
		for i := 0; i < len(segments); i++ {
			buildChainFullTraversal(n, segments, best, i)
		}

		// Check if close to opt
		change := percentageChange(best.Weight, timeOverall)

		if (change >= 0) && (change <= diffPercent) {
			successful++
		}

		if change < 0 {
			change = 0
			successfulTimes--
		}
		averageChangeSum += change
	}

	fmt.Println("Average change is:", averageChangeSum/float64(successfulTimes))
	fmt.Println("Times:", successfulTimes, "Successful (15%):", successful)
	fmt.Println("Success percentage:", (float64(successful)/float64(successfulTimes))*100.0)
}

func percentageChange(old, new float64) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}
