package cmd

import (
	"fmt"

	"github.com/1e9y/leetcode/kit"
	"github.com/spf13/cobra"
)

var problem int

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate solution scaffold",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
		fmt.Println(problem)

		//lc := new(kit.LeetCode)
		if err := kit.Init(); err != nil {
			panic(err)
		}
		p, _ := kit.Get(problem)
		//fmt.Printf("%v", g)
		kit.Generate(p)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.
	genCmd.Flags().IntVarP(&problem, "problem", "p", 0, "Problem ID")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
