package cmd

import (
	"fmt"
	"github.com/taschenbergerm/reverse/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number",
	Long: "Print the version number of reverse",
	Run: func(cmd *cobra.Command, args  []string){
		fmt.Println("Build Date", version.BuildDate)
		fmt.Println("Build Arch", version.OsArch)
		fmt.Println("GO version", version.GoVersion)
	},
}