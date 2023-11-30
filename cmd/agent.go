/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/rahn-it/svalin/config"
	"github.com/rahn-it/svalin/pki"
	"github.com/rahn-it/svalin/rmm"
	"github.com/rahn-it/svalin/rpc"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("agent called")

		err := config.SetSubdir("agent")
		if err != nil {
			panic(err)
		}

		credentials, err := pki.GetHostCredentials()
		if err != nil {
			if errors.Is(err, pki.ErrNotInitialized) {
				log.Printf("agent not yet initialized")
				credentials, err = rpc.EnrollWithUpstream()
				if err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}

		agent, err := rmm.AgentConnect(context.Background(), credentials)

		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			err = agent.Run()
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-interrupt
			err := agent.Close()
			if err != nil {
				err := fmt.Errorf("error shutting down program: error closing agent: %w", err)
				log.Println(err)
			} else {
				log.Println("Agent closed without errors")
			}
		}()

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
