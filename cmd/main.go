package main

import (
	"fmt"

	"github.com/baubletech/traveling-gopher/segment"
)

func buildChain(segments []segment.Segment, start int) (result []*segment.Segment) {
	// Get a starting segment
	seg := &segments[start]
	segLinks := seg.Links
	result = append(result, seg)

	for {
		// Check links are enabled, break otherwise
		var availableLinks []*segment.Segment
		for _, value := range segLinks {
			if !value.Disabled {
				availableLinks = append(availableLinks, value)
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
	}

	return result
}

func main() {
	segments := segment.GenerateMatrix(5, 5.0)
	fmt.Println(segments)

	builtSegments := buildChain(segments, 0)
	fmt.Println(builtSegments)

	var result []int
	for _, value := range builtSegments {
		result = append(result, value.ID)
	}
	fmt.Println(result)
}
