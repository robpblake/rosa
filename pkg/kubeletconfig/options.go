package kubeletconfig

import (
	"context"
	"fmt"
	"github.com/openshift/rosa/pkg/rosa"
)

type KubeletConfigOptions struct {
	PodPidsLimit int
}

func NewKubeletConfigOptions() *KubeletConfigOptions {
	return &KubeletConfigOptions{}
}

// Validate that the options the user has selected are correct
func (k *KubeletConfigOptions) Validate(_ context.Context, runtime *rosa.Runtime) error {
	maxPidsLimit, err := GetMaxPidsLimit(runtime.OCMClient)
	if err != nil {
		return fmt.Errorf("Failed to check maximum allowed Pids limit for cluster '%s'",
			runtime.ClusterKey)
	}

	if k.PodPidsLimit < MinPodPidsLimit {
		return fmt.Errorf("The minimum value for --pod-pids-limit is '%d'. You have supplied '%d'",
			MinPodPidsLimit, k.PodPidsLimit)
	}

	if k.PodPidsLimit > maxPidsLimit {
		return fmt.Errorf("The maximum value for --pod-pids-limit is '%d'. You have supplied '%d'",
			maxPidsLimit, k.PodPidsLimit)
	}
	return nil
}
