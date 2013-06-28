package decisiontree

import "testing"

func TestDedupe(t *testing.T) {
	test_unique := make([]pairFloat64, 4)
	test_dupe := make([]pairFloat64, 4)
	for i := 0; i < len(test_dupe); i++ {
		test_unique[i] = pairFloat64{float64(i), 1.}
		test_dupe[i] = pairFloat64{float64(i % 2), 1.}
	}
	if len(Dedupe(test_unique)) != 4 {
		t.Errorf("False positive!")
	}
	r := Dedupe(test_dupe)
	if len(r) != 2 {
		t.Errorf("False negative! Expected length 2, found length %d", len(r))
	}
}
