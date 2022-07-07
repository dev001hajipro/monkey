package evaluator

import (
	"github.com/dev001hajipro/monkey/ast"
	"github.com/dev001hajipro/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
