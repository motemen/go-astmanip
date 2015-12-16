package astmanip

import (
	"testing"

	"go/ast"
	"go/parser"
	"go/token"
)

func TestNextSibling(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "testdata/nextsibling/nextsibling.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	if NextSibling(f, f.Decls[0]) != f.Decls[1] {
		t.Fail()
	}

	if NextSibling(f, f.Decls[1]) != f.Decls[2] {
		t.Fail()
	}

	if NextSibling(f, f.Decls[2]) != nil {
		t.Fail()
	}

	fun := f.Decls[2].(*ast.FuncDecl)
	if NextSibling(f, fun.Body.List[0]) != fun.Body.List[1] {
		t.Fail()
	}

	if NextSibling(f, fun.Body.List[1]) != fun.Body.List[2] {
		t.Fail()
	}

	if NextSibling(f, fun.Body.List[2]) != nil {
		t.Fail()
	}

	ifStmt := fun.Body.List[2].(*ast.IfStmt)
	if NextSibling(f, ifStmt.Body.List[0]) != ifStmt.Body.List[1] {
		t.Fail()
	}

	if NextSibling(f, ifStmt.Body.List[1]) != nil {
		t.Fail()
	}
}
