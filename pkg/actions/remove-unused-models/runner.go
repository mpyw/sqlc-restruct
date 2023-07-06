package removeunusedmodels

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path"

	"github.com/mpyw/sqlc-restruct/pkg/actions/internal/astutil"
)

type runner struct {
	input ActionInput
	fset  *token.FileSet
}

func (r *runner) Run() error {
	mf, err := parser.ParseFile(r.fset, path.Join(r.input.ModelsDir, r.input.ModelsFileName), nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse models file: %w", err)
	}

	iff, err := parser.ParseFile(r.fset, path.Join(r.input.IfaceDir, r.input.QuerierFileName), nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse interface file: %w", err)
	}

	// Generate used model type indexes
	modelsToBeKept := make(map[string]bool)
	if r.input.IfaceDir == r.input.ModelsDir {
		ast.Walk(
			astutil.NewExportedExprIdentUpdater(func(ident *ast.Ident) ast.Expr {
				modelsToBeKept[ident.Name] = true
				return nil
			}),
			iff,
		)
	} else {
		ast.Walk(
			astutil.NewExportedSelectorExprUpdater(func(expr *ast.SelectorExpr) ast.Expr {
				if x, ok := expr.X.(*ast.Ident); ok {
					if x.Name == r.input.ModelsPkgName {
						modelsToBeKept[expr.Sel.Name] = true
					}
				}
				return nil
			}),
			iff,
		)
	}

	// Remove unused decls
	var newDecls []ast.Decl
	for _, decl := range mf.Decls {
		newDecls = append(newDecls, decl)
		switch decl := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.ValueSpec:
					if ident, ok := spec.Type.(*ast.Ident); ok {
						if _, shouldKeep := modelsToBeKept[ident.Name]; !shouldKeep {
							newDecls = newDecls[:len(newDecls)-1]
						}
					}
				case *ast.TypeSpec:
					if ident, ok := spec.Type.(*ast.Ident); ok {
						if _, shouldKeep := modelsToBeKept[ident.Name]; !shouldKeep {
							newDecls = newDecls[:len(newDecls)-1]
						}
					}
				}
			}
		}
	}

	byt, err := r.intoBytes(mf)
	if err != nil {
		return fmt.Errorf("failed to generate new content: %w", err)
	}

	_ = os.Remove(path.Join(r.input.ModelsDir, r.input.ModelsFileName))
	if err := os.WriteFile(path.Join(r.input.ModelsDir, r.input.ModelsFileName), byt, 0644); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

func (r *runner) intoBytes(node any) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := format.Node(buf, r.fset, node); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
