package main

import (
	"github.com/parkervcp/fsquota"

	"github.com/spf13/cobra"
)

var cmdProject = &cobra.Command{
	Use:   "project",
	Short: "Project quota management",
}

func init() {
	cmdRoot.AddCommand(cmdProject)
}

func lookupProject(projectIdOrName string) (prj *fsquota.Project, err error) {
	if isNumeric(projectIdOrName) {
		prj.ID = projectIdOrName
		return
	}

	return fsquota.LookupProject(projectIdOrName)
}
