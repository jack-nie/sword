package algo

import "testing"

func TestReversedList(t *testing.T) {
	list1 := new(LinkList)
	list1.Val = 1
	list1.Next = new(LinkList)
	list1.Next.Val = 2
	list1.Next.Next = new(LinkList)
	list1.Next.Next.Val = 3
	list1.Next.Next.Next = new(LinkList)
	list1.Next.Next.Next.Val = 4

	list2 := new(LinkList)
	list2.Val = 4
	list2.Next = new(LinkList)
	list2.Next.Val = 3
	list2.Next.Next = new(LinkList)
	list2.Next.Next.Val = 2
	list2.Next.Next.Next = new(LinkList)
	list2.Next.Next.Next.Val = 1

	got := reverseList(list1)
	for list2 != nil && got != nil {
		if got.Val != list2.Val {
			t.Errorf("Failed, expected %d, got %d", list2.Val, got.Val)
		}
		list2 = list2.Next
		got = got.Next
	}

}
