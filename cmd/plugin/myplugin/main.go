package main

import (
	"os"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
    "github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
    "github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:        "myplugin",
	Description: "demo plugin",
	Target:      types.TargetUnknown, // <<<FIXME! set the Target of the plugin to one of {TargetGlobal,TargetK8s,TargetTMC}
	Version:     buildinfo.Version,
	BuildSHA:    buildinfo.SHA,
	Group:       plugin.ManageCmdGroup, // set group
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "")
	}
	p.AddCommands(
		// Add commands
	)
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
