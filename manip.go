package astmanip

import (
	"go/ast"
)

func NextSibling(parent, pivot ast.Node) (result ast.Node) {
	capture := false
	ast.Inspect(parent, func(node ast.Node) bool {
		if result != nil {
			return false
		} else if node == pivot {
			capture = true
			return false
		} else if capture {
			result = node
			return false
		} else {
			return true
		}
	})

	return
}

func InsertStmtAfter(list []ast.Stmt, stmt, ref ast.Stmt) []ast.Stmt {
	for i, s := range list {
		if s == ref {
			return append(list[0:i+1], append([]ast.Stmt{stmt}, list[i+1:]...)...)
		}
	}

	return append(list, stmt)
}
