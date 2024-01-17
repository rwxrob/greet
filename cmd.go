package greet

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(hiCmd)
}

var Cmd = &cobra.Command{
	Use:   "greet",
	Short: "Print a simple greeting",
	Run: func(x *cobra.Command, args []string) {
		fmt.Println(`Hello there.`)
	},
}

var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "Say hi",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi there.")
	},
}
