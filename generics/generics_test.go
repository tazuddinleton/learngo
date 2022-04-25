package generics

import "testing"

func TestNonGenericSum(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	intSum := sumInts(ints)
	floatSum := sumFloats(floats)

	if intSum != 46 {
		t.Errorf("wanted 46, got %d", intSum)
	}

	if floatSum != 62.97 {
		t.Errorf("wanted 62.97, got %f", floatSum)
	}

}

func TestGenericSum(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	intSum := sumNumbers(ints)

	floatSum := sumNumbers(floats)

	if intSum != 46 {
		t.Errorf("wanted 46, got %d", intSum)
	}

	if floatSum != 62.97 {
		t.Errorf("wanted 62.97, got %f", floatSum)
	}
}
