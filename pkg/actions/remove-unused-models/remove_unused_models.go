package removeunusedmodels

import (
	"context"
	"go/token"
)

type ActionInput struct {
	// IfaceDir The directory path where the separated interface will be located.
	IfaceDir string
	// ModelsPkgName The package name where the separated models will be located.
	ModelsPkgName string
	// ModelsDir The directory path where the separated models will be located.
	ModelsDir string
	// ModelsFileName The file name for the sqlc-generated models file.
	ModelsFileName string
	// QuerierFileName The file name for the sqlc-generated `Querier` file.
	QuerierFileName string
}

func Action(_ context.Context, input ActionInput) error {
	r := &runner{
		input: input,
		fset:  token.NewFileSet(),
	}
	return r.Run()
}
