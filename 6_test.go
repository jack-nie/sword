package algo

import (
	"reflect"
	"testing"
)

func TestPrintLinkListReversinglyIteratively(t *testing.T) {
	list := generateListFromSlice([]int{1, 2, 3, 4, 5})
	for _, tt := range []struct {
		list   *LinkList
		wanted []int
	}{
		{list, []int{1, 2, 3, 4, 5}},
	} {
		if got := printLinkListReversinglyIteratively(tt.list); !reflect.DeepEqual(tt.wanted, got) {
			t.Errorf("failed, expected %d, got %d", tt.wanted, got)
		}
	}
}
