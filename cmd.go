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
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run hugo...")
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

	fmt.Println(os.Args)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
