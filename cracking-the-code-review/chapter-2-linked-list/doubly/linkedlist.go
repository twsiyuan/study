package list

// List represents a doubly linked list
type List struct {
	root  *Element
	count int
}

// First returns the first element in the list
func (l *List) First() *Element {
	if l == nil {
		return nil
	}
	return l.root.Next()
}

// Last returns the last element in the list
func (l *List) Last() *Element {
	if l == nil {
		return nil
	}
	return l.root.Previous()
}

// Count returns the number of elements in the list
func (l *List) Count() int {
	if l == nil {
		return 0
	}
	return l.count
}

func (l *List) lazyInit() {
	if l.root == nil {
		elem := &Element{list: l}
		elem.next = elem
		elem.prev = elem
		l.root = elem
	}
}

// PushLast adds the given value into the last in the list and returns the element if successed
func (l *List) PushLast(v interface{}) *Element {
	if l == nil {
		return nil
	}
	l.lazyInit()

	elem := &Element{Value: v, list: l}
	at := l.root.prev
	elem.prev = at
	elem.next = at.next
	elem.prev.next = elem
	elem.next.prev = elem

	l.count++
	return elem
}

// PushLastSlice adds the given values into the last in the list and returns the last element
func (l *List) PushLastSlice(v ...interface{}) *Element {
	var elem *Element
	for _, vv := range v {
		elem = l.PushLast(vv)
	}
	return elem
}

// PushLastList adds the given list into the last in the list and returns the last element
func (l *List) PushLastList(v *List) *Element {
	var elem *Element
	for at := v.First(); at != nil; at = at.Next() {
		elem = l.PushLast(at.Value)
	}
	return elem
}

// Remove removes the given element and returns true if it's an element of list
func (l *List) Remove(element *Element) bool {
	if l == nil || element == nil || element.list != l {
		return false
	}

	prev := element.prev
	next := element.next

	prev.next = next
	next.prev = prev

	element.prev = nil // avoid memory leak
	element.next = nil // avoid memory leak
	element.list = nil // avoid memory leak
	l.count--

	return true
}

// Element represents an element in a list
type Element struct {
	Value interface{}

	next *Element
	prev *Element
	list *List
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if e == nil || e.list == nil || e.next == e.list.root {
		return nil
	}
	return e.next
}

// Previous returns the previous list element or nil.
func (e *Element) Previous() *Element {
	if e == nil || e.list == nil || e.prev == e.list.root {
		return nil
	}
	return e.prev
}

// ToArray converts the given linked list into a array
func ToArray(l *List) []interface{} {
	a := make([]interface{}, 0, l.Count())
	for at := l.First(); at != nil; at = at.Next() {
		a = append(a, at.Value)
	}
	return a
}
