// descriptivestats is a package to do simple 1-variable statistics (currently
// only count, mean, and standard deviation) without keeping a list of variables
// around.
package descriptivestats

import (
	"math"
)

// Stats contains the information needed to compute statistics. The zero value
// of Stats is a valid Stats with no data.
type Stats struct {
	Count float64
	sum   float64
	sumsq float64
}

// Add records a data point.
func (s *Stats) Add(data float64) {
	s.Count++
	s.sum += data
	s.sumsq += data * data
}

// Mean returns the mean of the recorded data points.
func (s *Stats) Mean() float64 {
	return s.sum / s.Count
}

// Variance returns the variance of the recorded data points.
func (s *Stats) Variance() (variance float64) {
	if s.Count < 2 {
		return 0
	}
	mean := s.Mean()
	variance = s.sumsq - s.Count*mean*mean
	if variance < 0 {
		variance = 0
	} else {
		variance /= (s.Count - 1)
	}
	return
}

// StdDev returns the standard deviation of the recorded data points.
func (s *Stats) StdDev() float64 {
	return math.Sqrt(s.Variance())
}
