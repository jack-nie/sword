package algo

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	list1 := generateListFromSlice([]int{1, 2, 3, 4})
	list2 := generateListFromSlice([]int{3, 4})
	deleteNode(&list1, list2)
	wanted := convertListToSlice(list1)
	got := convertListToSlice(generateListFromSlice([]int{1, 2, 4}))
	if !reflect.DeepEqual(wanted, got) {
		t.Errorf("Failed, expected %d, got %d", wanted, got)
	}
}
