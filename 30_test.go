package algo

import "testing"

func TestMinStackPush(t *testing.T) {
	stack := new(minStack)
	stack.push(3)
	if stack.min() != 3 {
		t.Errorf("Failed, expected %d, got %d", 3, stack.min())
	}
	stack.push(4)
	if stack.min() != 3 {
		t.Errorf("Failed, expected %d, got %d", 3, stack.min())
	}
	stack.push(2)
	if stack.min() != 2 {
		t.Errorf("Failed, expected %d, got %d", 2, stack.min())
	}
	stack.pop()
	if stack.min() != 3 {
		t.Errorf("Failed, expected %d, got %d", 3, stack.min())
	}
}
