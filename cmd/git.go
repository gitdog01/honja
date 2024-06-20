package cmd

import (
	"fmt"

	"github.com/gitdog01/honja/internal/git"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		repoURL, _ := cmd.Flags().GetString("repo")
		destDir, _ := cmd.Flags().GetString("dest")
		if err := git.CloneRepo(repoURL, destDir); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringP("repo", "r", "", "Repository URL")
	cloneCmd.Flags().StringP("dest", "d", ".", "Destination directory")
	cloneCmd.MarkFlagRequired("repo")
}
