package cmd

import (
	"fmt"
	"os"
	"strings"

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
	var t string = args[0]

	if strings.HasPrefix(t, "github.com") {
		fmt.Printf("Read t from github: %s\n", t)
	} else {
		err := target.ValidateLocalTarget(t)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = target.ExecuteLocalTarget(t)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
