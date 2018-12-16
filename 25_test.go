package algo

import "testing"

func TestMergeTwoSortedList(t *testing.T) {
	list1 := new(LinkList)
	list1.Val = 1
	list1.Next = new(LinkList)
	list1.Next.Val = 2
	list1.Next.Next = new(LinkList)
	list1.Next.Next.Val = 3
	list1.Next.Next.Next = new(LinkList)
	list1.Next.Next.Next.Val = 4

	list2 := new(LinkList)
	list2.Val = 1
	list2.Next = new(LinkList)
	list2.Next.Val = 3

	list3 := new(LinkList)
	list3.Val = 2
	list3.Next = new(LinkList)
	list3.Next.Val = 4

	got := mergeTwoSortedList(list2, list3)
	for got != nil {
		if got.Val != list1.Val {
			t.Errorf("Failed, expected%d, got %d", list1.Val, got.Val)
		}
		got = got.Next
		list1 = list1.Next
	}
}
