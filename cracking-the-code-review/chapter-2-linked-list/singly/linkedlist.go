package list

// List represents a singly linked list
type List struct {
	first *Element
	last  *Element
	count int
}

// First returns the first element in the list
func (l List) First() *Element {
	return l.first
}

// Last returns the last element in the list
func (l List) Last() *Element {
	return l.last
}

// Count returns the number of elements in the list
func (l List) Count() int {
	return l.count
}

// PushLast adds the given value into the last in the list and returns the element if successed
func (l *List) PushLast(v interface{}) *Element {
	if l == nil {
		return nil
	}

	e := &Element{Value: v}
	if l.count <= 0 {
		l.first = e
		l.last = e
		l.count++
		return e
	}

	l.last.next = e
	l.last = e
	l.count++
	return e
}

// PushLastSlice adds the given values into the last in the list and returns the last element
func (l *List) PushLastSlice(v ...interface{}) *Element {
	var elem *Element
	for _, vv := range v {
		elem = l.PushLast(vv)
	}
	return elem
}

// Remove removes the given element and returns true if it's an element of list
func (l *List) Remove(element *Element) bool {
	if l == nil || element == nil || l.count <= 0 {
		return false
	}

	defer func() {
		if l.count <= 1 {
			l.last = l.first
		}
	}()

	if l.first == element {
		l.first = element.next
		l.count--
		element.next = nil // avoid memory leak
		return true
	}

	at := l.first
	for i := 1; i < l.count; i++ {
		if at.next == element {
			at.next = element.next
			if l.last == element {
				l.last = at
			}
			l.count--
			element.next = nil // avoid memory leak
			return true
		}

		at = at.next
	}
	return false
}

// Element represents an element in a list
type Element struct {
	Value interface{}

	next *Element
}

// Next returns the next list element or nil.
func (e Element) Next() *Element {
	return e.next
}

// ToArray converts the given linked list into a array
func ToArray(l *List) []interface{} {
	a := make([]interface{}, 0, l.Count())
	for at := l.First(); at != nil; at = at.Next() {
		a = append(a, at.Value)
	}
	return a
}
