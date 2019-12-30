package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"log"

	"github.com/baubletech/traveling-gopher/segment"
)

// BestSegment is a segment with minimal weight overall value
type BestSegment struct {
	Weight   float64
	Segments []*segment.Segment
}

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
			if !value.Disabled {
				if len(dots) < (dotsSize - 1) {
					if !intInSlice(value.End, dots) {
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

	return sum / float64(len(initial)*len(initial)-len(initial))
}

func isStartInSegments(segments []*segment.Segment, start int) bool {
	for _, v := range segments {
		if v.Start == start {
			return true
		}
	}
	return false
}

// Check segment comparing to average
func checkSegment(segment *segment.Segment, segments []*segment.Segment, averageWeight float64, n int) (bool, []*segment.Segment) {
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
			res, segs := checkSegment(value, append(freshSegs, value), averageWeight, dotsSize-2)
			if res {
				return segs
			}
		}
	}

	return nil
}

func getChainWeight(segments []*segment.Segment) (sum float64) {
	for _, seg := range segments {
		sum += seg.Weight
	}

	return sum
}

// Check segment using full traversal
func checkSegmentFull(segment *segment.Segment, segments []*segment.Segment, best *BestSegment, n int) []*segment.Segment {
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

func buildChainFullTraversal(dotsSize int, segments []segment.Segment, best *BestSegment, start int) *BestSegment {
	var dots []int
	// Get a starting segment
	seg := &segments[start]
	segLinks := seg.Links
	//result = append(result, seg)
	dots = append(dots, seg.Start)

	var freshSegs []*segment.Segment
	freshSegs = append(freshSegs, seg)

	for _, value := range segLinks {
		checkSegmentFull(value, append(freshSegs, value), best, dotsSize-2)
	}

	return best
}

func outputFullReport(n int) {
	initial := segment.GenerateInitialMatrix(n, 10.0)
	segments := segment.GenerateMatrix(initial)

	segment.ShowInitialMatrix(initial)
	segment.ShowMatrix(segments)

	// Get average weight
	averageWeight := getAverageWeight(initial)
	fmt.Println("Average element weight is:", averageWeight)

	// Get overall average
	fmt.Println("Overall matrix average is:", averageWeight*float64(n-1))

	// Build with average
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

func testAlgo(n int, times int) {
	diffPercent := 15.0

	var successful int
	var averageChangeSum float64

	for i := 0; i < times; i++ {
		initial := segment.GenerateInitialMatrix(n, 10.0)
		segments := segment.GenerateMatrix(initial)

		// Get average weight
		averageWeight := getAverageWeight(initial)

		// Build with average
		var builtSegments []*segment.Segment
		for i := 0; i < len(segments); i++ {
			builtSegments = buildChainByAverage(n, segments, i, averageWeight)

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

		averageChangeSum += change
	}

	fmt.Println("Average change is:", averageChangeSum / float64(times))
	fmt.Println("Times:", times, "Successful (15%):", successful)
	fmt.Println("Success percentage:", (float64(successful) / float64(times)) * 100.0)
}

func percentageChange(old, new float64) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}

func getIntFromUser(prompt string) int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(result)
}

func main() {
	// Size of matrix
	n := getIntFromUser("Enter matrix size: ")

	if a := getUserInput("Would you like to run testing? (y/n): "); a == "y" {
		times := getIntFromUser("How many iterations?: ")
		testAlgo(n, times)
	} else {
		outputFullReport(n)
	}
}
