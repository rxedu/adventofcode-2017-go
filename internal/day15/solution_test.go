package day15

import (
	"sync"
	"testing"
)

type example struct {
	i Generators
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: Generators{
			Generator{65, aFactor, aMultiple},
			Generator{8921, bFactor, bMultiple},
		}, o: 588},
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
		{i: Generators{
			Generator{65, aFactor, aMultiple},
			Generator{8921, bFactor, bMultiple},
		}, o: 309},
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
