package linkedlist

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	l := New(1, New(2, nil))

	if l.head != 1 || l.tail.head != 2 {
		t.Errorf("New does not work")
	}
}

func TestApply(t *testing.T) {
	l := Apply(1, 2, 3, 4)
	if l.head != 1 || l.tail.head != 2 || l.tail.tail.head != 3 {
		t.Errorf("Apply does not work")
	}
}

func TestSetHeadOO(t *testing.T) {
	l := Apply("h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d", "!")
	l.setHead("H")
	h := l.getHead()

	if h != "H" {
		t.Errorf("Wanted H, but got %s", h)
	}
}

func TestSetHeadFP(t *testing.T) {
	l := Apply("h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d", "!")
	k := "K"
	l = SetHead(l, k)
	h := l.getHead()

	if h != k {
		t.Errorf("Wanted %s, but got %s", k, h)
	}
}

func TestGetTail(t *testing.T) {
	l := Apply("h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d", "!")

	tail := "e"
	tl := l.getTail()

	if tl.head != tail {
		t.Errorf("wanted %s, but got %s", tail, tl.head)
	}
}

func TestDropOO(t *testing.T) {
	t.Skip("skipping oo style drop...")
	// l := Apply("h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d", "!")

	// l.drop(2)
	// if l.head != "l" {
	// 	t.Errorf("wanted l, got %s", l.head)
	// }

	// l1 := Apply(1, 2, 3, 4, 5)
	// l1.drop(4)
	// if l1.head != 5 {
	// 	t.Errorf("wanted 5, got %d", l1.head)
	// }

	l1 := Apply(1, 2, 3, 4, 5)
	fmt.Println(&l1)
	// l1 = l1.tail

	l1.drop(1)

	if l1.head != 2 {
		t.Errorf("Did not work")
	}
}

func TestDropFP(t *testing.T) {
	l := Apply(1, 2, 3, 4)
	l = Drop(l, 2)

	if l.head != 3 {
		t.Errorf("wanted 3, go %d", l.head)
	}
}

func TestDropWhileFP(t *testing.T) {
	l := Apply(2, 4, 6, 8, 5, 6, 7)

	l = DropWhile(l, func(t int) bool { return t%2 == 0 })

	if l.head != 5 {
		t.Errorf("wanted 5, got %d", l.head)
	}

}

func TestAppend(t *testing.T) {
	l1 := New(1, nil)
	l2 := Apply(2, 3, 4, 54)

	l3 := Append(l1, l2)

	if l3.tail.head != 2 {
		t.Errorf("wanted 2, got %d", l3.tail.head)
	}
}

func TestLast(t *testing.T) {
	l := Apply(1, 2, 3, 4, 5)
	last := Last(l)
	if last != 5 {
		t.Errorf("wanted 5, got %d", last)
	}
}

func TestDropLast(t *testing.T) {
	l := Apply(1, 2, 4, 5)
	l = DropLast(l)
	last := Last(l)

	if last != 4 {
		t.Errorf("wanted 4, got %d", last)
	}
}
