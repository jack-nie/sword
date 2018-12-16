package algo

import "testing"

func TestFindFirstCommonNode(t *testing.T) {
	list1 := new(LinkList)
	list1.Val = 1
	list1.Next = new(LinkList)
	list1.Next.Val = 2
	list1.Next.Next = new(LinkList)
	list1.Next.Next.Val = 3
	list1.Next.Next.Next = new(LinkList)
	list1.Next.Next.Next.Val = 4

	list2 := new(LinkList)
	list2.Val = 5
	list2.Next = new(LinkList)
	list2.Next.Val = 6
	list2.Next.Next = new(LinkList)
	list2.Next.Next.Val = 3
	list2.Next.Next.Next = new(LinkList)
	list2.Next.Next.Next.Val = 4

	target := list2.Next.Next
	for _, v := range []struct {
		list1  *LinkList
		list2  *LinkList
		wanted *LinkList
	}{
		{list1, list2, target},
	} {
		got := findFirstCommonNode(list1, list2)
		if got != v.wanted {
			t.Fatalf("Failed, expected %d, got %d", v.wanted.Val, got.Val)
		}
	}

}
