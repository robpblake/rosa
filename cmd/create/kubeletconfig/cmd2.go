/*
Copyright (c) 2023 Red Hat, Inc.

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

package kubeletconfig

import (
	"context"
	"fmt"
	"github.com/openshift/rosa/pkg/interactive/confirm"
	"os"

	"github.com/spf13/cobra"

	"github.com/openshift/rosa/pkg/interactive"
	. "github.com/openshift/rosa/pkg/kubeletconfig"
	"github.com/openshift/rosa/pkg/ocm"
	"github.com/openshift/rosa/pkg/rosa"
)

// RosaCommandRun defines a function that should be returned to implement the actual action of the command
type RosaCommandRun func(ctx context.Context, r *rosa.Runtime, command *cobra.Command, args []string) error

// RosaRuntimeVisitor defines a function that can visit the Runtime and configure it as required. Default implementation
// of this would exist to facilitate the most common scenarios e.g with OCM only
type RosaRuntimeVisitor func(ctx context.Context, r *rosa.Runtime, command *cobra.Command, args []string) error

// This is the default runnable function for all commands. It takes care of instantiating the runtime and context
// and then invokes the runtime visitor and finally, the commandrun.
func DefaultRosaCommandRun(visitor RosaRuntimeVisitor, commandRun RosaCommandRun) func(command *cobra.Command, args []string) {
	return func(command *cobra.Command, args []string) {
		ctx := context.Background()
		runtime := rosa.NewRuntime()
		defer runtime.Cleanup()

		err := visitor(ctx, runtime, command, args)
		if err != nil {
			runtime.Reporter.Errorf(err.Error())
			os.Exit(1)
		}

		err = commandRun(ctx, runtime, command, args)
		if err != nil {
			runtime.Reporter.Errorf(err.Error())
			os.Exit(1)
		}
	}
}

// Return the command fully initialised, but not globally scoped within the package
func CreateKubeletConfig() *cobra.Command {

	options := NewKubeletConfigOptions()

	cmd := &cobra.Command{
		Use:     "kubeletconfig",
		Aliases: []string{"kubelet-config"},
		Short:   "Create a custom kubeletconfig for a cluster",
		Long:    "Create a custom kubeletconfig for a cluster",
		Example: `  # Create a custom kubeletconfig with a pod-pids-limit of 5000
  rosa create kubeletconfig --cluster=mycluster --pod-pids-limit=5000
  `,
		Run: DefaultRosaCommandRun(CreateKubeletConfigRuntimeVisitor(), CreateKubeletConfigRun(options)),
	}

	flags := cmd.Flags()
	flags.SortFlags = false
	flags.IntVar(
		&options.PodPidsLimit,
		PodPidsLimitOption,
		PodPidsLimitOptionDefaultValue,
		PodPidsLimitOptionUsage)

	ocm.AddClusterFlag(cmd)
	interactive.AddFlag(flags)
	return cmd
}

// Visitor function that customises the runtime for this command
func CreateKubeletConfigRuntimeVisitor() RosaRuntimeVisitor {
	return func(ctx context.Context, r *rosa.Runtime, command *cobra.Command, args []string) error {
		r.WithOCM()
		return nil
	}
}

// The actual logic of the command. os.Exit is not allowed during the execution of this function.
// errors are bubbled upwards and handled centrally.
func CreateKubeletConfigRun(options KubeletConfigOptions) RosaCommandRun {
	return func(ctx context.Context, r *rosa.Runtime, command *cobra.Command, args []string) error {

		client := NewKubeletConfigClient(ctx, r)
		_, err := client.CanCreateKubeletConfig(r.ClusterKey)
		if err != nil {
			return err
		}

		// Interactive mode is a function of the cobra layer and not ROSA core
		if options.PodPidsLimit == PodPidsLimitOptionDefaultValue && !interactive.Enabled() {
			interactive.Enable()
			r.Reporter.Infof("Enabling interactive mode")
			options.PodPidsLimit, err = interactive.GetInt(GetInteractiveInput(MaxPodPidsLimit, nil))
			if err != nil {
				return err
			}
		}

		// Could also be delegated to the KubeletConfigClient
		err = options.Validate(ctx, r)
		if err != nil {
			return err
		}

		prompt := fmt.Sprintf("Creating the custom KubeletConfig for cluster '%s' will cause all non-Control Plane "+
			"nodes to reboot. This may cause outages to your applications. Do you wish to continue?", r.ClusterKey)

		if confirm.ConfirmRaw(prompt) {

			_, err = client.CreateKubeletConfig(r.ClusterKey, options)
			if err != nil {
				return err
			}

			r.Reporter.Infof("Successfully created custom KubeletConfig for cluster '%s'", r.ClusterKey)
		}

		r.Reporter.Infof("Creation of custom KubeletConfig for cluster '%s' aborted.", r.ClusterKey)
		return nil
	}
}
