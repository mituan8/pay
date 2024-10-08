package cmd

import (
	"context"

	"github.com/mituan8/pay/internal/app"
	"github.com/mituan8/pay/pkg/graceful"
	"github.com/spf13/cobra"
)

var runSchedulerCommand = &cobra.Command{
	Use:   "run-scheduler",
	Short: "Start Scheduler Service",
	Run:   runScheduler,
}

func runScheduler(_ *cobra.Command, _ []string) {
	ctx := context.Background()

	service := app.New(ctx, resolveConfig())
	service.RunScheduler()

	if err := graceful.WaitShutdown(); err != nil {
		service.Logger().Error().Err(err).Msg("unable to shutdown service gracefully")
		return
	}

	service.Logger().Info().Msg("shutdown complete")
}
