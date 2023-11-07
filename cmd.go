package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var (
	host    string
	port    int
	address string
)

var rootCmd = &cobra.Command{
	Use:   "anicat",
	Short: "anicat-cli is a command-line client used to control anicat",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Subscribe to anime series",
	Run: func(cmd *cobra.Command, args []string) {

	},
	TraverseChildren: true,
}

var feed = &cobra.Command{
	Use:   "feed",
	Short: "Subscribe to anime series via rss feed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
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

var stop = &cobra.Command{
	Use:   "stop",
	Short: "Stop progress",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "Server dial host")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 12314, "Server dial port")
	addFlags()
	lsi.Flags().BoolVarP(&searchList, "search", "s", false, "Show search list")
	rootCmd.AddCommand(add)
	add.AddCommand(feed)
	rootCmd.AddCommand(ls)
	rootCmd.AddCommand(lsi)
	rootCmd.AddCommand(stat)
	rootCmd.AddCommand(rm)
	rootCmd.AddCommand(stop)
	address = solveAddress()
}

var (
	contain    string
	exclude    string
	useRegEXP  bool
	group      string
	index      string
	searchList bool
	// add feed --name
	infoNameForFeed string
)

func addFlags() {
	add.Flags().StringVarP(&contain, "contain", "c", "", "Contained keywords")
	add.Flags().StringVarP(&exclude, "exclude", "e", "", "Excluded keywords")
	add.Flags().BoolVarP(&useRegEXP, "regexp", "r", false, "Use regular expressions")
	add.Flags().StringVarP(&group, "group", "g", "", "Subtitle group keywords")
	add.Flags().StringVarP(&index, "index", "i", "", "Index of torrent list")
	add.Flags().StringVarP(&infoNameForFeed, "name", "n", "", "KeyName for solve info on AddFeed")
}

func solveAddress() string {
	def := "127.0.0.1:12314"
	if port != 0 && host != "" {
		def = host + ":" + strconv.Itoa(port)
	}
	return def
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
