package cmd

import (
	"muxtemp/internal/server"

	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start muxtemp HTTP server",
		Long:  "Start muxtemp HTTP server",
		Run:   executeMuxtemp,
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

func executeMuxtemp(cmd *cobra.Command, args []string) {
	server.Start()
}
