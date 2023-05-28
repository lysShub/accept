package tree

import (
	c "golang.org/x/exp/constraints"
)

// https://www.bilibili.com/video/BV1xE411h7dd

const (
	leftHeavy  = -1
	balanced   = 0
	rightHeavy = +1
)

type ordered = c.Ordered

// AVL tree
type Tree[T ordered] struct {
	// b[0] is entry, 初始即存在, 存类型的0值
	b []block[T]
}

type block[T ordered] struct {
	l, r   uint32
	p      uint32
	factor int8

	key T
	val int
}

// func (b *block[T]) end() bool {
// 	return b.l == b.r
// }

/*
	2
  1   3

*/

func (t *Tree[T]) Set(key T, val int) {

	var idx uint32 = 0
	var n = t.get()
	t.b[n].val = val

	for {
		if t.b[idx].key < key {
			if t.b[idx].l == 0 {
				t.b[idx].l = n
				t.b[n].p = idx

				t.lReshape(n)
				return
			} else {
				idx = t.b[idx].l
			}
		} else if key < t.b[idx].key {
			if t.b[idx].r == 0 {
				t.b[idx].r = n
				t.b[n].p = idx

				t.rReshape(n)
				return
			} else {
				idx = t.b[idx].r
			}
		} else {
			t.b[idx].val = val
			return
		}
	}
}

func (t *Tree[T]) get() (idx uint32) {
	return
}

// 插入在左边
func (t *Tree[T]) lReshape(idx uint32) {

	// child index
	var cidx uint32 = idx
	print(cidx)
	for idx = t.b[idx].p; idx != 0; idx = t.b[idx].p {

	}

	if t.b[idx].factor > rightHeavy {

	} else if t.b[idx].factor < leftHeavy {

	}
}

func (t *Tree[T]) rReshape(idx uint32) {

	if t.b[idx].factor > rightHeavy {

	} else if t.b[idx].factor < leftHeavy {

	}
}

// 右旋
func (t *Tree[T]) rotateRight(idx uint32) {

	/*
		5 9 3 4 2 1
	*/

	p := t.b[idx].p
	n := t.b[idx].l

	if idx == t.b[p].r {
		t.b[p].r = n
	} else {
		t.b[p].l = n
	}

	t.b[n].r = idx
	t.b[idx].p = n

	t.b[idx].l = t.b[n].r
	if t.b[n].r != 0 {
		t.b[t.b[n].r].p = idx
		// t.b[n].factor += rightHeavy
	}

	if t.b[n].factor == balanced {
		t.b[idx].factor = leftHeavy
		t.b[n].factor = rightHeavy
	} else {
		t.b[idx].factor = balanced
		t.b[n].factor = balanced
	}
}

func (t *Tree[T]) rotateLeft(idx uint32) {
	n := t.b[idx].r
	p := t.b[idx].p

	if idx == t.b[p].l {
		t.b[p].l = n
	} else {
		t.b[p].r = n
	}

	t.b[n].l = idx
	t.b[idx].p = n

	t.b[idx].r = t.b[n].l
	if t.b[n].l != 0 {
		t.b[t.b[n].l].p = idx
	}

	if t.b[n].factor == balanced {
		t.b[idx].factor = rightHeavy
		t.b[n].factor = leftHeavy
	} else {
		t.b[idx].factor = balanced
		t.b[n].factor = balanced
	}
}
