package main

import (
	"errors"

	"github.com/parkervcp/fsquota"
	"github.com/spf13/cobra"
)

func lookupProjectNameByID(id string) string {
	if p, err := fsquota.LookupProjectID(id); err == nil {
		return p.Name
	}
	return id
}

var cmdProjectReport = &cobra.Command{
	Use:   "report path",
	Short: "Retrieves quota report for a given path",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) != 1 {
			err = errors.New("exactly one argument required")
			return
		}

		var report *fsquota.Report
		if report, err = fsquota.GetProjectReport(args[0]); err != nil {
			return
		}

		lookupFn := lookupProjectNameByID

		if wantNumeric, _ := cmd.Flags().GetBool("numeric"); wantNumeric {
			lookupFn = noopLookup
		}

		printReport(cmd, report, "project", lookupFn)
		return
	},
}

func init() {
	cmdProjectReport.Flags().BoolP("numeric", "n", false, "Print numeric 		project IDs")
	cmdProject.AddCommand(cmdProjectReport)
}
