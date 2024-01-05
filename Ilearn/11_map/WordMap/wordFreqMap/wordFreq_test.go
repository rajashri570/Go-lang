package wordFreqMap

import (
	"reflect"
	"testing"
)

func TestFreq(t *testing.T) {
	result := GetFreq("I am rajashri. rajashri is good person")
	expected := map[string]int{"I": 1, "am": 1, "rajashri": 2, "is": 1, "good": 1, "person": 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test fails. Got %v, expected %v", result, expected)
	}
}
