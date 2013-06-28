package decisiontree

import "testing"

func TestFindOptimalCut(t *testing.T) {
	xvals := make([]float64, 4)
	yvals := make([]float64, 4)
	for i:=0; i < len(xvals); i++ {
		xvals[i] = float64(i)
	}
	yvals[len(yvals)-1] = 1.
	cut := findOptimalCut(xvals, yvals)
	if cut != 3.0 {
		t.Errorf("Expected cut at 3.0, got %f", cut)
	}
}