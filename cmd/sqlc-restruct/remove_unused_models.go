package main

import (
	"errors"

	"github.com/mpyw/sqlc-restruct/pkg/actions/remove-unused-models"
	"github.com/urfave/cli/v2"
)

var RemoveUnusedModelsCommand = &cli.Command{
	Name:    "remove-unused-models",
	Usage:   "Removes unused models that aren't referenced by any interface methods. This helps in reducing clutter and keeps the codebase lean.",
	Aliases: []string{},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "iface-dir",
			Usage:    "The directory path where the interface will be located.",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "models-dir",
			Usage: "The directory path where the model will be located. (default: --iface-dir value)",
		},
		&cli.StringFlag{
			Name:  "models-pkg-name",
			Usage: "The package name where the model will be located. Required if the model directory path is different from the interface directory path.",
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
		iDir := c.String("iface-dir")
		mDir := c.String("models-dir")
		if mDir == "" {
			mDir = iDir
		}

		mPkgName := c.String("models-pkg-name")
		if mDir != iDir && mPkgName == "" {
			return errors.New("models-pkg-name is required when models-dir is different from iface-dir")
		}

		return removeunusedmodels.Action(c.Context, removeunusedmodels.ActionInput{
			IfaceDir:        iDir,
			ModelsPkgName:   mPkgName,
			ModelsDir:       mDir,
			ModelsFileName:  c.String("models-file-name"),
			QuerierFileName: c.String("querier-file-name"),
		})
	},
}

func init() {
	App.Commands = append(App.Commands, RemoveUnusedModelsCommand)
}
