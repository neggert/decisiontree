// Utilities used in the decision tree
package decisiontree

// pair datatype so that we can sort by one field
// and the other goes along for the ride
type pairFloat64 struct {
	sort_val float64
	other_val float64
}


// slice of pairs
type pairFloat64Collection []pairFloat64
// Add methods so we can sort the pairCollection

func (l pairFloat64Collection) Len() int {
	return len(l)
}

// Less by sort key
func (l pairFloat64Collection) Less(i,j  int) bool {
	return l[i].sort_val < l[j].sort_val
}

// swap both values
func (l pairFloat64Collection) Swap(i, j int) {
	l[i].sort_val, l[j].sort_val = l[j].sort_val, l[i].sort_val
	l[i].other_val, l[j].other_val = l[j].other_val, l[i].other_val
}

// Dedupe removes items with duplicate sort values
func Dedupe(l pairFloat64Collection) []pairFloat64 {
	unique := make(map[float64]bool)
	out := make([]pairFloat64, 0, len(l))
	for i:=0; i<len(l); i++ {
		_, ok := unique[l[i].sort_val]
		if !ok {
			out = append(out, l[i])
			unique[l[i].sort_val] = true
		}
	}
	return out
}