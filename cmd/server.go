/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"rahnit-rmm/config"
	"rahnit-rmm/pki"
	"rahnit-rmm/rpc"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		config.SetSubdir("server")

		addr := "localhost:1234"

		fmt.Printf("\nListening on localhost:%s\n", addr)

		commands := rpc.NewCommandCollection()

		_, err := pki.GetCaCert()

		if err == nil {
			// Server has a CA certificate

		} else {
			if errors.Is(err, pki.ErrNoCaCert) {
				// Server doesn't have a CA certificate yet
				commands.Add(rpc.UploadCaHandler)
			} else {
				panic(err)
			}
		}

		server, err := rpc.NewRpcServer(addr, commands)
		if err != nil {
			panic(err)
		}
		server.Run()

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
