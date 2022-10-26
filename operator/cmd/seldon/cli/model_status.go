package cli

import (
	"k8s.io/utils/env"

	"github.com/seldonio/seldon-core/operatorv2/pkg/cli"
	"github.com/spf13/cobra"
)

func createModelStatus() *cobra.Command {
	cmdModelStatus := &cobra.Command{
		Use:   "status <modelName>",
		Short: "get status for model",
		Long:  `get the status for a model`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			schedulerHost, err := cmd.Flags().GetString(flagSchedulerHost)
			if err != nil {
				return err
			}
			authority, err := cmd.Flags().GetString(flagAuthority)
			if err != nil {
				return err
			}
			showRequest, err := cmd.Flags().GetBool(flagShowRequest)
			if err != nil {
				return err
			}
			showResponse, err := cmd.Flags().GetBool(flagShowResponse)
			if err != nil {
				return err
			}
			modelWaitCondition, err := cmd.Flags().GetString(flagWaitCondition)
			if err != nil {
				return err
			}
			modelName := args[0]

			schedulerClient, err := cli.NewSchedulerClient(schedulerHost, authority)
			if err != nil {
				return err
			}

			err = schedulerClient.ModelStatus(modelName, showRequest, showResponse, modelWaitCondition)
			return err
		},
	}

	cmdModelStatus.Flags().String(flagSchedulerHost, env.GetString(envScheduler, defaultSchedulerHost), helpSchedulerHost)
	cmdModelStatus.Flags().String(flagAuthority, "", helpAuthority)
	cmdModelStatus.Flags().StringP(flagWaitCondition, "w", "", "model wait condition")

	return cmdModelStatus
}
