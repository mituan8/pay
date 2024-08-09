package cmd

import (
	"context"

	"github.com/mituan8/pay/internal/kms"
	"github.com/mituan8/pay/pkg/graceful"
	"github.com/spf13/cobra"
)

var serverKMSCommand = &cobra.Command{
	Use:   "serve-kms",
	Short: "Start KMS (Key Management Server)",
	Run:   serveKMS,
}

func serveKMS(_ *cobra.Command, _ []string) {
	service := kms.NewApp(context.Background(), resolveConfig())
	service.Run()

	if err := graceful.WaitShutdown(); err != nil {
		service.Logger().Error().Err(err).Msg("unable to shutdown service gracefully")
		return
	}

	service.Logger().Info().Msg("shutdown complete")
}
