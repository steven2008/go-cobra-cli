/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called", "foo:", *foo)
		fmt.Println("serve called", "print:", print)
		fmt.Println("serve called", "show:", show)
		fmt.Println("serve called", "foo2:", *foo2)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	foo2 = serveCmd.Flags().String("foo2", "fo2", "help for foo2")
	serveCmd.PersistentFlags().StringVar(&print2, "print2", "defaultPrint2", "print2")
	serveCmd.Flags().BoolP("serve", "s", false, "Help message for serve")

}
