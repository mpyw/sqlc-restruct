package main

import (
	"github.com/mpyw/sqlc-restruct/pkg/actions/separate-interface"
	"github.com/urfave/cli/v2"
)

var SeparateInterfaceCommand = &cli.Command{
	Name:    "separate-interface",
	Usage:   "Separates models and the `Querier` interface from the `Queries` struct. This is typically done to adhere to the Dependency Inversion Principle (DIP), allowing for more flexible and testable code.",
	Aliases: []string{},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "iface-pkg-name",
			Usage:    "The package name where the separated models and Querier will be located.",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "iface-pkg-url",
			Usage:    "The package URL where the separated models and Querier will be located (e.g. \"github.com/<user>/<repo>/path/to/pkg\").",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "iface-dir",
			Usage:    "The directory path where the separated models and Querier will be located.",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "impl-dir",
			Usage: "The original directory where the sqlc-generated code is located.",
			Value: ".",
		},
		&cli.StringFlag{
			Name:  "impl-sql-suffix",
			Usage: "The suffix for sqlc-generated files from SQL files.",
			Value: ".sql.go",
		},
		&cli.StringFlag{
			Name:  "models-file-name",
			Usage: "The file name for the sqlc-generated models file.",
			Value: "models.go",
		},
		&cli.StringFlag{
			Name:  "querier-file-name",
			Usage: "The file name for the sqlc-generated Querier file.",
			Value: "querier.go",
		},
	},
	Action: func(c *cli.Context) error {
		return separateinterface.Action(c.Context, separateinterface.ActionInput{
			IfacePkgName:    c.String("iface-pkg-name"),
			IfacePkgURL:     c.String("iface-pkg-url"),
			IfaceDir:        c.String("iface-dir"),
			ImplDir:         c.String("impl-dir"),
			ImplSQLSuffix:   c.String("impl-sql-suffix"),
			ModelsFileName:  c.String("models-file-name"),
			QuerierFileName: c.String("querier-file-name"),
		})
	},
}

func init() {
	App.Commands = append(App.Commands, SeparateInterfaceCommand)
}
