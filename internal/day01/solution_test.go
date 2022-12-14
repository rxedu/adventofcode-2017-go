package day01

import (
	"sync"
	"testing"
)

type example struct {
	i []int
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: []int{}, o: 0},
		{i: []int{1, 1, 2, 2}, o: 3},
		{i: []int{1, 1, 1, 1}, o: 4},
		{i: []int{1, 2, 3, 4}, o: 0},
		{i: []int{9, 1, 2, 1, 2, 1, 2, 9}, o: 9},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartOne(e.i)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}

func TestPartTwoExamples(t *testing.T) {
	examples := []example{
		{i: []int{}, o: 0},
		{i: []int{1}, o: 0},
		{i: []int{1, 2, 1, 2}, o: 6},
		{i: []int{1, 2, 2, 1}, o: 0},
		{i: []int{1, 2, 3, 4, 2, 5}, o: 4},
		{i: []int{1, 2, 3, 1, 2, 3}, o: 12},
		{i: []int{1, 2, 1, 3, 1, 4, 1, 5}, o: 4},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartTwo(e.i)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}
