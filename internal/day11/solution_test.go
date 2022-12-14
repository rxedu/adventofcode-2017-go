package day11

import (
	"sync"
	"testing"
)

type example struct {
	i []Step
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: []Step{}, o: 0},
		{i: []Step{
			{1, -1, 0},
			{1, -1, 0},
			{1, -1, 0},
		}, o: 3},
		{i: []Step{
			{1, -1, 0},
			{1, -1, 0},
			{-1, 1, 0},
			{-1, 1, 0},
		}, o: 0},
		{i: []Step{
			{1, -1, 0},
			{1, -1, 0},
			{0, 1, -1},
			{0, 1, -1},
		}, o: 2},
		{i: []Step{
			{1, 0, -1},
			{-1, 1, 0},
			{1, 0, -1},
			{-1, 1, 0},
			{-1, 1, 0},
		}, o: 3},
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
		{i: []Step{}, o: 0},
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
