package cmd

import (
	"fmt"

	"github.com/gitdog01/honja/internal/docker"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Manage Docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		image, _ := cmd.Flags().GetString("image")
		if err := docker.RunDockerContainer(image); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().StringP("image", "i", "", "Docker image to run")
	dockerCmd.MarkFlagRequired("image")
}
