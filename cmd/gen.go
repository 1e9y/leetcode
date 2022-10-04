package cmd

import (
	"github.com/1e9y/leetcode/kit"
	"github.com/spf13/cobra"
)

var problem int

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate solution template",
	Run: func(cmd *cobra.Command, args []string) {
		if err := kit.Init(); err != nil {
			panic(err)
		}
		p, _ := kit.Get(problem)
		kit.Generate(p)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().IntVarP(&problem, "problem", "p", 0, "Problem ID")
}
