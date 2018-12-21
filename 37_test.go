package algo

import (
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestSerialize(t *testing.T) {
	root := constructBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	r, w := io.Pipe()
	go serialize(root, w)
	content, err := ioutil.ReadAll(r)
	if err != nil {
		t.Errorf("Failed, %s", err)
	}
	if got := string(content); got != "136$8$$5$$2$47$$$" {
		t.Errorf("Failed, expected %s, got %s", "136$8$$5$$2$47$$$", got)
	}
}

func TestDeserialize(t *testing.T) {
	root := constructBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	r, w := io.Pipe()
	go serialize(root, w)
	node := new(TreeNode)
	deserialize(r, &node)
	var stack []int
	got := inOrder(node, &stack)
	if !reflect.DeepEqual(got, []int{6, 8, 3, 5, 1, 2, 7, 4}) {
		t.Errorf("Failed, expected %d, got %d", []int{6, 8, 3, 5, 1, 2, 7, 4}, got)
	}
}
