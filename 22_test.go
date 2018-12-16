package algo

import "testing"

func TestFindKThNodeToTail(t *testing.T) {
	list1 := new(LinkList)
	list1.Val = 1
	list1.Next = new(LinkList)
	list1.Next.Val = 2
	list1.Next.Next = new(LinkList)
	list1.Next.Next.Val = 3
	list1.Next.Next.Next = new(LinkList)
	list1.Next.Next.Next.Val = 4
	got := findKThNodeToTail(list1, 2)
	if got != list1.Next.Next {
		t.Errorf("Failed, expected %d, got %d", list1.Next.Next, got)
	}
}
