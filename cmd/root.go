package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "honja",
	Short: "Honja CLI application",
	Long:  `A CLI application to manage Docker and GitHub repositories.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// 여기에서 하위 명령어를 추가할 수 있습니다.
}
