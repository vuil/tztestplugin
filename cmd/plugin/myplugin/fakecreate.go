package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
)

var FakeCreateClusterCmd = &cobra.Command{
	Use:   "fakecreate clustername",
	Short: "Fake simulation of create cluster",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Fake create cluster %s\n", args[0])
		fmt.Printf("Sync-ing plugins...\n")
		output, err := plugin.SyncPluginsForTarget(types.TargetK8s)
		if err == nil {
			fmt.Printf("SYNC OK: output=\n%s\n", output)
		} else {
			fmt.Printf("SYNC FAILED: output=\n%v\n%s\n", err, output)
		}
	},
}
