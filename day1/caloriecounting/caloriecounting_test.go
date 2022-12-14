package caloriecounting_test

import (
	"reflect"
	"testing"

	"github.com/AldoFusterTurpin/AdventOfCode-2022/day1/caloriecounting"
)

func TestGetTotalCaloriesOfTheTopNElvesCarryingTheMostCalories(t *testing.T) {
	expected := 45000
	elvesCalories := [][]int{
		[]int{
			1000,
			2000,
			3000,
		},
		[]int{4000},
		[]int{5000, 6000},
		[]int{7000, 8000, 9000},
		[]int{10000},
	}

	got := caloriecounting.GetTotalCaloriesOfTheTopNElvesCarryingTheMostCalories(elvesCalories, 3)
	if expected != got {
		t.Fatalf("expected %v, but got %v", expected, got)
	}
}

func TestGetTotalCaloriesOfTheElfCarryingMostCalories(t *testing.T) {
	expected := 24000
	elvesCalories := [][]int{
		[]int{
			1000,
			2000,
			3000,
		},
		[]int{4000},
		[]int{5000, 6000},
		[]int{7000, 8000, 9000},
		[]int{10000},
	}

	got := caloriecounting.GetTotalCaloriesOfTheElfCarryingMostCalories(elvesCalories)
	if expected != got {
		t.Fatalf("expected %v, but got %v", expected, got)
	}
}

func TestGetInventoryFromInputString(t *testing.T) {
	type testData struct {
		inputData string
		expected  [][]int
	}

	tests := map[string]testData{
		"simple_input": {
			inputData: `1
2
3
4
5
6

7
8

9
10`,
			expected: [][]int{
				[]int{1, 2, 3, 4, 5, 6},
				[]int{7, 8},
				[]int{9, 10},
			},
		},
		"single_char": {
			inputData: `1`,
			expected: [][]int{
				[]int{1},
			},
		},
		"empty_input": {
			inputData: ``,
			expected:  [][]int{},
		},
		"input_is_a_space": {
			inputData: ` `,
			expected:  [][]int{},
		},
		"input_is_a_blank_line": {
			inputData: `
`,
			expected: [][]int{},
		},
		"input_from_example": {
			inputData: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
			expected: [][]int{
				[]int{1000, 2000, 3000},
				[]int{4000},
				[]int{5000, 6000},
				[]int{7000, 8000, 9000},
				[]int{10000},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := caloriecounting.GetAllElvesCalories(tc.inputData)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
