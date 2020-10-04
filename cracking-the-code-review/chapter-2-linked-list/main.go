package main

type Node struct {
	Value interface{}
	Next  *Node
}

func ToArray(n *Node) []interface{} {
	r := make([]interface{}, 0)
	for ; n != nil; n = n.Next {
		r = append(r, n.Value)
	}
	return r
}

func main() {

}
