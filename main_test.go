package main

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	head := fromSlice([]int{4, 5, 1, 9})
	if !DeleteNode(head.Next) {
		t.Fatal("deleting a middle node should succeed")
	}
	if got := toSlice(head); !reflect.DeepEqual(got, []int{4, 1, 9}) {
		t.Errorf("after delete: %v, want [4 1 9]", got)
	}

	// Deleting the head works the same way.
	head2 := fromSlice([]int{1, 2, 3})
	DeleteNode(head2)
	if got := toSlice(head2); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("after head delete: %v, want [2 3]", got)
	}

	// Tail and nil are refused.
	head3 := fromSlice([]int{1, 2})
	if DeleteNode(head3.Next) {
		t.Error("tail delete should return false")
	}
	if DeleteNode(nil) {
		t.Error("nil delete should return false")
	}
}
