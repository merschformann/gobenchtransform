package example

type RangeChecker interface {
	// InRange returns true if the given integer is in the range of any of the
	// ranges in the RangeChecker.
	InRange(int) bool
}

// Range defines a range of integers.
type Range struct {
	// Min is the minimum value of the range (inclusive).
	Min int
	// Max is the maximum value of the range (inclusive).
	Max int
}

// --> Slice implementation

type rangeCheckerSlice struct {
	ranges []Range
}

func (r *rangeCheckerSlice) InRange(i int) bool {
	for _, r := range r.ranges {
		if i >= r.Min && i <= r.Max {
			return true
		}
	}
	return false
}

func NewRangeCheckerSlice(ranges []Range) RangeChecker {
	return &rangeCheckerSlice{ranges}
}

// --> Map implementation

type rangeCheckerMap struct {
	rangeMap map[int]bool
}

func (r *rangeCheckerMap) InRange(i int) bool {
	return r.rangeMap[i]
}

func NewRangeCheckerMap(ranges []Range) RangeChecker {
	rangeMap := make(map[int]bool)
	for _, r := range ranges {
		for i := r.Min; i <= r.Max; i++ {
			rangeMap[i] = true
		}
	}
	return &rangeCheckerMap{rangeMap}
}
