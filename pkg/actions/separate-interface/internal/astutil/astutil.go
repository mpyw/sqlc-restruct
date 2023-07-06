package astutil

import (
	"go/ast"
	"go/token"
	"unicode"
)

func HasCapitalName(ident *ast.Ident) bool {
	return ident != nil && len(ident.Name) > 0 && unicode.IsUpper(rune(ident.Name[0]))
}

func ExportedIndividualTypeOrValueDecls(decls ...ast.Decl) []ast.Decl {
	return individualTypeOrValueDecls(true, decls...)
}

func UnexportedIndividualTypeOrValueDecls(decls ...ast.Decl) []ast.Decl {
	return individualTypeOrValueDecls(false, decls...)
}

func FuncDecls(decls ...ast.Decl) []ast.Decl {
	var exported []ast.Decl
	for _, decl := range decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			exported = append(exported, decl)
		}
	}
	return exported
}

func ExtractImportDecls(decls ...ast.Decl) []ast.Decl {
	var specs []ast.Spec
	for _, decl := range decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					specs = append(specs, spec)
				}
			}
		}
	}
	if len(specs) < 1 {
		return nil
	}
	return []ast.Decl{&ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs,
	}}
}

func SymbolNameFromTypeOrValueDecls(decls ...ast.Decl) []string {
	var symbols []string
	for _, decl := range decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.ValueSpec:
					for _, name := range spec.Names {
						symbols = append(symbols, name.Name)
					}
				case *ast.TypeSpec:
					symbols = append(symbols, spec.Name.Name)
				}
			}
		}
	}
	return symbols
}

func individualSpecs(exp bool, specs ...ast.Spec) []ast.Spec {
	var exported []ast.Spec
	for _, spec := range specs {
		switch spec := spec.(type) {
		case *ast.TypeSpec:
			// type XXX ...
			if HasCapitalName(spec.Name) == exp {
				exported = append(exported, spec)
			}
		case *ast.ValueSpec:
			// const X, Y, Z = 1, 2, 3
			// var X, Y, Z = 1, 2, 3
			// var X, Y, Z int
			for i, name := range spec.Names {
				if HasCapitalName(name) == exp {
					var values []ast.Expr
					if len(spec.Values) > i {
						values = []ast.Expr{spec.Values[i]}
					}
					exported = append(exported, &ast.ValueSpec{
						Doc:     spec.Doc,
						Names:   []*ast.Ident{spec.Names[i]},
						Type:    spec.Type,
						Values:  values,
						Comment: spec.Comment,
					})
				}
			}
		}
	}
	return exported
}

func individualTypeOrValueDecls(exp bool, decls ...ast.Decl) []ast.Decl {
	var exported []ast.Decl
	for _, decl := range decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				if specs := individualSpecs(exp, spec); len(specs) > 0 {
					exported = append(exported, &ast.GenDecl{
						Tok:   decl.Tok,
						Specs: specs,
					})
				}
			}
		}
	}
	return exported
}

func NewExportedExprIdentUpdater(updater func(*ast.Ident) ast.Expr) *ExportedExprIdentUpdater {
	return &ExportedExprIdentUpdater{updater: updater}
}

type ExportedExprIdentUpdater struct {
	updater func(*ast.Ident) ast.Expr
}

func (r *ExportedExprIdentUpdater) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncDecl:
		r.walkFuncDecl(n)
		return nil
	case *ast.Field:
		if expr := r.resolveExpr(n.Type); expr != nil {
			n.Type = expr
		}
		return nil
	case *ast.ValueSpec:
		if expr := r.resolveExpr(n.Type); expr != nil {
			n.Type = expr
		}
		for i, value := range n.Values {
			if expr := r.resolveExpr(value); expr != nil {
				n.Values[i] = expr
			}
		}
		return nil
	case *ast.TypeSpec:
		if expr := r.resolveExpr(n.Type); expr != nil {
			n.Type = expr
		}
	case *ast.AssignStmt:
		for i, rh := range n.Rhs {
			if expr := r.resolveExpr(rh); expr != nil {
				n.Rhs[i] = rh
			}
		}
	case *ast.InterfaceType:
		ast.Inspect(n, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.Field:
				if _, isInterfaceMethod := n.Type.(*ast.FuncType); !isInterfaceMethod {
					if expr := r.resolveExpr(n.Type); expr != nil {
						n.Type = expr
					}
					return false
				}
			}
			return true
		})
	}
	return r
}

func (r *ExportedExprIdentUpdater) walkFuncDecl(n *ast.FuncDecl) {
	// Explicitly exclude method receiver
	ast.Walk(r, n.Type.Params)
	ast.Walk(r, n.Type.Results)
	ast.Walk(r, n.Body)
}

func (r *ExportedExprIdentUpdater) resolveExpr(n ast.Expr) ast.Expr {
	switch n := n.(type) {
	case *ast.StarExpr:
		if expr := r.resolveExpr(n.X); expr != nil {
			n.X = expr
		}
	case *ast.ArrayType:
		if expr := r.resolveExpr(n.Elt); expr != nil {
			n.Elt = expr
		}
	case *ast.Ellipsis:
		if expr := r.resolveExpr(n.Elt); expr != nil {
			n.Elt = expr
		}
	case *ast.SliceExpr:
		if expr := r.resolveExpr(n.X); expr != nil {
			n.X = expr
		}
	case *ast.CompositeLit:
		if expr := r.resolveExpr(n.Type); expr != nil {
			n.Type = expr
		}
	case *ast.Ident:
		if HasCapitalName(n) {
			return r.updater(n)
		}
	}
	return n
}
