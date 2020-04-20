package sort

import (
	"testing"

	"github.com/Johnny4Fun/InterviewPreparation/helper"
)

type args struct {
	arr      []int
	expected []int
}
type testCase struct {
	name string
	args args
}

func generateArraySamples() []testCase {
	tests := []testCase{
		{
			name: "normal test",
			args: args{
				arr: []int{
					8, 6, 9, 5, 4, 7, 3, 1, 2,
				},
				expected: []int{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
		},
		{
			name: "ascending array",
			args: args{
				arr: []int{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
				expected: []int{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
		},
		{
			name: "descending array",
			args: args{
				arr: []int{
					9, 8, 7, 6, 5, 4, 3, 2, 1,
				},
				expected: []int{
					1, 2, 3, 4, 5, 6, 7, 8, 9,
				},
			},
		},
		{
			name: "empty",
			args: args{
				arr:      []int{},
				expected: []int{},
			},
		},
	}
	return tests
}
func TestHeapSort(t *testing.T) {

	tests := generateArraySamples()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
			helper.AssertSliceEqual(t, tt.args.arr, tt.args.expected)
		})
	}
}

func TestBubbleSort(t *testing.T) {

	tests := generateArraySamples()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.arr)
			helper.AssertSliceEqual(t, tt.args.arr, tt.args.expected)
		})
	}
}

func TestSelectionSort(t *testing.T) {

	tests := generateArraySamples()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.arr)
			helper.AssertSliceEqual(t, tt.args.arr, tt.args.expected)
		})
	}
}

func TestQuickSort(t *testing.T) {

	tests := generateArraySamples()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			helper.AssertSliceEqual(t, tt.args.arr, tt.args.expected)
		})
	}
}

func TestCountingSort(t *testing.T) {

	tests := generateArraySamples()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountingSort(tt.args.arr)
			helper.AssertSliceEqual(t, tt.args.arr, tt.args.expected)
		})
	}
}
