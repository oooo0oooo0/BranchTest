package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"fobrain/cmd/migrate"
	"fobrain/cmd/webhook"
)

var rootCmd = &cobra.Command{
	Use:          "fobrain",
	Short:        "fobrain",
	SilenceUsage: true,
	Long:         `fobrain`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	fmt.Printf("%s\n", `
 ______   ______     ______     ______     ______     __     __   __    
/\  ___\ /\  __ \   /\  == \   /\  == \   /\  __ \   /\ \   /\ "-.\ \   
\ \  __\ \ \ \/\ \  \ \  __<   \ \  __<   \ \  __ \  \ \ \  \ \ \-.  \  
 \ \_\    \ \_____\  \ \_____\  \ \_\ \_\  \ \_\ \_\  \ \_\  \ \_\\"\_\ 
  \/_/     \/_____/   \/_____/   \/_/ /_/   \/_/\/_/   \/_/   \/_/ \/_/  
`)
	fmt.Println("fobrain-feat/a")
}

func init() {
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(webhook.WebhookCmd)

	fmt.Println("fobrain-feat/a")
}

// Execute : apply commands
func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
