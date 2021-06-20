package cmd

import (
	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"github.com/skmonir/mango/src/config"
	"github.com/skmonir/mango/src/intel"
	"github.com/skmonir/mango/src/system"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses a contest/problem",
	Run: func(cmd *cobra.Command, args []string) {
		argc := len(args)

		if argc == 0 {
			cfg, err := config.GetConfig()
			if err != nil {
				ansi.Println(color.RedString(err.Error()))
				return
			}
			if intel.GuessWorkspace(&cfg) {
				if err := system.Parse(cfg, cfg.CurrentContestId); err == nil {
					return
				}
			}
		}

		if argc < 1 {
			ansi.Println(color.RedString("too few argument for 'parse' command"))
		} else if argc > 1 {
			ansi.Println(color.RedString("too much argument for 'parse' command"))
		} else {
			cfg, err := config.GetConfig()
			if err != nil {
				ansi.Println(color.RedString(err.Error()))
				return
			}
			if cfg.Workspace == "" {
				ansi.Println(color.RedString("workspace directory path missing. run 'mango configure' to set workspace"))
				if !intel.GuessWorkspace(&cfg) {
					return
				}
			}
			if err := system.Parse(cfg, args[0]); err != nil {
				ansi.Println(color.RedString(err.Error()))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
