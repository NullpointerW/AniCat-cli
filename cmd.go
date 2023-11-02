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
	Short: "anicat-cli is a command-line client used to control anicat",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Subscribe to anime series",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var ls = &cobra.Command{
	Use:   "ls",
	Short: "Show detailed information of subjects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var lsi = &cobra.Command{
	Use:   "lsi",
	Short: "Show resource list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
var stat = &cobra.Command{
	Use:   "stat",
	Short: "Show downloading status with the subject.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var rm = &cobra.Command{
	Use:   "rm",
	Short: "Delete a subject",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "server dial host")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 12314, "server dial port")
	rootCmd.AddCommand(add)
	rootCmd.AddCommand(ls)
	rootCmd.AddCommand(lsi)
	rootCmd.AddCommand(stat)
	rootCmd.AddCommand(rm)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
