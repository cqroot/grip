package cmd

import (
	"fmt"
	"os"

	"github.com/cqroot/grip/target"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grip",
	Short: "grip",
	Long:  `grip`,
	Args:  cobra.ExactArgs(1),
	Run:   runRootCmd,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runRootCmd(cmd *cobra.Command, args []string) {
	var targetName string = args[0]

	t := target.NewTarget(targetName)
	err := t.Validate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = t.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
