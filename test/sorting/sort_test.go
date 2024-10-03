package sorting

import (
	"testing"

	"github.com/santakdalai90/GolangDSA/src/sorting"
	"github.com/stretchr/testify/assert"
)

type sortAlgo func([]int)

type sorter struct {
	name string
	algo sortAlgo
}

var sorters = []sorter{
	{"Bubble Sort", sorting.BubbleSort},
	{"Selection Sort", sorting.SelectionSort},
	{"Insertion Sort", sorting.InsertionSort},
	{"Merge Sort", sorting.MergeSort},
}

type testCase struct {
	inputArr  []int
	sortedArr []int
}

func TestSort(t *testing.T) {
	var testData = []testCase{
		{
			[]int{65, 55, 45, 35, 25, 15, 10},
			[]int{10, 15, 25, 35, 45, 55, 65},
		},
		{
			[]int{4, 2, 1, 7, 3},
			[]int{1, 2, 3, 4, 7},
		},
		{
			[]int{95, 21, 47, 63, 31, 53, 19, 101},
			[]int{19, 21, 31, 47, 53, 63, 95, 101},
		},
		{
			[]int{65, 55},
			[]int{55, 65},
		},
		{
			[]int{65},
			[]int{65},
		},
	}

	for _, s := range sorters {
		t.Logf("running %s", s.name)
		for _, test := range testData {
			s.algo(test.inputArr)
			assert.Equalf(t, test.sortedArr, test.inputArr, "failed sorting, algo name:%s", s.name)
		}
	}
}
