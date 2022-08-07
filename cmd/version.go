package cmd

import (
	"fmt"

	"muxtemp/pkg/config"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "golang api using gorilla mux",
	Long:  `golang api using gorilla mux`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadAppConfig()
		fmt.Printf("muxtemp %s version\n", config.AppConfig.Version)
	},
}
