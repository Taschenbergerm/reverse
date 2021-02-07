package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	reverse "github.com/taschenbergerm/reverse/pkg"
)

func init() {
	var fileName string
	var inputArguments bool

	pflag.StringVarP(&fileName, "read", "r", "", "read file to reverse")
	pflag.BoolVarP(&inputArguments, "input", "i",  false,"read input from arguments")
	var reverseCmd = &cobra.Command{
		Use: "reverse",
		Short: "reverse a string",
		Run: func ( cmd *cobra.Command, args []string){

			if inputArguments && fileName != "" {
				fmt.Println("Reverse cannot work with -r and -i together please use only one")
			}else if inputArguments {
				fmt.Println(reverse.RevSlice(args))
								
			}else {
				fmt.Println("read is not yet implemented")
				file, err := os.Open(fileName)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					fmt.Println(reverse.RevString(scanner.Text()))
				}

				if err := scanner.Err(); err != nil{
					log.Fatal(err)
				}
			}
		},
	}
	rootCmd.AddCommand(reverseCmd)
}
