package evaluator

import (
	"github.com/dev001hajipro/monkey/ast"
	"github.com/dev001hajipro/monkey/object"
)

// Node object of the argument is wrapped in a Quote object
func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
