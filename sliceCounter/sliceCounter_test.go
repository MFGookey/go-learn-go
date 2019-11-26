package main

import (
	"testing"
)

func Test_Empty_Slice(t *testing.T) {
	haystack := []int{}
	needle := 0
	expectedNeedles := 0
	result := CountSlices(haystack, needle)
	if result != expectedNeedles {
		t.Errorf("Empty slice count was incorrect.  Expected \"%d\" but got \"%d\"", expectedNeedles, result)
	}
}

func Test_No_Match_Slice(t *testing.T){
	haystack := []int{1, 2, 3}
	needle := 0
	expectedNeedles := 0
	result := CountSlices(haystack, needle)
	if result != expectedNeedles {
		t.Errorf("No-match slice count was incorrect.  Expected \"%d\" but got \"%d\"", expectedNeedles, result)
	}
}

func Test_Single_Match_Slice(t *testing.T){
	haystack := []int{1, 2, 3, 1}
	needle := 2
	expectedNeedles := 1
	result := CountSlices(haystack, needle)
	if result != expectedNeedles {
		t.Errorf("Single match slice count was incorrect.  Expected \"%d\" but got \"%d\"", expectedNeedles, result)
	}
}


func Test_Multiple_Match_Slice(t *testing.T){
	haystack := []int{1, 2, 3, 1}
	needle := 1
	expectedNeedles := 2
	result := CountSlices(haystack, needle)
	if result != expectedNeedles {
		t.Errorf("Multiple match slice count was incorrect.  Expected \"%d\" but got \"%d\"", expectedNeedles, result)
	}
}