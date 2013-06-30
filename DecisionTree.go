/*
Decision tree package. So far only handles float data.
Categorial data should be coded as dummy variables
*/

package decisiontree

import (
	"fmt"
	"github.com/neggert/stats"
	"sort"
)

// A node in a decision tree
type DecisionTree struct {
	value  float64
	column int
	cutoff float64
	Low    *DecisionTree
	High   *DecisionTree
}

// Walk sends a data point through the decision tree and
// returns the prediction
func (t *DecisionTree) Walk(item []float64) float64 {
	if t.Low == nil && t.High == nil {
		return t.value
	}
	if item[t.column] < t.cutoff {
		return t.Low.Walk(item)
	} else {
		return t.High.Walk(item)
	}
}

func (t *DecisionTree) String() string {
	return getString(t, 0)
}

func getString(t *DecisionTree, depth int) string {
	out := ""
	for i := 0; i < depth; i++ {
		out += fmt.Sprint(" ")
	}
	if t.Low == nil && t.High == nil {
		out += fmt.Sprintf("Value: %f\n", t.value)
	} else {
		out += fmt.Sprintf("Column: %d    Cutoff: %f\n", t.column, t.cutoff)
	}
	if t.Low != nil {
		out += getString(t.Low, depth+1)
	}
	if t.High != nil {
		out += getString(t.High, depth+1)
	}
	return out
}

// Given a column of data, find the cutoff that gives the
// smallest sum of RMS
func findOptimalCut(column, target []float64) (bestCut, bestRMS float64) {
	// join the column and the targets together
	paired := make(pairFloat64Collection, len(column))
	for i := 0; i < len(column); i++ {
		paired[i].sort_val = column[i]
		paired[i].other_val = target[i]
	}
	// sort the indices by the column values
	sort.Sort(paired)
	// eliminate duplicates
	paired = Dedupe(paired)
	if len(paired) <= 1 {
		bestCut = 0.
		bestRMS = 1e30
		return bestCut, bestRMS
	} else if stats.RMS(target) == 0. {
		bestCut = 0.
		bestRMS = 0.
		return bestCut, bestRMS
	}
	// extract it back into two different arrays which are now sorted and deduplicated
	for i, pair := range paired {
		column[i] = pair.sort_val
		target[i] = pair.other_val
	}

	// now loop through cuts to find the best one
	bestRMS, bestCut = 1.e30, 0.
	var rms float64
	for i := 1; i < len(column); i++ {
		rms = stats.RMS(target[:i]) + stats.RMS(target[i:])
		if rms < bestRMS {
			bestRMS = rms
			bestCut = (column[i] + column[i-1]) / 2
		}
	}
	return bestCut, bestRMS
}

func createDecisionNode(data [][]float64, target []float64, minSamples int) *DecisionTree {
	// ending conditions
	if (len(target) < minSamples) || (stats.RMS(target) == 0) {
		return &DecisionTree{stats.Mean(target), 0, 0., nil, nil}
	}
	// Find the best variable to split on
	bestRMS := 1e30
	var bestCut float64
	var bestCol int
	cut := 0.
	rms := 0.
	for i := 0; i < len(data[0]); i++ {
		// make sure there's some variance in the column
		cut, rms = findOptimalCut(data[:][i], target)
		if rms < bestRMS {
			bestRMS = rms
			bestCol = i
			bestCut = cut
		}
	}
	lowData := make([][]float64, 0, len(data)/2)
	lowTarget := make([]float64, 0, len(data)/2)
	highData := make([][]float64, 0, len(data)/2)
	highTarget := make([]float64, 0, len(data)/2)

	for i, row := range data {
		if row[bestCol] <= bestCut {
			lowData = append(lowData, row)
			lowTarget = append(lowTarget, target[i])
		} else {
			highData = append(highData, row)
			highTarget = append(highTarget, target[i])
		}
	}

	node := DecisionTree{stats.Mean(target), bestCol, bestCut,
		createDecisionNode(lowData, lowTarget, minSamples), createDecisionNode(highData, highTarget, minSamples)}
	return &node
}

// CreateDecisionTree builds a decision tree from training data
// returns a pointer to the created tree
func CreateDecisionTree(data [][]float64, target []float64, minSamples int) *DecisionTree {
	return createDecisionNode(data, target, minSamples)
}
