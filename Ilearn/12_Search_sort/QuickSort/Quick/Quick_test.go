package Quick

import (
	"reflect"
	"testing"
)

func Test_Quick(t *testing.T) {
	input := []float64{4.494965654325735, 2.4884998774532328, 9.928963008404843, 4.391562995201735, 8.356103323753441}
	expected := []float64{2.4884998774532328, 4.391562995201735, 4.494965654325735, 8.356103323753441, 9.928963008404843}
	result := Quick_sort(input)
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

}
