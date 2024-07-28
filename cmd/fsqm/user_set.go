package main

import (
	"errors"
	"os/user"

	"github.com/parkervcp/fsquota"
	"github.com/spf13/cobra"
)

var cmdUserSet = &cobra.Command{
	Use:   "set path user",
	Short: "Sets quota configuration for a given user",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) != 2 {
			err = errors.New("exactly two arguments required")
			return
		}

		var bytesSoft, bytesHard, filesSoft, filesHard uint64
		var bytesPresent, filesPresent bool

		if bytesSoft, bytesHard, bytesPresent, err = parseLimitsFlag(cmd, "bytes"); err != nil {
			return err
		}

		if filesSoft, filesHard, filesPresent, err = parseLimitsFlag(cmd, "files"); err != nil {
			return err
		}

		if err != nil {
			return
		}

		var u *user.User
		if u, err = lookupUser(args[1]); err != nil {
			return
		}

		var info *fsquota.Info
		limits := fsquota.Limits{}

		if bytesPresent {
			limits.Bytes.SetSoft(bytesSoft)
			limits.Bytes.SetHard(bytesHard)
		}

		if filesPresent {
			limits.Files.SetSoft(filesSoft)
			limits.Files.SetHard(filesHard)
		}

		if !filesPresent && !bytesPresent {
			err = errors.New("nothing to set")
			return
		}

		if info, err = fsquota.SetUserQuota(args[0], u, limits); err != nil {
			return
		}

		printQuotaInfo(cmd, info)
		return
	},
}

func init() {
	cmdUserSet.Flags().StringP("bytes", "b", "", "Byte limit in soft,hard format. ie. 1MiB,2GiB")
	cmdUserSet.Flags().StringP("files", "f", "", "File limit in soft,hard format, ie. 1M,2G")
	cmdUser.AddCommand(cmdUserSet)
}
