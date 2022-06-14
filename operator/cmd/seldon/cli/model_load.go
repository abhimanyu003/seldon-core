package cli

import (
	"os"

	"k8s.io/utils/env"

	"github.com/seldonio/seldon-core/operatorv2/pkg/cli"
	"github.com/spf13/cobra"
)

func createModelLoad() *cobra.Command {
	cmdModelLoad := &cobra.Command{
		Use:   "load",
		Short: "load a model",
		Long:  `load a model`,
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			schedulerHost, err := cmd.Flags().GetString(schedulerHostFlag)
			if err != nil {
				return err
			}
			filename, err := cmd.Flags().GetString(fileFlag)
			if err != nil {
				return err
			}
			showRequest, err := cmd.Flags().GetBool(showRequestFlag)
			if err != nil {
				return err
			}
			showResponse, err := cmd.Flags().GetBool(showResponseFlag)
			if err != nil {
				return err
			}
			schedulerClient := cli.NewSchedulerClient(schedulerHost)
			err = schedulerClient.LoadModel(loadFile(filename), showRequest, showResponse)
			return err
		},
	}
	cmdModelLoad.Flags().StringP(fileFlag, "f", "", "model file to load")
	if err := cmdModelLoad.MarkFlagRequired(fileFlag); err != nil {
		os.Exit(-1)
	}
	cmdModelLoad.Flags().String(schedulerHostFlag, env.GetString(EnvScheduler, DefaultScheduleHost), "seldon scheduler host")
	return cmdModelLoad
}