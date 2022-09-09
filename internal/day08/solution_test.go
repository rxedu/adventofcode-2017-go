package day08

import (
	"sync"
	"testing"
)

type exampleOne struct {
	i []Instruction
	o int
}

type exampleTwo struct {
	i []Instruction
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []exampleOne{
		{i: []Instruction{}, o: 0},
		{i: []Instruction{{reg: "a", change: 10, condition: Condition{reg: "a", op: "=", val: 0}}}, o: 10},
		{i: []Instruction{
			{reg: "a", change: 10, condition: Condition{reg: "a", op: "=", val: 0}},
		}, o: 10},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e exampleOne) {

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
	examples := []exampleTwo{
		{i: []Instruction{}, o: 0},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e exampleTwo) {

			defer wg.Done()
			got := solvePartTwo(e.i)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}
