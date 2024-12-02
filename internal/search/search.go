package search

import (
	"errors"
	"math"
	"slices"
)

type SearchResult struct {
	Index  int `json:"index"`
	Number int `json:"number"`
}

func FindIndexOrClosest(numbers []int, target int) (SearchResult, error) {
	if len(numbers) == 0 {
		return SearchResult{}, errors.New("empty list")
	}

	index, found := slices.BinarySearch(numbers, target)

	if found {
		return SearchResult{Index: index, Number: target}, nil
	}

	if index == 0 {
		return closestValueIfExists(target, numbers[index], index)
	}

	if index == len(numbers) {
		return closestValueIfExists(target, numbers[len(numbers)-1], index-1)
	}

	before := numbers[index-1]
	after := numbers[index]

	if target-before <= after-target {
		return closestValueIfExists(target, before, index-1)
	}

	return closestValueIfExists(target, after, index)
}

func closestValueIfExists(target int, value int, index int) (SearchResult, error) {
	fTarget := float64(target)
	minThreshold := int(math.Ceil(fTarget - 0.1*fTarget))
	maxThreshold := int(math.Floor(fTarget + 0.1*fTarget))

	if minThreshold <= value && value <= maxThreshold {
		return SearchResult{Index: index, Number: value}, nil
	}

	return SearchResult{}, errors.New("Number not found")
}
