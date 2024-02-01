package kubeletconfig

import (
	"fmt"
	"github.com/openshift/rosa/pkg/interactive"
	"github.com/openshift/rosa/pkg/interactive/confirm"
	"github.com/openshift/rosa/pkg/kubeletconfig"
	"github.com/openshift/rosa/pkg/rosa"
	"github.com/spf13/pflag"
)

// Each command defines it's own Options struct, which represents what it can be configured by on the CLI
type Options struct {
	podPidsLimit int
}

// Bind the command flags to the values of the Options
func (o *Options) AddFlags(flags *pflag.FlagSet) {
	flags.IntVar(
		&o.podPidsLimit,
		"pod-pids-limit",
		0,
		"This is how to use --pod-pids-limit")

	interactive.AddFlag(flags)
	confirm.AddFlag(flags)
}

// Validate the flags that have been set by the user on the command line
func (o *Options) Complete(flags *pflag.FlagSet, r *rosa.Runtime) error {

	// TODO - Where are we handling interactive mode?
	if o.podPidsLimit == 0 && !interactive.Enabled() {
		interactive.Enable()
		r.Reporter.Infof("Enabling interactive mode")
	}

	maxPidsLimit, err := kubeletconfig.GetMaxPidsLimit(r.OCMClient)
	if err != nil {
		return err
	}

	if interactive.Enabled() {
		//o.podPidsLimit, err = interactive.GetInt(GetInteractiveInput(maxPidsLimit, kubeletConfig))
		if err != nil {
			return err
		}
	}

	if o.podPidsLimit < kubeletconfig.MinPodPidsLimit {
		return fmt.Errorf("The minimum value for --pod-pids-limit is '%d'. You have supplied '%d'",
			kubeletconfig.MinPodPidsLimit, o.podPidsLimit)
	}

	if o.podPidsLimit > maxPidsLimit {
		return fmt.Errorf("The maximum value for --pod-pids-limit is '%d'. You have supplied '%d'",
			maxPidsLimit, o.podPidsLimit)
	}
	return nil
}

func NewCreateOptions() *Options {
	return &Options{
		podPidsLimit: 0,
	}
}
