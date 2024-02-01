package runtime

import (
	"context"
	"github.com/openshift/rosa/pkg/rosa"
	"github.com/spf13/cobra"
	"os"
)

// CommandHandler is a function that does the work of the command. It expects to be passed a context, a runtime
// and then command being executed upon.
type CommandHandler func(ctx context.Context, r *rosa.Runtime, command *cobra.Command, args []string) error

func RuntimeWithOCM(handler CommandHandler) func(command *cobra.Command, args []string) {
	r := rosa.NewRuntime().WithOCM()
	return wrapAndExecuteHandler(r, handler)
}

// Commands that need both OCM and AWS clients
func RuntimeWithOCMAndAWS(handler CommandHandler) func(command *cobra.Command, args []string) {
	r := rosa.NewRuntime().WithOCM().WithAWS()
	return wrapAndExecuteHandler(r, handler)
}

func wrapAndExecuteHandler(runtime *rosa.Runtime, handler CommandHandler) func(command *cobra.Command, args []string) {
	return func(command *cobra.Command, args []string) {
		ctx := context.Background()
		defer runtime.Cleanup()

		err := handler(ctx, runtime, command, args)
		if err != nil {
			runtime.Reporter.Errorf(err.Error())
			os.Exit(1)
		}
	}
}
