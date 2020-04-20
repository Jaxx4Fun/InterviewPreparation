package helper

import (
	"reflect"
	"testing"
)

func AssertSliceEqual(t *testing.T, got, want []int) {
	t.Helper()
	if gotLen, wantLen := len(got), len(want); gotLen != wantLen {
		t.Errorf("got length(%d) not equals want length(%d)", gotLen, wantLen)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
