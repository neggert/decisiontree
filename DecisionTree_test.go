package decisiontree

import (
	"testing"
)

func TestFindOptimalCut(t *testing.T) {
	xvals := []float64{1, 2, 3, 4}
	yvals := []float64{0, 0, 0.9, 1}
	cut, std := findOptimalCut(xvals, yvals)
	if cut != 2.5 {
		t.Errorf("Expected cut at 2.5, got %f", cut)
	}
	if !FloatEqual(std, 0.05, 1e-4) {
		t.Errorf("Expected std for best cut at 0.05, got %f", std)
	}
}

func TestFindOptimalCutNoVarianceData(t *testing.T) {
	xvals := []float64{0, 0, 0, 0}
	yvals := []float64{0, 0, 0.9, 1}
	cut, std := findOptimalCut(xvals, yvals)
	if cut != 0. {
		t.Errorf("Expected cut at 0., got %f", cut)
	}
	if std != 1e30 {
		t.Errorf("Expected std for best cut at 1e30, got %f", std)
	}
}

func TestFindOptimalCutNoVarianceTarget(t *testing.T) {
	xvals := []float64{1, 2, 3, 4}
	yvals := []float64{1, 1, 1, 1}
	cut, std := findOptimalCut(xvals, yvals)
	if cut != 0. {
		t.Errorf("Expected cut at 0., got %f", cut)
	}
	if std != 0. {
		t.Errorf("Expected std for best cut at 0., got %f", std)
	}
}

func TestCreateDecisionNode(t *testing.T) {
	xvals := [][]float64{{1, 2},
		{2, 4},
		{1, 6},
		{2, 8}}
	yvals := []float64{0, 1, 0, 1}
	tree := createDecisionNode(xvals, yvals, 2)
	if tree.value != 0.5 {
		t.Errorf("Expected node value of 0.5, found %f", tree.value)
	}
	if tree.cutoff != 1.5 {
		t.Errorf("Expected cutoff of 1.5, found %f", tree.cutoff)
	}
	if tree.column != 0 {
		t.Errorf("Expected column 0, found %d", tree.column)
	}
}

func TestWalk(t *testing.T) {
	xvals := [][]float64{{1, 2},
		{2, 4},
		{1, 6},
		{2, 8}}
	yvals := []float64{0, 1, 0, 1}
	output := make([]float64, 4)
	tree := CreateDecisionTree(xvals, yvals, 2)
	for i, data := range xvals {
		output[i] = tree.Walk(data)
	}
	if !SliceEqual(output, yvals) {
		t.Errorf("Expected prediction %v, got %v", yvals, output)
	}
}
