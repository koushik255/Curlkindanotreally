package cmd

import (
	"fmt"
	"os"

	// funcs "BLAH/my-cli-app/funcs"

	"github.com/shikmade/my-cli-app/funcs"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Define the get command
var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Get url",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		for arg,x := range args {
			fmt.Println(arg,x)
			funcs.SendGet(x)

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
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-cli-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add the get command to the root command
	rootCmd.AddCommand(getCmd)
}