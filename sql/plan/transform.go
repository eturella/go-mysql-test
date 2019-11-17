package plan

import (
	"github.com/eturella/go-mysql-test/sql"
)

// TransformUp applies a transformation function to the given tree from the
// bottom up.
func TransformUp(node sql.Node, f sql.TransformNodeFunc) (sql.Node, error) {
	if o, ok := node.(sql.OpaqueNode); ok && o.Opaque() {
		return f(node)
	}

	children := node.Children()
	if len(children) == 0 {
		return f(node)
	}

	newChildren := make([]sql.Node, len(children))
	for i, c := range children {
		c, err := TransformUp(c, f)
		if err != nil {
			return nil, err
		}
		newChildren[i] = c
	}

	node, err := node.WithChildren(newChildren...)
	if err != nil {
		return nil, err
	}

	return f(node)
}
