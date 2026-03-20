package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type podsOptions struct {
	namespace string
	output    string
}

// rootCmd is a pointer to &cobra.Command{}; rootCmd is of type *cobra.Command
var rootCmd = &cobra.Command{
	Use:   "devtool",
	Short: "A CLI tool for DevOps tasks",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("devtool v0.0.1")
	},
}

var mainOptions = &podsOptions{}

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "List pods in a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing pods in namespace: %s\n", mainOptions.namespace)
		fmt.Printf("Output format: %s\n", mainOptions.output)
	},
}

func main() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	podsCmd.Flags().StringVarP(&mainOptions.namespace, "namespace", "n", "default", "Kubernetes namespace")
	podsCmd.Flags().StringVarP(&mainOptions.output, "output", "o", "table", "Output format (table|json|yaml)")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(podsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
