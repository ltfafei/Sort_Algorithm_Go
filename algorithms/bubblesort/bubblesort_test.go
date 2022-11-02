//-*- coding: UTF-8 -*-
// Author: afei00123

package bubblesort

import "testing"

func TestBubblesort1(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	Bubblesort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("Bubblesort() failed. Got", values, "Excepted 1 2 3 4 5")
	}
}

func TestBubblesort2(t *testing.T) {
	values := []int{5, 5, 3, 4, 5}
	Bubblesort(values)
	if values[0] != 5 || values[1] != 5 || values[2] != 5 || values[3] != 3 || values[4] != 4 {
		t.Error("Bubblesort() failed. Got", values, "Excepted 5 5 5 3 4")
	}
}