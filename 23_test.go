package algo

import "testing"

func TestEntryNodeOfLoop(t *testing.T) {
	list1 := new(LinkList)
	list1.Val = 1
	list1.Next = new(LinkList)
	list1.Next.Val = 2
	list1.Next.Next = new(LinkList)
	list1.Next.Next.Val = 3
	list1.Next.Next.Next = new(LinkList)
	list1.Next.Next.Next.Val = 4
	list1.Next.Next.Next.Next = new(LinkList)
	list1.Next.Next.Next.Next.Val = 5
	list1.Next.Next.Next.Next.Next = new(LinkList)
	list1.Next.Next.Next.Next.Next.Val = 6
	list1.Next.Next.Next.Next.Next.Next = list1.Next.Next

	got := entryNodeOfLoop(list1)
	if got != list1.Next.Next {
		t.Fatalf("Failed, expected %d, got %d", list1.Next.Next, got)
	}

}
