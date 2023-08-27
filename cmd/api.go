package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/kevinsantana/wallet-view-api/pkg/version"
	"github.com/kevinsantana/wallet-view-api/internal/config"
	"github.com/kevinsantana/wallet-view-api/internal/server"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the http server.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		log.WithField("project_version", version.PROJECT_VERSION)
		config.InitConfig(ctx)
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
