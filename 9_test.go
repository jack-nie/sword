package algo

import "testing"

func TestAppendTail(t *testing.T) {
	t.Run("Append 1", func(t *testing.T) {
		queue := &StackQueue{
			stack1: new(Stack),
			stack2: new(Stack),
		}
		queue.AppendTail(1)
		if got := queue.stack1.Pop(); got != 1 {
			t.Errorf("failed, expected %d, got %d", 1, got)
		}
	})

	t.Run("Append 1, 2", func(t *testing.T) {
		queue := &StackQueue{
			stack1: new(Stack),
			stack2: new(Stack),
		}
		queue.AppendTail(1)
		queue.AppendTail(2)

		if got := queue.stack1.Pop(); got != 2 {
			t.Errorf("failed, expected %d, got %d", 2, got)
		}
	})
}

func TestDeleteHead(t *testing.T) {
	t.Run("Delete from empty queue", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did  not panic")
			}
		}()

		queue := &StackQueue{
			stack1: new(Stack),
			stack2: new(Stack),
		}
		queue.DeleteHead()
	})

	t.Run("Delete from queue", func(t *testing.T) {
		queue := &StackQueue{
			stack1: new(Stack),
			stack2: new(Stack),
		}
		queue.AppendTail(1)
		queue.DeleteHead()
		if got := queue.stack1.Size(); got != 0 {
			t.Errorf("failed, expected %d, got %d", 0, got)
		}
	})

	t.Run("Delete until empty", func(t *testing.T) {
		queue := &StackQueue{
			stack1: new(Stack),
			stack2: new(Stack),
		}
		queue.AppendTail(1)
		queue.AppendTail(2)
		queue.AppendTail(3)
		queue.AppendTail(4)
		queue.DeleteHead()
		queue.DeleteHead()
		queue.DeleteHead()
		queue.DeleteHead()
		if got := queue.stack1.Size(); got != 0 {
			t.Errorf("failed, expected %d, got %d", 0, got)
		}
	})
}
