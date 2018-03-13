package handySlices

import "testing"

type item struct {
	k string
	i int
}

func (i item) KeyStringValue() string {
	return i.k
}

var (
	a = []item{
		{"k0", 0},
		{"k1", 1},
		{"k2", 2},
	}
	b = []item{
		{"k1", 1},
		{"k2", 1},
		{"k3", 1},
	}
)

func expect(t *testing.T, result []item, expectedResult []item) {
	if len(result) != len(expectedResult) {
		t.Errorf("Unexpected result: %v (%v) != %v (%v)", result, len(result), expectedResult, len(expectedResult))
		return
	}
	for idx, _ := range result {
		if result[idx].k != expectedResult[idx].k || result[idx].i != expectedResult[idx].i {
			t.Errorf("Unexpected result: %v != %v", result[idx], expectedResult[idx])
			return
		}
	}

	return
}

func TestGetSubtraction(t *testing.T) {
	r := GetSubtraction(a, b).([]item)
	expect(t, r, []item{{"k0", 0}})
}

func TestGetIntersection(t *testing.T) {
	r := GetIntersection(a, b).([]item)
	expect(t, r, []item{{"k1", 1}, {"k2", 2}})
}

func TestGetDiffedIntersection(t *testing.T) {
	r := GetDiffedIntersection(a, b).([]item)
	expect(t, r, []item{{"k2", 2}})
}
func TestMapToSlice(t *testing.T) {
	r := MapToSlice(map[string]item{"k0":{"k0", 0}, "k1":{"k1", 1}}).([]item)
	expect(t, r, []item{{"k0", 0}, {"k1", 1}})
}
