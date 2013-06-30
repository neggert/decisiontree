package decisiontree

import "testing"

func TestDedupe(t *testing.T) {
	test_unique := []pairFloat64{{1., 1.}, {2., 1.}, {3., 1.}, {4., 1.}}
	test_dupe := []pairFloat64{{1., 1.}, {1., 1.}, {2., 1.}, {2., 1.}}
	if len(Dedupe(test_unique)) != 4 {
		t.Errorf("False positive!")
	}
	r := Dedupe(test_dupe)
	if len(r) != 2 {
		t.Errorf("False negative! Expected length 2, found length %d", len(r))
	}
}

func TestFloatEqual(t *testing.T) {
	if !FloatEqual(0.6, 0.65, 0.1) {
		t.Errorf("False positive! FloatEqual(0.6, 0.65, 0.1) returned false, should be true")
	}
	if FloatEqual(0.6, 0.65, 0.01) {
		t.Errorf("False negative! FloatEqual(0.6, 0.65, 0.01) returned true, should be false")
	}
}

func TestSliceEqual(t *testing.T) {
	a := []float64{1, 2, 3, 4}
	b := []float64{1, 2, 3, 4}
	if !SliceEqual(a, b) {
		t.Errorf("False negative! %v and %v should be equal", a, b)
	}
	c := []float64{1, 2, 3, 5}
	if SliceEqual(a, c) {
		t.Errorf("False positive! %v and %v should not be equal", a, c)
	}
	d := []float64{1, 2, 3, 4, 5}
	if SliceEqual(a, d) {
		t.Errorf("False positive! %v and %v should not be equal", a, d)
	}
}
