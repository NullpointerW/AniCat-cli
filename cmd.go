package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	host string
	port int
)

var rootCmd = &cobra.Command{
	Use:   "anicat",
	Short: "anicat-cli is a command-line client used to control anicat.",
	Long: `Use the "help" command to obtain usage instructions. For more details
                For more information, please refer to https://github.com/NullpointerW/AniCat and
                 https://github.com/NullpointerW/AniCat-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "O", "localhost", "server dial host")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 12314, "server dial port")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
