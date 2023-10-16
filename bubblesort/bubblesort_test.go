package math

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T)	{

	sli := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	BubbleSort(sli)

	if !reflect.DeepEqual(expected, sli) {
		t.Errorf("Expected %v, received %v", expected, sli)
	}
}

func TestSort2(t *testing.T)	{

	sli := []int{6, 7, 8, 9, 0, 2, 1, 3, 5, 4}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	BubbleSort(sli)

	if !reflect.DeepEqual(expected, sli) {
		t.Errorf("Expected %v, received %v", expected, sli)
	}
}

func TestEmpty(t *testing.T)	{

	sli := []int{}
	expected := []int{}

	BubbleSort(sli)

	if !reflect.DeepEqual(expected, sli) {
		t.Errorf("Expected %v, received %v", expected, sli)
	}
}