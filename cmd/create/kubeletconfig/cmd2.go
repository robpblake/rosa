package kubeletconfig

import (
	"context"
	"fmt"
	"github.com/openshift/rosa/pkg/output"
	"github.com/openshift/rosa/pkg/rosa"
	"github.com/openshift/rosa/pkg/runtime"
	"github.com/spf13/cobra"
)

func Create() *cobra.Command {

	o := NewCreateOptions()

	cmd := &cobra.Command{
		Use:     "kubeletconfig2",
		Aliases: []string{"kubelet-config2"},
		Short:   "Create a custom kubeletconfig for a cluster",
		Long:    "Create a custom kubeletconfig for a cluster",
		Example: `  # Create a custom kubeletconfig with a pod-pids-limit of 5000
  rosa create kubeletconfig2 --cluster=mycluster --pod-pids-limit=5000
  `,
		Run: runtime.RuntimeWithOCM(NewCreateCommandHandler(o)),
	}

	flags := cmd.Flags()
	flags.SortFlags = false
	output.AddFlag(cmd)
	o.AddFlags(flags)

	return cmd
}

func NewCreateCommandHandler(options *Options) runtime.CommandHandler {
	return func(ctx context.Context, r *rosa.Runtime, command *cobra.Command, args []string) error {
		//TODO - should this return a bool to indicate interactive input is required?
		err := options.Complete(command.Flags(), r)

		if err != nil {
			return err
		}

		fmt.Printf("I have pod_pids_limit of '%d'\n", options.podPidsLimit)

		//// TODO - Feels like this should live on the command specific options
		//clusterKey := r.GetClusterKey()
		//// TODO -
		//cluster, err := r.OCMClient.GetClusterByID(clusterKey, r.Creator)
		//if err != nil {
		//	return err
		//}
		//
		//if cluster.Hypershift().Enabled() {
		//	return fmt.Errorf("Hosted Control Plane clusters do not support custom KubeletConfig configuration.")
		//}
		//
		//if cluster.State() != cmv1.ClusterStateReady {
		//	return fmt.Errorf("Cluster '%s' is not yet ready. Current state is '%s'", clusterKey, cluster.State())
		//}
		//
		//kubeletConfig, err := r.OCMClient.GetClusterKubeletConfig(cluster.ID())
		//if err != nil {
		//	return fmt.Errorf("Failed getting KubeletConfig for cluster '%s': %s",
		//		cluster.ID(), err)
		//}
		//
		//if kubeletConfig != nil {
		//	return fmt.Errorf("A custom KubeletConfig for cluster '%s' already exists. "+
		//		"You should edit it via 'rosa edit kubeletconfig'", clusterKey)
		//}
		//
		//prompt := fmt.Sprintf("Creating the custom KubeletConfig for cluster '%s' will cause all non-Control Plane "+
		//	"nodes to reboot. This may cause outages to your applications. Do you wish to continue?", clusterKey)
		//
		//if confirm.ConfirmRaw(prompt) {
		//
		//	r.Reporter.Debugf("Creating KubeletConfig for cluster '%s'", clusterKey)
		//	kubeletConfigArgs := ocm.KubeletConfigArgs{PodPidsLimit: o.podPidsLimit}
		//
		//	_, err = r.OCMClient.CreateKubeletConfig(cluster.ID(), kubeletConfigArgs)
		//	if err != nil {
		//		return fmt.Errorf("Failed creating custom KubeletConfig for cluster '%s': '%s'",
		//			clusterKey, err)
		//	}
		//
		//	r.Reporter.Infof("Successfully created custom KubeletConfig for cluster '%s'", clusterKey)
		//}
		//
		//r.Reporter.Infof("Creation of custom KubeletConfig for cluster '%s' aborted.", clusterKey)
		return nil
	}
}
