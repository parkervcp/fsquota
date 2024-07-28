package main

import (
	"errors"

	"github.com/parkervcp/fsquota"
	"github.com/spf13/cobra"
)

var cmdProjectGet = &cobra.Command{
	Use:   "get path project",
	Short: "Retrieves quota information for a given project",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) != 2 {
			err = errors.New("exactly two arguments required")
			return
		}

		var p *fsquota.Project
		if p, err = lookupProject(args[1]); err != nil {
			return
		}

		var info *fsquota.Info
		if info, err = fsquota.GetProjectInfo(args[0], p); err != nil {
			return
		}

		printQuotaInfo(cmd, info)

		return
	},
}

func init() {
	cmdGroup.AddCommand(cmdProjectGet)
}
