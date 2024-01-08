package Bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort_asc(t *testing.T) {
	input := []rune("rajashri")

	expected := []rune("aahhijrrs")

	result := Bubble_sort(input, func(a, b rune) bool {
		return a > b
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestBubbleSort_dec(t *testing.T) {
	input := []rune("raj")

	expected := []rune("rja")

	result := Bubble_sort(input, func(a, b rune) bool {
		return a < b
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
