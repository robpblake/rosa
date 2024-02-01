/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package create

import (
	"github.com/spf13/cobra"

	"github.com/openshift/rosa/cmd/create/kubeletconfig"
	"github.com/openshift/rosa/pkg/arguments"
	"github.com/openshift/rosa/pkg/interactive/confirm"
)

var Cmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a resource from stdin",
	Long:    "Create a resource from stdin",
}

func init() {
	//Cmd.AddCommand(accountroles.Cmd)
	//Cmd.AddCommand(admin.Cmd)
	//Cmd.AddCommand(cluster.Cmd)
	//Cmd.AddCommand(idp.Cmd)
	//Cmd.AddCommand(machinepool.Cmd)
	//Cmd.AddCommand(oidcconfig.Cmd)
	//Cmd.AddCommand(oidcprovider.Cmd)
	//Cmd.AddCommand(operatorroles.Cmd)
	//Cmd.AddCommand(userrole.Cmd)
	//Cmd.AddCommand(ocmrole.Cmd)
	//Cmd.AddCommand(service.Cmd)
	//Cmd.AddCommand(tuningconfigs.Cmd)
	//Cmd.AddCommand(dnsdomains.Cmd)
	//Cmd.AddCommand(autoscaler.Cmd)
	//Cmd.AddCommand(kubeletconfig.Cmd)
	//
	//flags := Cmd.PersistentFlags()
	//arguments.AddProfileFlag(flags)
	//arguments.AddRegionFlag(flags)
	//confirm.AddFlag(flags)
	//
	//globallyAvailableCommands := []*cobra.Command{
	//	accountroles.Cmd, operatorroles.Cmd,
	//	userrole.Cmd, ocmrole.Cmd,
	//	oidcprovider.Cmd,
	//}
	//arguments.MarkRegionHidden(Cmd, globallyAvailableCommands)
}

// Sanity check that the command meets our requirements:

func RegisterCommand(parent *cobra.Command, child *cobra.Command) {
	if child.Run == nil {
		// TODO - fail, you haven't set the run function
	}

	if child.Use == "" {
		// TODO - fail, you haven't set use
	}

	if child.Long == "" {

	}

	if child.Short == "" {

	}

	// All good, register the command with the parent
	parent.AddCommand(child)
}

func Create() *cobra.Command {
	create := &cobra.Command{
		Use:     "create",
		Aliases: []string{"add"},
		Short:   "Create a resource from stdin",
		Long:    "Create a resource from stdin",
	}

	flags := Cmd.PersistentFlags()
	arguments.AddProfileFlag(flags)
	arguments.AddRegionFlag(flags)
	confirm.AddFlag(flags)

	RegisterCommand(create, kubeletconfig.Create())
	return create
}
