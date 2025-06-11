package cmd

import (
	"fmt"
	"os"

	// funcs "BLAH/my-cli-app/funcs"

	"github.com/shikmade/my-cli-app/funcs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Get url",
	Args:  cobra.MaximumNArgs(1), // Expects at most ONE argument
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 { 
			url := args[0] 
			fmt.Println("Getting:", url)
			funcs.SendGet(url)
		} else {
			fmt.Println("Please provide a URL")
		}
	},
}


// make the post cli command
// take the cli as the params of the Get_data functino
var postCmd = &cobra.Command{
	Use:   "post [url] [a1] [a2] [a3]",
	Short: "Post url",
	Args:  cobra.MaximumNArgs(10),
	Run: func(cmd *cobra.Command, args []string) {

		arg1 := args[1]
		arg2 := args[2]
		arg3 := args[3]

		if len(args) > 0 {
			url := args[0]
			fmt.Println("Posting to:", url)
			data := funcs.Get_data(arg1,arg2,arg3)
			fmt.Println(data)
			funcs.SendPost(url,data)
		} else {
			fmt.Println("Please provide a URL")
		}
	},
}


		// for arg,x  := range args {
		// 	// fmt.Println("bludad:",arg)
		// 	fmt.Printf("ARG: %d, %s\n",arg,x)
		// 	// fmt.Println("bludadsadad:",x)
		// 	if arg == 1 {

		// 	}
		// }
		// url := strings.Join(args, " ")

		// fmt.Println("blud: ", url)


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add the get command to the root command
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(postCmd)
}