package greet

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "greet",
	Short: "Print a simple greeting",
	Run: func(x *cobra.Command, args []string) {
		fmt.Println(`Hello there.`)
	},
}

func init() {
	Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
