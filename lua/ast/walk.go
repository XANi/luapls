package ast

import (
	"reflect"

	"github.com/raiguard/luapls/lua/token"
)

type Visitor func(node Node) bool

// Walk performs a depth-first traversal of the AST, calling the visitor for each node.
// If the visitor returns false, this node's children are not traversed.
func Walk(node Node, visitor Visitor) {
	if node == nil || reflect.ValueOf(node).IsNil() {
		return
	}
	if !visitor(node) {
		return
	}

	WalkList(node.Leaves(), visitor)
}

// WalkList traverses a slice of Nodes.
func WalkList[N Node](nodes []N, v Visitor) {
	for _, n := range nodes {
		Walk(n, v)
	}
}

type NodePath struct {
	Node    Node
	Parents []Node
}

// GetNode returns the innermost node at the given position, and its parent nodes.
func GetNode(base Node, pos token.Pos) NodePath {
	var node Node
	parents := []Node{}
	Walk(base, func(n Node) bool {
		if n.Pos() <= pos && pos < n.End() {
			if node != nil {
				parents = append(parents, node)
			}
			node = n
			return true
		}
		return false
	})
	return NodePath{node, parents}
}
