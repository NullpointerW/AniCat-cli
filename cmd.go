package main

import (
	"encoding/json"
	"fmt"
	CMD "github.com/NullpointerW/anicat/net/cmd"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var (
	host    string
	port    int
	address string
)

var rootCmd = &cobra.Command{
	Use:   "anicat",
	Short: "AniCat-Cli is a command-line client used to control AniCat",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在命令行参数解析之后，再调用 solveAddress
		address = solveAddress()
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Subscribe to anime series",
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, " ")
		flag := CMD.AddFlag{
			MustContain:    contain,
			MustNotContain: exclude,
			UseRegexp:      useRegEXP,
			Group:          group,
			Index:          index,
		}
		raw, _ := json.Marshal(flag)
		c := CMD.Cmd{
			Arg: arg,
			Cmd: CMD.Add,
			Raw: raw,
		}
		resp, err := Send(address, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(1),
}

var feed = &cobra.Command{
	Use:   "feed",
	Short: "Subscribe to anime series via rss feed",
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, "")
		flag := CMD.AddFlag{
			MustContain:    contain,
			MustNotContain: exclude,
			UseRegexp:      useRegEXP,
			FeedInfoName:   infoNameForFeed,
		}
		raw, _ := json.Marshal(flag)
		c := CMD.Cmd{
			Arg: arg,
			Cmd: CMD.Add,
			Raw: raw,
		}
		resp, err := Send(address, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
	Args: cobra.MinimumNArgs(1),
}

var ls = &cobra.Command{
	Use:   "ls",
	Short: "Show detailed information of subjects",
	Run: func(cmd *cobra.Command, args []string) {
		c := CMD.Cmd{
			Cmd: CMD.Ls,
			Raw: nil,
		}
		resp, err := Send(address, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
}

var lsi = &cobra.Command{
	Use:   "lsi",
	Short: "Show resource list",
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, "")
		flag := CMD.LsiFlag{SearchList: searchList}
		raw, _ := json.Marshal(flag)
		c := CMD.Cmd{
			Arg: arg,
			Cmd: CMD.LsItems,
			Raw: raw,
		}
		resp, err := Send(address, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
	Args: cobra.MinimumNArgs(1),
}
var stat = &cobra.Command{

	Use:   "stat",
	Short: "Show downloading status with the subject.",
	Run: func(cmd *cobra.Command, args []string) {
		c := CMD.Cmd{
			Arg: args[0],
			Cmd: CMD.Status,
		}
		resp, err := Send(address, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
	Args: cobra.ExactArgs(1),
}

var rm = &cobra.Command{
	Use:   "rm",
	Short: "Delete a subject",
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, " ")
		c := CMD.Cmd{
			Arg: arg,
			Cmd: CMD.Remove,
		}
		resp, err := Send(address, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
	Args: cobra.MinimumNArgs(1),
}

var stop = &cobra.Command{
	Use:   "stop",
	Short: "Stop progress",
	Run: func(cmd *cobra.Command, args []string) {
		c := CMD.Cmd{
			Cmd: CMD.Stop,
		}
		resp, err := Send(address, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "Server dial host")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 12314, "Server dial port")
	addFlags()
	feed.Flags().StringVarP(&infoNameForFeed, "name", "n", "", "KeyName for solve info on AddFeed")
	lsi.Flags().BoolVarP(&searchList, "search", "s", false, "Show search list")
	rootCmd.AddCommand(add)
	add.AddCommand(feed)
	rootCmd.AddCommand(ls)
	rootCmd.AddCommand(lsi)
	rootCmd.AddCommand(stat)
	rootCmd.AddCommand(rm)
	rootCmd.AddCommand(stop)
}

var (
	contain    string
	exclude    string
	useRegEXP  bool
	group      string
	index      int
	searchList bool
	// add feed --name
	infoNameForFeed string
)

func addFlags() {
	add.Flags().StringVarP(&contain, "contain", "c", "", "Contained keywords")
	add.Flags().StringVarP(&exclude, "exclude", "e", "", "Excluded keywords")
	add.Flags().BoolVarP(&useRegEXP, "regexp", "r", false, "Use regular expressions")
	add.Flags().StringVarP(&group, "group", "g", "", "Subtitle group keywords")
	add.Flags().IntVarP(&index, "index", "i", 0, "Index of torrent list")

}

func solveAddress() string {
	def := "127.0.0.1:12314"
	if port != 0 && host != "" {
		def = host + ":" + strconv.Itoa(port)
	}
	return def
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
