/*
Decision tree package. So far only handles float data.
Categorial data should be coded as dummy variables
*/

package decisiontree

import (
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

// Given a column of data, find the cutoff that gives the
// smallest sum of RMS
func findOptimalCut(column, target []float64) float64 {
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
	// extract it back into two different arrays which are now sorted and deduplicated
	for i, pair := range paired {
		column[i] = pair.sort_val
		target[i] = pair.other_val
	}

	// now loop through cuts to find the best one
	bestRMS, bestCut := 1.e30, 0.
	var rms float64
	for i := 1; i < len(column); i++ {
		rms = stats.RMS(target[:i]) + stats.RMS(target[i:])
		if rms < bestRMS {
			bestRMS = rms
			bestCut = column[i]
		}
	}
	return bestCut
}

func createDecisionNode(data [][]float64, target []float64) {
	return
}

// CreateDecisionTree builds a decision tree from training data
// returns a pointer to the created tree
func CreateDecisionTree(data [][]float64, target []float64) *DecisionTree {
	return &DecisionTree{1., 2, 3., nil, nil}
}
