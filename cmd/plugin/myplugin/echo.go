package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var EchoCmd = &cobra.Command{
	Use:   "echo (text to echo)",
	Short: "Prints text providied to stdout",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, " "))
	},
}
