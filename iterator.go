// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

type Iterator func(i Item) bool

// Ascend will call iterator once for each element greater or equal than pivot
// in ascending order. It will stop whenever the iterator returns false.
func (t *Rbtree) Ascend(pivot Item, iterator Iterator) {
	t.ascend(t.root, pivot, iterator)
}

func (t *Rbtree) ascend(x *Node, pivot Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(x.Item, pivot) {
		if !t.ascend(x.Left, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.ascend(x.Right, pivot, iterator)
}

// Descend will call iterator once for each element less or equal than pivot
// in descending order. It will stop whenever the iterator returns false.
func (t *Rbtree) Descend(pivot Item, iterator Iterator) {
	t.descend(t.root, pivot, iterator)
}

func (t *Rbtree) descend(x *Node, pivot Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(pivot, x.Item) {
		if !t.descend(x.Right, pivot, iterator) {
			return false
		}
		if !iterator(x.Item) {
			return false
		}
	}

	return t.descend(x.Left, pivot, iterator)
}

// AscendRange will call iterator once for elements greater or equal than @ge
// and less than @lt, which means the range would be [ge, lt).
// It will stop whenever the iterator returns false.
func (t *Rbtree) AscendRange(ge, lt Item, iterator Iterator) {
	t.ascendRange(t.root, ge, lt, iterator)
}

func (t *Rbtree) ascendRange(x *Node, inf, sup Item, iterator Iterator) bool {
	if x == t.NIL {
		return true
	}

	if !less(x.Item, sup) {
		return t.ascendRange(x.Left, inf, sup, iterator)
	}
	if less(x.Item, inf) {
		return t.ascendRange(x.Right, inf, sup, iterator)
	}

	if !t.ascendRange(x.Left, inf, sup, iterator) {
		return false
	}
	if !iterator(x.Item) {
		return false
	}
	return t.ascendRange(x.Right, inf, sup, iterator)
}

func (t *Rbtree) AscNext(pivot Item) Item {
	return t.ascnext(t.root, pivot).Item
}

func (t *Rbtree) ascnext(x *Node, pivot Item) *Node {
	if x == t.NIL {
		return t.NIL
	}

	if !less(pivot, x.Item) {
		return t.ascnext(x.Right, pivot)
	} else {
		if x.Left != t.NIL {
			if !less(pivot, t.max(x.Left).Item) {
				return x
			} else {
				return t.ascnext(x.Left, pivot)
			}
		}
		return x
	}
}

func (t *Rbtree) DescNext(pivot Item) Item {
	return t.descnext(t.root, pivot).Item
}

func (t *Rbtree) descnext(x *Node, pivot Item) *Node {
	if x == t.NIL {
		return t.NIL
	}

	if !less(x.Item, pivot) {
		return t.descnext(x.Left, pivot)
	} else {
		if x.Right != t.NIL {
			if !less(t.min(x.Right).Item, pivot) {
				return x
			} else {
				return t.descnext(x.Right, pivot)
			}
		}
		return x
	}
}
