package main

import (
	"fmt"

	"github.com/baubletech/traveling-gopher/segment"
)

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func buildChain(dotsSize int, segments []segment.Segment, start int) (result []*segment.Segment) {
	var dots []int
	// Get a starting segment
	seg := &segments[start]
	segLinks := seg.Links
	result = append(result, seg)
	dots = append(dots, seg.Start)

	for {
		// Check links are enabled, break otherwise
		var availableLinks []*segment.Segment
		for _, value := range segLinks {
			if (!value.Disabled) {
				if (len(dots) < (dotsSize - 1)) {
					if (!intInSlice(value.End, dots)) {
						availableLinks = append(availableLinks, value)
					}
				}
			}
		}

		if len(availableLinks) == 0 {
			break
		}

		// Choose correct path + disable them
		minWeightSeg := availableLinks[0]
		for _, value := range availableLinks {
			value.Disabled = true
			if minWeightSeg.Weight > value.Weight {
				minWeightSeg = value
			}
		}

		// Get to next linked segment
		seg = minWeightSeg
		segLinks = seg.Links
		result = append(result, seg)
		dots = append(dots, seg.Start)
	}

	return result
}

func getAverageWeight(initial [][]float64) float64 {
	var sum float64
	for _, i := range initial {
		for _, j := range i {
			if j != -1 {
				sum += j
			}
		}
	}

	return sum / float64(len(initial) * len(initial) - len(initial))
}

func isStartInSegments(segments []*segment.Segment, start int) bool {
	for _, v := range segments {
		if v.Start == start {
			return true
		}
	}
	return false
}

func checkSegment(segment *segment.Segment, segments []*segment.Segment, averageWeight float64, n int) (bool, []*segment.Segment) {
	if n <= 0 {
		return true, segments
	}

	if isStartInSegments(segments, segment.End) {
		return false, segments
	}

	for _, seg := range segment.Links {
		if seg.Weight < averageWeight {
			res, segs := checkSegment(seg, append(segments, seg), averageWeight, n - 1)
			if res {
				return true, segs
			}
		}
	}

	return false, segments
}

func buildChainByAverage(dotsSize int, segments []segment.Segment, start int, averageWeight float64) (result []*segment.Segment) {
	var dots []int
	// Get a starting segment
	seg := &segments[start]
	segLinks := seg.Links
	//result = append(result, seg)
	dots = append(dots, seg.Start)

	var freshSegs []*segment.Segment
	freshSegs = append(freshSegs, seg)

	for _, value := range segLinks {
		if value.Weight < averageWeight {
			res, segs := checkSegment(value, append(freshSegs, value), averageWeight, dotsSize - 2)
			if res {
				return segs
			}
		}
	}

	return nil
}

func main() {
	// Size of matrix
	n := 5

	initial := segment.GenerateInitialMatrix(n, 10.0)

	segments := segment.GenerateMatrix(initial)

	segment.ShowInitialMatrix(initial)
	segment.ShowMatrix(segments)

	// Basic algo
	// builtSegments := buildChain(n, segments, 0)
	// fmt.Println(builtSegments)

	// Get average weight
	averageWeight := getAverageWeight(initial)
	fmt.Println("Average element weight is:", averageWeight)

	// Get overall average
	fmt.Println("Overall matrix average is:", averageWeight * float64(n - 1))

	var builtSegments []*segment.Segment
	for i := 0; i < len(segments); i++ {
		builtSegments = buildChainByAverage(n, segments, i, averageWeight)
		
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
}
