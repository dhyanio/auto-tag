package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	appVersion  = "v0.0.1-dev"
	longAppDesc = "AutoTag: Automatic multi-cloud resource tagging for cost allocation and ownership tracking"
)

var (
	rootCMD = &cobra.Command{
		Long: longAppDesc,
	}
)

// flagError is an erro type for the flag errors
type flagError struct{ err error }

func (e flagError) Error() string { return e.err.Error() }

// Execute runs the root command
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		if !errors.As(err, &flagError{}) {
			panic(err)
		}
	}
}

func init() {
	rootCMD.AddCommand(versionCMD())
}
